package main

var (
	a bool    = true
	b int     = 1
	c float64 = 1.1
	e string  = "gilberto"
)

func main() {
	// we need to run all files withing the same package to automatically import the constants
	// go run *.go  or go run main.go env.go
	println("the value of a:", a)
	println("the value of b:", b)
	println("the value of c:", c)
	println("the value of e:", e)

	user := "Gilberto" //auto infered type

	println("the value of user:", user)
}
