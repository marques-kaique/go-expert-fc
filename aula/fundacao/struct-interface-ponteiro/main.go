package main

import (
	"fmt"
	"unsafe"
)

var (
	e ID
	f User
)

type ID int

type Andress struct {
	Street string
	Number int
}

type People interface { // interface é um contrato
	Ativar() // somente assinatura do método
			 // precisa ser implementado indentico do contrato
}

type User struct {
	ID     int
	Name   string
	Active bool
	Andress Andress // composição - um struct dentro de outro
	Date interface{} // qualquer tipo
}

func NewUser() *User {
	return &User{Name: "John"}
}

func (u User) Ativar() {
	u.Active = true
}

func (u *User) Ativacao() {
	u.Active = true
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Active: %v, Andress: %v", u.ID, u.Name, u.Active, u.Andress)
}

func alterandoValores(a, c *int)  {
	*a = 1
	*c = 2
}

func main() {
	println("\n\n*** ID - Novo tipo")
	println("ID: ", e, "tamanho: ", unsafe.Sizeof(e))

	print("\n\n*** Struct\n")
	fmt.Printf("%T: %v tamanho: %v\n", f, f, unsafe.Sizeof(f))
	fmt.Printf("%T: %v tamanho: %v\n\n", f.ID, f.ID, unsafe.Sizeof(f.ID))

	//inicializando um struct
	user := User{
		ID:     1,
		Name:   "John",
		Active: true,
		Andress: Andress{
			Street: "Main Street",
			Number: 123,
		},
	}

	print("*** Struct inicializada\n")
	fmt.Printf("%T: %v tamanho: %v\n", user, user, unsafe.Sizeof(user))

	user.Active = false
	user.Andress.Number = 456
	println("*** Struct alterada\n")
	fmt.Printf("%T: %v tamanho: %v\n", user, user, unsafe.Sizeof(user))

	print("\n*** Struct com metodos\n")
	user.Ativar()
	print(user.String())

	// quando um struct implementa uma interface, ele herda todos os métodos da interface
	People.Ativar(user)

	// ponteiro
	println("\n\n*** Ponteiro\n")

	userPonteiro := &user

	println("Posição memoria user: %v", &user)
	println("Posição memoria userPonteiro: %v", userPonteiro)

	//* acesa o valor do ponteiro
	// & acessa o endereço de memoria
	*&userPonteiro.Name = "Doe" 

	println("Name: ", user.Name)

	// exmplo de ponteiro simples
	a := 1
	b := &a
	*b = 2
	println("\na: ", a)
	println("b: ", *b)

	var c int 

	// passando o endereço de memoria
	alterandoValores(&a, &c)

	// como o valor de 'a' foi alterado, o valor de 'b' também é alterado
	println("\na: ", a)
	println("b: ", *b)
	println("c: ", c)

	// ativar sem ponteiro
	print("\n*** Ativar sem ponteiro\n")
	user.Ativar()
	fmt.Printf("%T: %v tamanho: %v\n", user, user, unsafe.Sizeof(user))

	// ativar com ponteiro
	print("\n*** Ativar com ponteiro\n")
	user.Ativacao()
	fmt.Printf("%T: %v tamanho: %v\n", user, user, unsafe.Sizeof(user))

	// nova instancia de user via ponteiro
	// parece construtor de java
	// Também é utilizado para Singleton
	print("\n*** Nova instancia de user via ponteiro\n")
	user3 := NewUser()
	fmt.Printf("%T: %v tamanho: %v\n", user3, user3, unsafe.Sizeof(user3))

	// interface vazia 
	// pode receber qualquer tipo
	// é utilizada quando não sabemos o tipo que será recebido

	println("\n\n*** Interface vazia\n")
	user.Date = "2024-01-01"
	fmt.Printf("%T: %v tamanho: %v\n", user.Date, user.Date, unsafe.Sizeof(user.Date))

	user3.Date = 2024
	fmt.Printf("%T: %v tamanho: %v\n\n", user3.Date, user3.Date, unsafe.Sizeof(user3.Date))	

	// como não sabemos o tipo que será recebido, é necessário fazer um type assertion
	// para converter o tipo da interface para o tipo desejado
	
	println("Valor do user3.Date:",  user3.Date)
	println(user3.Date.(int)) // type assertion
	println()
	// type assertion quando não sabemos o tipo
	// ok é um booleano que indica se a conversão foi bem-sucedida ou não
	result, ok := user3.Date.(string) 
	
	if !ok {
		println(ok)
		println("Não é uma string")
	}

	println("Valor do result:",  result)
}
