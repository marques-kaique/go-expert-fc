package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product.Name)
	assert.NotNil(t, product.Price)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreateAt)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 100.0, product.Price)
}

func TestProduct_WhenNaimeIsRequired(t *testing.T) {
	product, err := NewProduct("", 100)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProduct_WhenPriceIsRequired(t *testing.T) {
product, err := NewProduct("Product 1", 0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProduct_WhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -100)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}
