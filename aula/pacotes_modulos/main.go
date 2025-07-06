package pacotesmodulos

import (
	"fmt"
	"github.com/marques-kaique/go-expert-fc/aula/pacotes_modulos/pacote_interno/math"
)

func main() {
	sum := Soma(1, 2)

	// não é possivel acessar o pacote interno, necessita do go mod
	// go mod init <repo_git> -> cria o go.mod
	// <repo_git> é o nome do repositório no git e serve para identificar o módulo e poder ser importado
	// go mod tidy -> atualiza as dependencias
	// go mod é o gerenciador de dependencias do go
	sub := math.Subtrai(2, 1)

	fmt.Println(sum)
	fmt.Println(sub)
}
