package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("http://gilbertosantos.com")
	if err != nil {
		panic(err)
	}
	// Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
	defer request.Body.Close()

	res, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))

}
