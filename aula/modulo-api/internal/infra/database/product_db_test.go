package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	db := DB.Begin()
	defer db.Rollback()

	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func Test_FindAllProducts(t *testing.T) {
	db := DB.Begin()
	defer db.Rollback()

	productDB := NewProduct(db)

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		err = productDB.Create(product)
		assert.NoError(t, err)
	}

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func Test_FindByID(t *testing.T) {
	db := DB.Begin()
	defer db.Rollback()

	productDB := NewProduct(db)

	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	err = productDB.Create(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10.0, product.Price)
}

func Test_UpdateProduct(t *testing.T) {
	db := DB.Begin()
	defer db.Rollback()

	productDB := NewProduct(db)

	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	err = productDB.Create(product)
	assert.NoError(t, err)

	product.Name = "Product 2"
	product.Price = 20
	err = productDB.Update(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
}

func Test_DeleteProduct(t *testing.T) {
	db := DB.Begin()
	defer db.Rollback()

	productDB := NewProduct(db)

	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	err = productDB.Create(product)
	assert.NoError(t, err)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
