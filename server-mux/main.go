package main

import "net/http"

func main() {
	println("Starting server on port 8081")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})
	mux.Handle("/blog", blog{})
	http.ListenAndServe(":8081", mux)
}

type blog struct{}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Blog"))
}
