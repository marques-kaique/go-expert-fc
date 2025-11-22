package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/dto"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/entity"
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/infra/database"
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
		return
	}

	w.WriteHeader(http.StatusCreated)
}
