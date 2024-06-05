package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	println("Busca Cep")
	for _, arg := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + arg + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao buscar cep %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta %v\n", err)
		}
		var data Cep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao decodificar json %v\n", err)
		}
		fmt.Printf("Cep: %s\nLogradouro: %s\nComplemento: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n",
			data.Cep,
			data.Logradouro,
			data.Complemento,
			data.Bairro,
			data.Localidade,
			data.Uf,
			data.Ibge, data.Gia, data.Ddd, data.Siafi)
	}
}
