package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	http.HandleFunc("/", BuscaCepHandler)

	// usando uma função anonima
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// r.URL.Query() -> retorna um map[string][]string
	// isso é um query param
	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// pega o resultado da busca (data) e transforma em json
	// o json gerado é atribuido ao corpo da resposta (w)
	// quando o json é gerado, ele é enviado para o cliente
	// se não fosse a itenção de enviar para o cliente, poderia ser atribuido a uma variavel
	// para atribuir a uma variavel, poderia ser utilizado o marshal
	json.NewEncoder(w).Encode(data)

	// atribui o json gerado a uma variavel
	// result, err := json.Marshal(data)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)
}

func BuscaCep(cep string) (*ViaCep, error) { // utilizando um ponteiro para ViaCep para retornar o endereço de memória
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	result, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data ViaCep

	err = json.Unmarshal(result, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
