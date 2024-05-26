package main

import "fmt"

// create your own type

type myType int

var a myType = 1

func main() {
	fmt.Printf("the type of a: %T value %v \n", a, a)

	var my_array_inline = [3]int{1, 2, 3}
	var my_array [5]string
	my_array[0] = "Gilberto"
	my_array[1] = "Davi"
	my_array[2] = "Jefte"
	my_array[3] = "Pedro"
	my_array[4] = "Isaac"

	fmt.Printf("the type of my_array: %T value %v \n", my_array_inline, my_array_inline)
	fmt.Printf("the type of my_array: %T value %v \n", my_array, my_array)
	fmt.Println(len(my_array_inline))

	for position, value := range my_array {
		fmt.Printf("the position %d has the value %v \n", position, value)
	}

	println("=================SLICE===================")

	// slice
	my_slice := []int{10, 20, 30, 40, 50, 70, 80, 90, 100}

	fmt.Printf("len=%d capacity=%d %v\n", len(my_slice), cap(my_slice), my_slice)

	//DROP everything, but keep the capacity
	fmt.Printf("len=%d capacity=%d %v\n", len(my_slice[:0]), cap(my_slice[:0]), my_slice[:0])

	//DROP until 4 index
	fmt.Printf("len=%d capacity=%d %v\n", len(my_slice[:4]), cap(my_slice[:4]), my_slice[:4])

	//DROP from 2 index
	fmt.Printf("len=%d capacity=%d %v\n", len(my_slice[2:]), cap(my_slice[2:]), my_slice[2:])

	//DROP from 2 index until 4 index (interval)
	fmt.Printf("len=%d capacity=%d %v\n", len(my_slice[2:4]), cap(my_slice[2:4]), my_slice[2:4])

	my_slice = append(my_slice, 110, 120, 130) // for every append the capacity is doubled
	fmt.Printf("len=%d capacity=%d %v\n", len(my_slice), cap(my_slice), my_slice)

	my_maps()

}
