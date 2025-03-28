package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Estrutura comum para armazenar o endereço (adaptada das duas APIs)
type Endereco struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"` // viacep
	Uf         string `json:"uf"`
	Api        string `json:"api"` // custom field para saber qual API respondeu
}

func main() {
	cep := "01153000"
	timeout := 1 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resultChan := make(chan Endereco, 2)
	errChan := make(chan error, 2)

	go fetchFromBrasilAPI(ctx, cep, resultChan, errChan)
	go fetchFromViaCEP(ctx, cep, resultChan, errChan)

	select {
	case res := <-resultChan:
		fmt.Printf("Resposta recebida da API %s:\n", res.Api)
		fmt.Printf("CEP: %s\nLogradouro: %s\nBairro: %s\nCidade: %s\nEstado: %s\n",
			res.Cep, res.Logradouro, res.Bairro, res.Localidade, res.Uf)
	case <-ctx.Done():
		fmt.Println("Erro: Timeout de 1 segundo atingido.")
	}
}

func fetchFromBrasilAPI(ctx context.Context, cep string, resultChan chan<- Endereco, errChan chan<- error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errChan <- fmt.Errorf("BrasilAPI retornou status %d", resp.StatusCode)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var data struct {
		Cep          string `json:"cep"`
		Street       string `json:"street"`
		Neighborhood string `json:"neighborhood"`
		City         string `json:"city"`
		State        string `json:"state"`
	}
	json.Unmarshal(body, &data)

	resultChan <- Endereco{
		Cep:        data.Cep,
		Logradouro: data.Street,
		Bairro:     data.Neighborhood,
		Localidade: data.City,
		Uf:         data.State,
		Api:        "BrasilAPI",
	}
}

func fetchFromViaCEP(ctx context.Context, cep string, resultChan chan<- Endereco, errChan chan<- error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errChan <- fmt.Errorf("ViaCEP retornou status %d", resp.StatusCode)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var data struct {
		Cep        string `json:"cep"`
		Logradouro string `json:"logradouro"`
		Bairro     string `json:"bairro"`
		Localidade string `json:"localidade"`
		Uf         string `json:"uf"`
	}
	json.Unmarshal(body, &data)

	resultChan <- Endereco{
		Cep:        data.Cep,
		Logradouro: data.Logradouro,
		Bairro:     data.Bairro,
		Localidade: data.Localidade,
		Uf:         data.Uf,
		Api:        "ViaCEP",
	}
}
