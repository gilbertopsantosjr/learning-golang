package main

// * points to the memory address of the variable
// it will change the value of the variable in memory
func sum(a *int, b *int) int {
	tax := 52
	*a = *a - tax
	*b = *b - tax
	return *a + *b
}
