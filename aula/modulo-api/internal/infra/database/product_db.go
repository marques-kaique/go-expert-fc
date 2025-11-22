package database

import (
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}

	// O findById busca o item, caso não encontre, retorna um erro
	// O save atualiza o item, caso não encontre, cria um novo
	// dessa forma, evitamos a criação de um novo item caso o id não exista
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	_, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(&entity.Product{}, "id = ?", id).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
	if sort == "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	var products []*entity.Product
	var err error
	
	if page != 0 && limit != 0 {
		offset := (page - 1) * limit
		err = p.DB.Offset(offset).Limit(limit).Order("create_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("create_at " + sort).Find(&products).Error
	}

	return products, err
}
