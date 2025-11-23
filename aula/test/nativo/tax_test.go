package tax

import (
	"testing"
)

// Para rodar o teste, execute o comando go test .
// ou go test -v
// para saber o nivel de cobertura, execute o comando go test -cover
// ou, para gerar um relatório de cobertura, execute o comando go test -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html

func TestCalculateTax(t *testing.T) { // padrão de nomenclatura: Test<Nome da função a ser testada><Cenário>
	amount := 500.0
	expectedTax := 5.0

	tax := CalculateTax(amount)

	if tax != expectedTax {
		t.Errorf("Expected tax of %f, but got %f", expectedTax, tax)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount      float64
		expectedTax float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1500.0, 10.0},
		{2000.0, 10.0},
		{100.0, 5.0},
	}

	for _, test := range table {
		tax := CalculateTax(test.amount)

		if tax != test.expectedTax {
			t.Errorf("Expected tax of %f for amount %f, but got %f", test.expectedTax, test.amount, tax)
		}
	}
}

// Para rodar o benchmark, execute o comando go test -bench=.
// para rodar apenas o benchmark, execute o comando go test -run=^$ -bench=.
// essa expressão regular ^$ é usada para rodar todos os testes que possuem $ no nome, como não há nenhum teste com $ no nome, nenhum teste será executado
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

// o motivo de duplicar é para simular um cenário aonde há dois algoritmos diferentes que podem ser usados
// assim, podemos comparar o desempenho deles e escolher o melhor
func BenchmarkCalculateTaxWithSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTaxWithSleep(500.0)
	}
}

// é possivel utilizar p -count = 10 para rodar o teste 10 vezes e obter uma média de desempenho
// go test -run=^$ -bench=. -count=10

// tem como utilizar benchtime para definir o tempo de execução do benchmark
// go test -run=^$ -bench=. -count=10 -benchtime=3s
// se passar desse tempo, o benchmark será interrompido

// para saber a respeito de alocação de memória, execute o comando -benchmem
// go test -run=^$ -bench=. -count=10 -benchtime=3s -benchmem

// Fuzz é uma função que permite gerar dados de entrada aleatórios para testes de fuzzing
func FuzzCalculateTax(f *testing.F) {
	// seed é como se desse um exemplo de dado de entrada para o fuzzing
	seed := []float64{500.0, 1500.0, 2000.0, 100.0, 0.0, -100.0}

	for _, amount := range seed {
		// Add é usado para adicionar dados de entrada para o fuzzing
		f.Add(amount)
	}

	// o amount é o dado de entrada gerado no seed
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0.0 {
			t.Errorf("Reveid %f but expected 0", result)
		}

		if amount > 20000 && result != 20.0 {
			t.Errorf("Reveid %f but expected 20", result)
		}
	})
}
// para rodar o fuzzing, execute o comando go test -fuzz=. -run=ˆ$
// quando um valor de entrada gera um erro, o fuzzing é interrompido
// cria uma pasta chamada testdata/test/fuzz
// e no console, aparece o comando para rodar o fuzzing novamente com o valor que quebrou o teste após a correção
// para esse caso, go test -run=FuzzCalculateTax/d76b4baeef1887a8
