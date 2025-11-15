package main

import (
	"fmt"

	pacoteinterno "banana/pacote_interno" // apartir de onde está o go.mod

	"github.com/google/uuid"
)

func main() {
	// como está no mesmo pacote, é possivel acessar a função sum
	// mesmo estando em minusculo (privado)
	// por estar no mesmo pacote, ao dar go run main.go, ele compila apenas o main.go
	// go run . -> compila todos os arquivos do pacote
	// go build . -> compila todos os arquivos do pacote e gera o executavel
	sum := sum(1, 2)

	// não é possivel acessar o pacote interno, necessita do go mod
	// go mod init <repo_git> -> cria o go.mod
	// <repo_git> é o nome do repositório no git e serve para identificar o módulo e poder ser importado
	// go mod tidy -> atualiza as dependencias
	// go mod é o gerenciador de dependencias do go
	// go get -u <nome_pacote> -> atualiza o pacote
	// go get -u ./... -> atualiza todos os pacotes
	// go get -u <nome_pacote>@<versao> -> atualiza o pacote para a versão especificada
	// go get -u <nome_pacote>@latest -> atualiza o pacote para a ultima versão
	// go get <> -> instala o pacote
	sub := pacoteinterno.Subtrai(2, 1)

	fmt.Println(sum)

	fmt.Println(sub)

	fmt.Println(uuid.New().String())
}
