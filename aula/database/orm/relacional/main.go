package main

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	Name string
	gorm.Model
	Products []Products // has many -> indica que a categoria tem mais de 1 produto
}

type Products struct {
	Name         string
	Price        float64
	CategoryID   uint         // foreign key, belongs to
	Category     Category     // belongs to
	SerialNumber SerialNumber `gorm:"foreignKey:ProductID"`      // has one
	Group        []Group      `gorm:"many2many:group_products;"` // many to many
	gorm.Model
}

type SerialNumber struct { // esses será o has one
	ID        int `gorm:"primaryKey"`
	ProductID uint
	Number    string
}

type Group struct {
	Name string
	Tipo int
	gorm.Model
	Products []Products `gorm:"many2many:group_products;"` // many to many -> group_products é a tabela de relacionamento
}

func main() {
	dsn := "root:123456@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Category{}, &Products{}, &SerialNumber{}, &Group{})

	// create category
	category := Category{Name: "Electronics"}
	db.Create(&category)
	// or
	//db.Create(&Category{Name: "Electronics"})

	// create product
	db.Create(&Products{
		Name:       "Laptop",
		Price:      1000,
		CategoryID: category.ID,
	})

	// find all products
	var products []Products
	db.Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Price, p.Category.Name)
	}

	// para tarzer dados das categorias, necessário utilizar o Preload
	db.Preload("Category").Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Price, p.Category.Name)
	}

	// --- HAS ONE

	product := Products{
		Name:       "Mouse",
		Price:      50,
		CategoryID: category.ID,
	}

	db.Create(&product)

	db.Create(&SerialNumber{
		ProductID: product.ID,
		Number:    uuid.New().String(),
	})

	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Price, p.Category.Name, p.SerialNumber.Number)
	}

	var categories []Category
	// .Error -> retorna o erro caso ocorra
	// Model(&Category{}) -> indica que a query é para a tabela de categorias
	// Preload("Products.SerialNumber") -> indica que queremos trazer os produtos e os números de série de cada produto
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	// --- HAS MANY
	category = Category{Name: "Cozinha"}
	db.Create(&category)

	db.Create(&Products{
		Name:       "Panela",
		Price:      320,
		CategoryID: category.ID,
	})

	for _, c := range categories {
		fmt.Println("*** ", c.Name)
		for _, p := range c.Products {
			fmt.Println("-- ", p.Name, p.Price, p.Category.Name, p.SerialNumber.Number)
		}
	}

	// --- MANY TO MANY
	group1 := Group{
		Name: "Group 1"}

	db.Create(&group1)

	group2 := Group{
		Name: "Group 2",
	}

	db.Create(&group2)

	db.Create(&Products{
		Name:       "Celular 1",
		Price:      4000,
		CategoryID: category.ID,
		Group:      []Group{group1, group2},
	})

	err = db.Model(&Category{}).Preload("Products.SerialNumber").Preload("Products.Group").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println("*** ", c.Name)
		for _, p := range c.Products {
			for _, g := range p.Group {
			fmt.Println("---- ", g.Name)
			
			fmt.Println("-- ", p.Name, p.Price, p.Category.Name, p.SerialNumber.Number, g.Name)
			}
		}
	}

}
