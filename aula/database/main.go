package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	// importa o driver do mysql, o _ é para importar e não utilizar
	//
	_ "github.com/go-sql-driver/mysql"
)

type Products struct {
	ID    string
	Name  string
	Price float64
}

func newProducts(name string, price float64) *Products {
	return &Products{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func insertProduct(db *sql.DB, p *Products) error {
	// para proteger de sql injection, é necessário utilizar o prepare
	// quando ele subsituir o ?, ele vai substituir pelo valor que está no slice
	// o prepare retorna um statement, que é uma representação do comando sql
	// o statement é utilizado para executar o comando sql
	stmt, err := db.Prepare("insert into products (id, name, price) values (?, ?, ?)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.ID, p.Name, p.Price)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/goexpert")

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connected to the database", db.Ping())
	p := newProducts("test", 10.02)

	err = insertProduct(db, p)
	if err != nil {
		panic(err)
	}
	println(p.ID)
}
