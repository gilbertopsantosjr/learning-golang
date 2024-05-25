// any file needs a package
// the package is the name of the folder where the file is located
// the main package is the entry point of the program
package main

func main() {
	// we need to run all files withing the same package to automatically import the constants
	// go run *.go  or go run main.go env.go
	println("the value of a:", a)
	println("the value of b:", b)
}
