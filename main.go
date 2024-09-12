package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {

	ip4 := getEgressv4IP()
	ip6 := getEgressv6IP()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	message, ok := os.LookupEnv("MESSAGE")
	if !ok {
		message = fmt.Sprintf("Hello World!\nfrom:\n%s (v4)\n%s (v6)\n", ip4, ip6)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(message))
	})

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func getEgressv4IP() string {
	conn, err := net.Dial("udp", "1.2.3.4:123")
	if err != nil {
		return err.Error()
	}
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}

func getEgressv6IP() string {
	conn, err := net.Dial("udp6", "[1:2:3:4:5:6:7:8]:123")
	if err != nil {
		return err.Error()
	}
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}
