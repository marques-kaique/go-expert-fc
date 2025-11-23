package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Para rodar o teste, execute o comando go test .
// ou go test -v
// para saber o nivel de cobertura, execute o comando go test -cover
// ou, para gerar um relatório de cobertura, execute o comando go test -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html
func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 5.0, tax)

	tax, err = CalculateTax(0.0)
	assert.Error(t, err, "amount must be greater than zero")
	assert.Equal(t, 0.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	// .On é usado para definir o comportamento esperado do mock
	// .On("SaveTax", 10.0) indica que o método SaveTax deve ser chamado com o argumento 10.0
	// .Return(nil) indica que o método deve retornar nil
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	err := CalculateTaxAndSave(10000, repository)
	// tem que passar o t pois se referencia ao teste
	assert.Nil(t, err)
 	err = CalculateTaxAndSave(0, repository)
	assert.Error(t, err, "error saving tax")
	// .AssertExpectations(t) é usado para verificar se todas as expectativas foram atendidas
	// ou seja, se todos os métodos esperados foram chamados com os argumentos esperados
	// e se o número de chamadas corresponde ao número de chamadas esperadas
	repository.AssertExpectations(t)

	// .AssertNumberOfCalls(t, "SaveTax", 1) é usado para verificar o número de chamadas de um método
	// no caso, se tivesse um if, o SaveTax poderia não ser chamado, assim voce valida se a regra esta condizente com o teste
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
