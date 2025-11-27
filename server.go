package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	fmt.Println("Server is running on", "http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
