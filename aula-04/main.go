package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))

	valor, erro := sumAndPersist(51, 2)
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println(valor)
	}

	receiveArray(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	total := func() int {
		println("I'm an anonymous function called")
		v, _ := sum(1, 2)
		return v * 2
	}()

	fmt.Println(total)
}

func sum(a, b int) (int, bool) {
	if a+b > 10 {
		return a + b, true
	}
	return a + b, false
}

func sumAndPersist(a, b int) (int, error) {
	if a+b > 10 {
		return a + b, errors.New("sum is greater than 10")
	}
	return a + b, nil
}

func receiveArray(arr ...int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}
