package main

type MyNumber int

type Number interface {
	// til force to accept other custom types like MyNumber
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func Compare[T comparable](a, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"a": 1, "b": 2}
	m2 := map[string]float64{"a": 1.2, "b": 2.2}
	m3 := map[string]MyNumber{"a": 10, "b": 20}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compare(1, 12334.0))
}
