package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", okHandler)
	fmt.Println("Starting server on port:", strings.Split(getPort(), ":")[1])
	log.Fatal(http.ListenAndServe(getPort(), nil))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()

	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "Version: 2.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
	fmt.Fprintf(w, "Host: %s", r.Host)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getPort() string {
	var port = getEnv("PORT", "8080")
	return ":" + port
}
