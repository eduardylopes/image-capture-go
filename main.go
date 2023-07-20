package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/eduardylopes/image-generator-go/api"
)

func main() {
	http.HandleFunc("/capture", handler.Handler)

	port := "8080"
	fmt.Printf("Server listening on http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
