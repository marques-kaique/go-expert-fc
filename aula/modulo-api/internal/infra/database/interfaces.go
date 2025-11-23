package database

import "github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/entity"

// Isso desacopla a implementação do banco de dados da aplicação
// e permite que você substitua facilmente o banco de dados por um mock
// ou outro banco de dados sem alterar o código da aplicação.
type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]*entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
