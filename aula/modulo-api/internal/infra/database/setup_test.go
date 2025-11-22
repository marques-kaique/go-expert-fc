package database

import (
	"os"
	"testing"
	"time"

	"github.com/marques-kaique/go-expert-fc/aula/modulo-api/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// isso aqui é um hook, é executado antes de todos os testes
// igual o @BeforeAll do JUnit
// seguindo 100% o exemplo do professor
// fica muito lento para rodar os testes
//
//	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{}) -> exemplo dele, abrindo a conexão em todos os testes
//	db.AutoMigrate(entity.Product{}) -> e, era necessario recriar a tabela em todos os testes
func TestMain(m *testing.M) {
	// 🔧 1. Abrir conexão
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	// 🔧 2. Configurar pool
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 🔧 3. Criar tabelas (migração única)
	if err := db.AutoMigrate(
		&entity.Product{},
		&entity.User{},
	); err != nil {
		panic(err)
	}

	DB = db // expõe para todos os testes

	// 🧪 4. Rodar os testes
	code := m.Run()

	// 🧹 5. Cleanup final
	sqlDB.Close()

	os.Exit(code)
}
