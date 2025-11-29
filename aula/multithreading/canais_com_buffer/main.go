package main

func main() {
	ch := make(chan string, 2) // canal com buffer de tamanho 2, não é muito recomendado usar canais com buffer muito grandes
	ch <- "Hello,"
	ch <- "world!"

	println(<-ch)
	println(<-ch)

}
