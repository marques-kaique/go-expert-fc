package main

import "fmt"

func recebe(nome string, hello chan<- string) {// chan<- indica que o canal é somente de enviar, encher o canal
	hello <- nome
}

func ler(data <-chan string) { // <-chan indica que o canal é somente de esvaziar, retirar valores do canal
	fmt.Println(<-data) // <- remove o valor do canal
}


func main() {
	hello := make(chan string)
	go recebe("hello", hello)
	ler(hello)
}
