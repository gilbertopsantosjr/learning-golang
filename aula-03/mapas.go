package main

import "fmt"

// can have only one main function per package ?
func my_maps() {
	salary := map[string]int{"Gilberto": 1000, "Davi": 2000, "Jefte": 3000, "Pedro": 4000, "Isaac": 5000}
	fmt.Println("my maps is", salary)

	delete(salary, "Gilberto")
	fmt.Println("my maps is", salary)

	salary["Gilberto"] = 1500
	fmt.Println("update maps is", salary)

	sal := make(map[string]int)
	sal["Gilberto"] = 1000
	sal["Davi"] = 2000
	sal["Jefte"] = 3000
	sal["Pedro"] = 4000
	sal["Isaac"] = 5000

	for key, value := range sal {
		fmt.Printf("the salario of %s is value $%v \n", key, value)
	}

	for _, value := range sal {
		fmt.Printf("the salario value $%v \n", value)
	}
}
