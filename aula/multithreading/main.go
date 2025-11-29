package main

import (
	"fmt"
	"sync"
	"time"
)

/*
/ quando trablhamos com threads em go, usamos Go Routines
/ uma go routine é uma função que é executada de forma assíncrona
/ para criar uma go routine, usamos a palavra go antes da chamada da função
/ uma go routine é executada em uma thread separada, mas compartilha o mesmo espaço de memória com a thread principal

/ para sincronizar o acesso a recursos compartilhados, usamos Mutexes
/ um Mutex é um mecanismo de sincronização que permite que apenas uma go routine acesse um recurso compartilhado por vez
/ para criar um Mutex, usamos a função sync.NewMutex()
/ para bloquear o acesso a um recurso compartilhado, usamos o método Lock() do Mutex
/ para desbloquear o acesso a um recurso compartilhado, usamos o método Unlock() do Mutex
/ quando uma go routine chama o método Lock() de um Mutex, ela bloqueia até que o Mutex seja desbloqueado
/ se uma go routine chama o método Lock() de um Mutex que já está bloqueado por outra go routine, ela também bloqueia

/ isso evita que duas go routines acessem o recurso compartilhado ao mesmo tempo, o que poderia causar condições de corrida
/ uma condição de corrida ocorre quando duas go routines tentam acessar o mesmo recurso ao mesmo tempo, e o resultado final depende da ordem em que as go routines são executadas
/ para evitar condições de corrida, usamos Mutexes para sincronizar o acesso a recursos compartilhados
/ um Mutex pode ser usado para proteger qualquer recurso compartilhado, como uma variável, um mapa ou um arquivo
/ quando um Mutex é usado para proteger um recurso compartilhado, apenas uma go routine pode acessar o recurso por vez
/ isso garante que o recurso seja acessado de forma segura, sem condições de corrida

/ além de Mutexes, go também fornece outros mecanismos de sincronização, como WaitGroups e Channels
/ um WaitGroup é um mecanismo de sincronização que permite que uma go routine espere por um grupo de go routines terminarem
/ para criar um WaitGroup, usamos a função sync.NewWaitGroup()
/ para adicionar uma go routine ao grupo, usamos o método Add() do WaitGroup
/ para esperar por todas as go routines do grupo terminarem, usamos o método Wait() do WaitGroup
/ um Channel é um mecanismo de comunicação que permite que go routines troquem dados de forma segura

/ um Channel tem um tipo, que define o tipo de dados que pode ser enviado e recebido pelo Channel
/ para criar um Channel, usamos a função make() com a sintaxe make(chan tipo)
/ para enviar dados para um Channel, usamos o operador <- com a sintaxe canal <- dados
/ para receber dados de um Channel, usamos o operador <- com a sintaxe dados <- canal
/ quando uma go routine envia dados para um Channel, ela bloqueia até que outra go routine receba os dados do Channel
*/

func task(name string, waitgroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Tasks %s is running\n", i, name)
		// somente para simular uma tarefa demorada
		time.Sleep(1 * time.Second)
		// indica que a tarefa foi concluída
		// o .Done() consome uma operação do waitgroup
		waitgroup.Done()
	}
}

// go routine -> Thread 1
func main() {
	waitgroup := sync.WaitGroup{}
	
	// numero total de operações a serem executadas
	waitgroup.Add(50)
	// Sem utiliza threads
	// passamos por ponteiro para que não seja criada uma cópia do waitgroup
	task("A - sem thread", &waitgroup)
	task("B - sem thread", &waitgroup)

	fmt.Println("\n\n--------------------------------------\n\n")

	// green threads 1
	go task("A - com thread", &waitgroup)
	// green threads 2
	go task("B - com thread", &waitgroup)

	// go tambem permite função anonimas
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Tasks %s is running\n", i, "C - com thread")
			time.Sleep(1 * time.Second)
			waitgroup.Done()
		}
	}()

	// espera todas as go routines terminarem
	// o total de .Add() for menor, o programa irá gerar um erro
	waitgroup.Wait()

	// automaticamente, a go routine principal termina quando o programa termina
	// para evitar que isso aconteça, podemos usar um time.Sleep() para esperar que as go routines terminem
	// utilizado somente para fins didáticos, sem o waitgroup
	//time.Sleep(30 * time.Second)
}
