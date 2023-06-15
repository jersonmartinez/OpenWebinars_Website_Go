package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "¡Hola, OpenWebinars!")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Servidor está corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
