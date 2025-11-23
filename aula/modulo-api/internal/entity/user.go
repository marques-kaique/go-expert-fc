package entity

import (
	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/pkg/entity"
	"golang.org/x/crypto/bcrypt" // realiza a criptografia da senha
)

// em DDD, são objetos de valor, pois não possuem identidade
type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"-"` // ignore password on marshalling, isso é para não retornar a senha no json
	Email    string    `json:"email"`
}

func NewUser(name, email, password string) (*User, error) {
	// gera o hash da senha
	// bcrypt.DefaultCost é o custo do poder processamento, quanto maior o custo, mais tempo leva para gerar o hash
	// e mais seguro é o hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	// compara a senha com o hash
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
