package main

type Endereco struct {
	Logradouro string
	Numero     int
	Bairro     string
	Cidade     string
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Pessoa interface {
	Deactivate()
}

// star * helps to change the value of the struct using point memory
func (c *Client) Deactivate() {
	c.Ativo = false
}

func NewClient() *Client {
	return &Client{}
}

// how to add func to struct
func (c Client) show() {
	println(c.Nome)
	println(c.Idade)
	println(c.Ativo)
	println(c.Endereco.Logradouro)
	println(c.Endereco.Numero)
	println(c.Endereco.Bairro)
	println(c.Endereco.Cidade)
}

func main() {

	cliente := Client{
		Nome:     "João",
		Idade:    25,
		Ativo:    true,
		Endereco: Endereco{"Rua 1", 10, "Centro", "São Paulo"},
	}

	println(cliente.Nome)

	cliente.show()
	cliente.Deactivate()

	// print the memory address of the struct
	println("Memory address:", &cliente.Ativo)
	println("Value in Memory address:", cliente.Ativo)

	salary := 1500
	bonus := 500
	total := sum(&salary, &bonus)

	println(total)

}
