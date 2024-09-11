package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {

	ip := getEgressIP()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	message, ok := os.LookupEnv("MESSAGE")
	if !ok {
		message = fmt.Sprintf("Hello World! from: %s\n", ip)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(message))
	})

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func getEgressIP() string {
	conn, _ := net.Dial("udp", "1.2.3.4:123")
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}
