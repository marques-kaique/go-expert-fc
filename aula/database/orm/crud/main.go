package main

// https://gorm.io/
// uma coisa que precisa tomar cuidado é com  Eager loading with Preload, Joins
// quando for fazer um select, sempre utilizar o select, para evitar que o gorm faça um select em todas as colunas

// orm quando é utilizado para um volume muito grande de dados, ele pode ser mais lento que o sql puro
// pois o orm faz muitas coisas automaticamente

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm utiliza o struct tag para mapear o campo do struct com a coluna do banco de dados
type Products struct {
	//ID    int `gorm:"primaryKey"` // agora gerenciado pelo gorm
	Name  string
	Price float64
	gorm.Model // cria os campos created_at, updated_at, deleted_at, ID
}

func main() {
	// é possivel definir utf, timezone, locale, etc
	dsn := "root:123456@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// auto migrate é utilizado para criar as tabelas no banco de dados
	// isso ocorre apenas se a tabela não existir
	// principalmente utilizado em testes
	db.AutoMigrate(&Products{})

	db.Create(&Products{
		Name:  "Test",
		Price: 1000.0,
	})

	// create em batch
	products := []Products{
		{Name: "Test", Price: 1000.0},
		{Name: "Test2", Price: 2000.0},
		{Name: "Test3", Price: 3000.0},
	}

	db.Create(&products)

	// select
	var product Products
	// firts é o primeiro registro que ele encontrar para a condição
	// se não passar a condição, ele retorna o primeiro registro da tabela
	db.First(&product, 1) // find product with id 1
	fmt.Println(product)

	db.First(&product, "name = ?", "Test") // find product with name Test

	// find all products
	var allProducts []Products
	db.Find(&allProducts)
	// para limitar a quantidade de registros, pode ser utilizado o limit
	// db.Limit(2).Find(&allProducts)
	// é possivel utilizar o offset para paginação
	// db.Offset(2).Limit(2).Find(&allProducts)

	for _, p := range allProducts {
		fmt.Printf("All: Product: %v, possui o preço de %.2f\n", p.ID, p.Price)
	}

	// where
	db.Where("price < ?", 2000).Find(&allProducts) // find all products with price less than 2000
	for _, p := range allProducts {
		fmt.Printf("Where: Product: %v, possui o preço de %.2f\n", p.ID, p.Price)
	}

	// update
	//db.Model(&product).Update("Price", 9999)
	product.Price = 9999
	db.Save(&product)

	// like
	db.Where("name LIKE ?", "%Test%").Find(&allProducts) // find all products with name containing Test
	for _, p := range allProducts {
		fmt.Printf("Like: Product: %v, possui o preço de %.2f\n", p.ID, p.Price)
	}

	db.Delete(&product) // soft delete
	//db.Unscoped().Delete(&product) // hard delete
}
