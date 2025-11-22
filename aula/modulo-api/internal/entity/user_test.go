package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jhon Doe", "jd@gmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user.Name)
	assert.NotNil(t, user.Email)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jhon Doe", user.Name)
	assert.Equal(t, "jd@gmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	senha := "123456"
	user, err := NewUser("Jhon Doe", "jd@gmail.com", senha)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(senha))
	assert.False(t, user.ValidatePassword("654321"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.False(t, user.ValidatePassword("12345"))
	assert.NotEqual(t, senha, user.Password)
}
