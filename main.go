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

	message, ok := os.LookupEnv("MESSAGE")
	if !ok {
		message = "Hello World!"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(message))
	})

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
