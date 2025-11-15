package main

import (
	"fmt"
	"unsafe"
)

type ID int

type User struct {
	ID   int
	Name string
	Active bool
}

var (
	a string
	b bool
	c int
	d float64
	e ID
	f User
	g [2]int        // array valor fixo
	h []int 		// slice valor variável
	i map[string]int
)

func main() {
	println("string: ", a, "tamanho: ", len(a))
	println("bool: ", b, "tamanho: ", unsafe.Sizeof(b))
	println("int: ", c, "tamanho: ", unsafe.Sizeof(c))

	println()
	println("float64: ", d, "tamanho: ", unsafe.Sizeof(d))
	fmt.Printf("%T  tamanho: %v\n", d, unsafe.Sizeof(d))
	fmt.Printf("%T  tamanho: %d\n\n", d, unsafe.Sizeof(d))

	println("\n\n*** ID - Novo tipo")
	println("ID: ", e, "tamanho: ", unsafe.Sizeof(e))

	print("\n\n*** Struct\n")
	fmt.Printf("%T: %v tamanho: %v\n", f, f, unsafe.Sizeof(f))
	fmt.Printf("%T: %v tamanho: %v\n\n", f.ID, f.ID, unsafe.Sizeof(f.ID))

	//inicializando um struct
	user := User{
		ID:   1,
		Name: "John",
		Active: true,
	}

	print("*** Struct inicializada\n")
	fmt.Printf("%T: %v tamanho: %v\n", user, user, unsafe.Sizeof(user))

	user.Active = false
	print("*** Struct alterada\n")
	fmt.Printf("%T: %v tamanho: %v\n", user, user, unsafe.Sizeof(user))
	
	println("\n\n*** Array")
	fmt.Printf("%T: %v tamanho: %v\n\n", g, g, unsafe.Sizeof(g))

	println("\n\n*** Slice")
	fmt.Printf("%T: %v tamanho: %v\n\n", h, h, unsafe.Sizeof(h))

	println("\n\n*** Map")
	fmt.Printf("%T: %v tamanho: %v\n", i, i, unsafe.Sizeof(i))

	i = make(map[string]int)
	// pode ser inicializado assim sem ser var
	// i := map[string]int{}
	// ou assim
	// i := make(map[string]int)

	i["a"] = 1
	i["b"] = 2
	println("*** Map precisa ser inicializado antes de ser usado")
	fmt.Printf("%T: %v tamanho: %v\n", i, i, unsafe.Sizeof(i))
	fmt.Printf("%T: %v tamanho: %v\n", i["a"], i["a"], unsafe.Sizeof(i["a"]))
}
