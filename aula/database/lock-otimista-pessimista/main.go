package main

// Lock otimista
// utilizado para realizar muitas operações que pode haver concorrência
// verificar a versão antes de atualizar
// se a versão for a mesma, atualiza
// se a versão for diferente, retorna um erro

// Lock pessimista
// utilizado para realizar uma operação que pode haver concorrência
// bloqueia o registro para que ninguém mais possa acessar
// quando a operação é finalizada, o registro é desbloqueado
// em sql, é utilizado o comando select for update
// select * from products where id = ? for update

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Products struct {
	Name  string
	Price float64
	gorm.Model
}

type Category struct {
	Name string
	gorm.Model
}

func main() {
	dsn := "root:123456@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Products{}, &Category{})

	// db.Begin() inicia uma transação
	tx := db.Begin()

	var c Category
	// tx.Debug() habilita o modo de depuração
	// clause.Locking{Strength: "UPDATE"} define o tipo de bloqueio, igual ao comando select for update
	// First(&c, 1) seleciona o primeiro registro da tabela e agrega a c inicializada anteriormente
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error

	if err != nil {
		panic(err)
	}

	c.Name = "Updated Category"

	tx.Debug().Save(&c)
	tx.Commit()

}
