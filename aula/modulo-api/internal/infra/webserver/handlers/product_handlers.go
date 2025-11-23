package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/dto"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/entity"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/infra/database"
	entityPkg "github.com/marques-kaique/go-expert-fc/aula/modulo-api/pkg/entity"
)

// Vai receber uma instancia de database.ProductInterface
// pois pode ser injetado qualquer implementação de banco de dados
type ProductHanlder struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHanlder {
	return &ProductHanlder{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body dto.CreateProductRequest true "product data"
// @Success 201
// @Failure 400
// @Failure 500 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHanlder) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// quando recebe dados de requisição, são dados puros
	// para manipular, precisa de um DTO
	// DTO é um objeto que representa os dados que estão sendo transferidos
	var product dto.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// isso é feito em caso de uso, não no handler
	// clean architecture aborda isso
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetProduct godoc
// @Summary Get product
// @Description get one product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product id" format(uuid)
// @Success 200 {array} entity.Product
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHanlder) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println("O id é: ", id)
	if id == "" {
		fmt.Println("ID is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// ListProducts godoc
// @Summary List products
// @Description get all products
// @Tags products
// @Accept json
// @Produce json
// @Param page query string false "page number"
// @Param limit query string false "limit"
// @Param sort query string false "sort"
// @Success 200 {array} entity.Product
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHanlder) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product id" format(uuid)
// @Param product body dto.CreateProductRequest true "product data"
// @Success 200 
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHanlder) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err = h.ProductDB.FindByID(product.ID.String())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product id" format(uuid)
// @Success 200 
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHanlder) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
