package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Escrevendo dados no arquivo!"))
	fmt.Printf("Tamanho: %d bytes\n", size)

	f.Close()

	file, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	arquivo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err2 := os.Remove("arquivo.txt")
	if err2 != nil {
		panic(err2)
	}
}
