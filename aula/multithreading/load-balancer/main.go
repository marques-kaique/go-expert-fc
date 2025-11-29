package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qtdWorks := 1000000

	// incializar os wokers
	for i := 0; i < qtdWorks; i++ {
		go worker(i, data)
	}

	go worker(2, data)

	// worker que terminar primeiro, pega o valor do canal
	for i := 0; i < 10000000; i++ {
		data <- i
	}

}
