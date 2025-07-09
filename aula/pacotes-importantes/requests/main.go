package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	rqts, err := http.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}

	// fecha a conexão
	// para não ter vazamento de memória
	// vazamento de memória é quando o programa não libera a memória alocada
	defer rqts.Body.Close() 

	results, err := io.ReadAll(rqts.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(results))
}
