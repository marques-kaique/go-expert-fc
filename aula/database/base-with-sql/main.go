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

func updateProduct(db *sql.DB, product *Products) error {
	smtp, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer smtp.Close()

	_, err = smtp.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func selectProduct(db *sql.DB, id string) (*Products, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Products

	// O queryRow retorna um único registro
	// O Scan é utilizado para ler os valores do registro e atribuí-los às variáveis
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	// pode ser utilizado o QueryRowContext para passar o contexto
	// assim teria tambem o timeout
	// ou cancelar a query
	//err = stmt.QueryRowContext(ctx, id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func selectProducts(db *sql.DB) ([]*Products, error) {
	// por ser varias linhas, sem utilizar o where
	// pode ser utilizado o query
	// nao teria preocupação com sql injection
	// rows, err := db.Query("select id, name, price from products")
	stmt, err := db.Prepare("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var products []*Products

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// o retorno volta em um slice de linhas
	// o Next é utilizado para iterar sobre as linhas
	for rows.Next() {
		var product Products
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		// adiciona o produto ao slice de produtos
		products = append(products, &product)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
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
	fmt.Printf("Product: %v, possui o preço de %.2f\n", p.ID, p.Price)

	p.Price = 52.02
	err = updateProduct(db, p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de %.2f\n", p.ID, p.Price)

	product, err := selectProduct(db, p.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de %.2f\n", product.ID, product.Price)

	products, err := selectProducts(db)
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", product.ID, product.Price)
	}

	err = deleteProduct(db, p.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Product deleted successfully")
}
