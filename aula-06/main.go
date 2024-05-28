package main

import "fmt"

// empty interface

func main() {

	var x interface{} = 10
	var y interface{} = "João"

	show(x)
	show(y)

	show(y.(string)) // force cast to string
	res, ok := y.(string)
	if ok {
		fmt.Println(res)
	} else {
		fmt.Println("Error in cast")
	}
}

func show(a interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", a, a)
}
