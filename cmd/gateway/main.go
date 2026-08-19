package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Resilient API Gateway listening on :8080...")
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
}
