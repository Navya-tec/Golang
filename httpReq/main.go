package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func check(w http.ResponseWriter, r *http.Request) {

	id := extractID(r.URL.Path)

	fmt.Fprintf(w, "Health check successful for ID: %s\n", id)
}

func extractID(path string) string {
	path = strings.Trim(path, "/")

	parts := strings.Split(path, "/")

	if len(parts) > 0 {
		return parts[0]
	}

	return ""
}

func main() {

	http.HandleFunc("/{id}/health", check)
	fmt.Println("Server is running...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
