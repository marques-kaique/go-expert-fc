package main

import (
	"fmt"
	"sync"
)

var number uint64 = 0

// Mesmo algoritmo do problema de concorrência, porém utilizando canais para resolver o problema
// sem o canal, o sistema não entende que há alguem mexendo em um valor x, e acaba alterando o valor ao mesmo tempo
// com o canal, o sistema entende que há alguem mexendo em um valor x, e bloqueia o acesso ao valor até que a operação seja concluída
// garantindo que o valor seja alterado de forma correta
// canal cheio, quando publica no canal o valor para outra thread
// canal vazio, quando não há valor no canal para ser consumido por outra thread

// thread 1
func main() {
	channel := make(chan string) // Cria um canal com buffer

	// thread 2
	go func() { 
		channel <- "Hello, world!" // Envia uma mensagem para o canal
	}()

	// thread 1
	msg := <-channel // Recebe a mensagem do canal, esvazia o canal
	fmt.Println(msg)

	forever := make(chan bool)
	
	
	go func() {
		forever <- true
	}()
	
	// Bloqueia a thread principal para que o programa não termine imediatamente
	// se remover o go func() acima, o programa irá travar aqui
	// pois para forever ficar cheio, necessita de outra thread para colocar um valor nele
	// daria para substituir o waitgroup := sync.WaitGroup{} por esse canal
	<-forever 

	ch := make(chan int)

	go publisher(ch)
	reader(ch) // não tem como ser go reader(ch), pois se for, a main thread irá terminar antes de ler os valores do canal
	// poderia ser colocado o valor do reader aqui, não mudaria nada

	// com o waitgroup, seguramos a thread principal até que todas as goroutines terminem
	// podemos assim trabalhar com varias threads sem que a main thread termine antes delas
	wg := sync.WaitGroup{}
	th := make(chan int)
	wg.Add(10)
	go publisher(th)
	go readerWithWG(th, &wg)

	wg.Wait() // espera todas as goroutines terminarem
}


func readerWithWG(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Println("Recebido:", x)
		wg.Done() // sinaliza que uma goroutine terminou
	}
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Println("Recebido:", x)
	} 
}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// fecha o canal após enviar todos os valores
	// assim o reader sabe que não há mais valores para serem lidos
	close(ch) 
}
