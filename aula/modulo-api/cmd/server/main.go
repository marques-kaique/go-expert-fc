package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/configs"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/entity"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/infra/database"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)

	productHandler := handlers.NewProductHandler(productDB)

	//implementação default do go
	/*
		http.HandleFunc("/products", productHandler.CreateProduct)

	   	http.ListenAndServe(":8000", nil) */

	// implementação com chi
	r := chi.NewRouter()
	// middleware é uma função que intercepta a requisição antes de chegar no handler
	// pode ser usado para log, autenticação, etc
	// o chi já tem um middleware de log
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Get("/products/", productHandler.GetProducts)

	http.ListenAndServe(":8000", r)

	// roteadores permitem atachar roteamentos das requições ate chegar no handler
	// isso significa que o handler não precisa saber qual é a rota
	// roteador do GO não permite trabalhar com variaveis
	// de agrupar rotas
	// de usar regex
}
