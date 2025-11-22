package tax

import (
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

// Como implementou o metodo SaveTax, entÃo, esse metodo é do tipo Repository declarado no arquivo tax.go
func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax) // registra no mock que o método foi chamado com o argumento tax
	return args.Error(0)  // retorna o erro registrado no mock
}
