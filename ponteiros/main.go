package main

import "fmt"

type Carro struct {
	Name string
}

type CityDBRepository struct {
}

func NewCityDBRepository() *CityDBRepository {
	return &CityDBRepository{}
}

func (c *Carro) andou() {
	c.Name = "Andou"
	fmt.Printf("O carro %s\n", c.Name)
}

func main() {
	a := 10
	var ponteiro *int = &a // & endereco da memoria
	*ponteiro = 20         // * alterando o valor no enderecoq
	fmt.Println(*ponteiro) // * valor no endereco
	b := &a                // b aponta pro endereco de a
	*b = 30                // altera o valor de a
	fmt.Println(b)

	// a funcao altera o valor da memoria
	abc(&a)

	repo := NewCityDBRepository()

	fmt.Println("here", *repo)
}

func abc(a *int) {
	*a = 100
}
