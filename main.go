package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
