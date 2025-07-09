package main


type MyNumber int

// isso é uma constraint
// pode ser int ou float64
// o simbolo ~ é usado para representar uma constraint
// por exemplo, pode ser usado MyNumber e ele entende que é um int
type Number interface {
	~int | float64
}

func Soma[T int | float64](a, b T) T {
	return a + b
}

func SomaMap[T Number](m map[string] T) T {
	var result T

	for _, v := range m {
		result += v
	}

	return result
}

// comparable é utilizado para comparar a igualdade de dois valores
// não é possivel utilizar para ver se é maior ou menor
func compara[T comparable](a, b T) bool {
	return a == b
}

func main() {

	println("*** Soma generics")
	somaInt := Soma(1, 2)
	somaFloat := Soma(1.5, 2.5)

	println(somaInt)
	println(somaFloat)	

	somaMapInt := SomaMap(map[string]int{"a": 1, "b": 2})
	somaMapFloat := SomaMap(map[string]float64{"a": 1.5, "b": 2.5})

	println("\n*** Soma Map generics")
	println(somaMapInt)
	println(somaMapFloat)


	// como agora está passando MyNumber
	// Na constraint Number, int precisa conter o sinal ~
	// indicando que qualquer type que seja compatível com int pode ser usado
	somaMapMyNumber := SomaMap(map[string]MyNumber{"a": 1, "b": 2})
	println(somaMapMyNumber)	


	println("\n*** Compara generics")

	println(compara(1, 1.0))
}
