package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Coin struct {
	USDBRL USDBRL
}

type USDBRL struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

var db *sql.DB
var err error

func main() {
	err := connectDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
		return
	}

	createTable()

	http.HandleFunc("/cotacao", handler)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	coin, err := getCoin(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	err = insertCoin(ctx, coin.USDBRL)
	if err != nil {
		http.Error(w, "erro ao salvar no banco", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"bid": coin.USDBRL.Bid,
	})
}

func getCoin(ctx context.Context) (Coin, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return Coin{}, errors.New("erro ao criar requisição")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Coin{}, errors.New("erro na requisição externa")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Coin{}, errors.New("erro ao ler resposta")
	}

	var data Coin
	if err := json.Unmarshal(body, &data); err != nil {
		return Coin{}, errors.New("erro no unmarshal")
	}

	return data, nil
}

func insertCoin(ctx context.Context, u USDBRL) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	query := `
		INSERT INTO cotacoes 
		(code, codein, name, high, low, varbid, pctchange, bid, ask, timestamp, createdate) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.ExecContext(ctx, query,
		u.Code, u.Codein, u.Name, u.High, u.Low,
		u.VarBid, u.PctChange, u.Bid, u.Ask,
		u.Timestamp, u.CreateDate,
	)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return errors.New("timeout ao inserir no banco")
		}
		return err
	}

	return nil
}

func connectDB() error {
	dsn := "root:123456@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
		return err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao fazer ping no banco:", err)
		return err
	}

	return nil
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INT AUTO_INCREMENT PRIMARY KEY,
		code VARCHAR(10),
		codein VARCHAR(10),
		name VARCHAR(80),
		high VARCHAR(30),
		low VARCHAR(30),
		varbid VARCHAR(30),
		pctchange VARCHAR(30),
		bid VARCHAR(30),
		ask VARCHAR(30),
		timestamp VARCHAR(30),
		createdate VARCHAR(30)
	);
	`
	if _, err := db.Exec(query); err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}
}
