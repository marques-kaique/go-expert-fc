package database

import (
	"testing"

	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// file:memory -> banco de dados em memoria
	// talves seja sugerido pelo copilot incluir cache aqui, porem prejudica muito no teste
	//db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	// devido demorar muito para rodar os testes, vamos usar o setup_test.go
	db := DB.Begin()
	defer db.Rollback()

	user, _ := entity.NewUser("John", "j@j.com", "123456")
	userDB := NewUser(db)
	err := userDB.Create(user)

	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func Test_FindByEmail(t *testing.T) {
	db := DB.Begin()
	defer db.Rollback()

	user, _ := entity.NewUser("John Due", "jd@gmail.com", "123456")
	userDB := NewUser(db)
	err := userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail("jd@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
