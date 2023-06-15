# OpenWebinars_Website_Gos

### Manejo de peticiones HTTP y respuestas

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Servidor corriendo en http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido a la página de inicio!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la página 'Acerca de nosotros'.")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Póngase en contacto con nosotros en jersonmartinez.com")
}
```

### Manejo de rutas y parámetros en Go

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/detail/", productDetailHandler)

	fmt.Println("Servidor corriendo en http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido a la página de inicio!")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aquí se muestran todos los productos disponibles")
}

func productDetailHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Path[len("/products/detail/"):]
	fmt.Fprintf(w, "Detalles del producto con ID: %s", productID)
}
```