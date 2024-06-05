package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Numero int     `json:"num"`
	Saldo  float64 `json:"saldo"`
}

func main() {
	person := Person{Numero: 1, Saldo: 100.0}
	_, err := json.Marshal(person) // parse to json
	if err != nil {
		panic(err)
	}

	erro := json.NewEncoder(os.Stdout).Encode(person)
	if erro != nil {
		panic(erro)
	}

	jsonRaw := []byte(`{"num": 2, "saldo": 200.0}`)
	var p Person
	err = json.Unmarshal(jsonRaw, &p) // parse to struct
	if err != nil {
		panic(err)
	}
	fmt.Printf(" Numero: %d Saldo %v \n", p.Numero, p.Saldo)
}
