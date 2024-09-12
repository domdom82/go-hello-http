package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	ip, ok := os.LookupEnv("POD_IP")
	if !ok {
		ip = "<POD_IP undefined>"
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	message, ok := os.LookupEnv("MESSAGE")
	if !ok {
		message = fmt.Sprintf("Hello World!\nfrom: %s\n", ip)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(message))
	})

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
