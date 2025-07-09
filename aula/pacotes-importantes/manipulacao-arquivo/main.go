package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./file.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close() // defer -> executa no final da função

	// assim escreve no arquivo apenas string
	//tamanho, err := f.WriteString("Hello, World!")

	// dessa forma escreve qualquer coisa no arquivo
	// forma mais eficiente
	tamanho, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("A quantidade de bytes escritos foi: %d bytes\n", tamanho)


	// para ler o arquivo

	// abre o arquivo, porem precisa fazer validação de erro
	//file, err := os.Open("./file.txt") 
	
	// abre o arquivo em modo somente leitura
	file, err := os.ReadFile("./file.txt")

	if err != nil {
		panic(err)
	}

	// se utilizar o os.ReadFile, não é necessario fechar o arquivo

	// converte o slice de bytes para string
	fmt.Println(string(file)) 

	// para carregamentos parcial de arquivos, se utiliza o os.Open
	// se utiliza o os.Open para arquivos grandes, pois o os.ReadFile carrega o arquivo inteiro na memória
	// isso é importante para arquivos grandes, pois pode estourar a memória
	fmt.Println("\nCarregamento parcial de arquivos")
	file_parcial, err := os.Open("./file.txt")
	if err != nil {
		panic(err)
	}

	defer file_parcial.Close() 

	reader := bufio.NewReader(file_parcial) // cria um leitor de arquivos

	buffer := make([]byte, 4) // cria um buffer de 4 bytes


	for	{
		n, err := reader.Read(buffer) // lê 4 bytes do arquivo
		if err != nil { // erro ocorre quando chega no final do arquivo
			break
		}

		fmt.Println(string(buffer[:n])) // converte o slice de bytes para string e imprime
	}

	// para remover um arquivo
	 err = os.Remove("./file.txt")
	 if err != nil {
		 panic(err)
	 }
}

// para abrir um arquivo pelo terminal se utiliza o comando cat
// cat <nome_arquivo> -> abre o arquivo
// cat <nome_arquivo> | wc -l -> conta as linhas do arquivo
// cat <nome_arquivo> | wc -w -> conta as palavras do arquivo
// cat <nome_arquivo> | wc -c -> conta os caracteres do arquivo

// parar criar um arquivo pelo terminal se utiliza o comando mkdir
// mkdir <nome_arquivo> -> cria o arquivo
// touch <nome_arquivo> -> cria o arquivo
// echo "Hello, World!" > <nome_arquivo> -> escreve no arquivo
// echo "Hello, World!" >> <nome_arquivo> -> escreve no arquivo sem sobrescrever
