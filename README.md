# OpenWebinars_Website_Go

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

### Cargar archivos estáticos JS y CSS

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("web/templates/", "home.html"))
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("web/templates/", "error.html"))
	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Agregar datos dinámicos a plantilla HTML

```go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type PageData struct {
	Title        string
	Message      template.HTML
	ErrorCode    int
	ErrorMessage string
}

func renderTemplate(w http.ResponseWriter, tmplFile string, data PageData) {
	tmpl, err := template.ParseFiles(tmplFile)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "OpenWebinars / Jerson",
		Message: template.HTML(`
			La plataforma donde encontrarás el curso <b>"Mi primera página web en Go"</b> impartido por el Ing. DevOps, Jerson Martínez. Recibido por los estudiantes más inteligentes de la plataforma."
		`),
	}

	tmplFile := filepath.Join("web/templates/", "home.html")
	renderTemplate(w, tmplFile, data)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:        "Error 404",
		ErrorCode:    500,
		ErrorMessage: "¡Página no encontrada!",
	}

	tmplFile := filepath.Join("web/templates/", "error.html")
	renderTemplate(w, tmplFile, data)
}

func main() {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/error", errorHandler)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Archivo `home.html`

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>{{ .Title }}</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet"
            integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
        <link rel="stylesheet" href="/static/css/custom.css">
    </head>
    <body>
        <main>
            <section class="py-5 text-center container">
                <div class="row py-lg-5">
                    <div class="col-lg-6 col-md-8 mx-auto">
                        <!-- <h1 class="fw-light">¡Bienvenido a OpenWebinars!</h1> -->

                        <img src="/static/images/logo-black.svg" id="logo" class="img-fluid w-50" alt="Logo OpenWebinars"/>

                        <p class="lead text-body-secondary m-5">
                            {{ .Message }}
                        </p>

                        <p>
                            <a href="#" class="btn btn-primary my-2">Ver más información</a>
                            <a href="#" class="btn btn-secondary my-2">Calificar con 5 estrellas</a>
                        </p>

                    </div>
                </div>
            </section>
        </main>

        <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"
            integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz"
            crossorigin="anonymous"></script>

        <script src="/static/js/custom.js"></script>
    </body>
</html>
```

Archivo `error.html`

```html
<!DOCTYPE html>
<html lang="en">

	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>{{ .Title }}</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet"
			integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">

	</head>

	<body>
		<main>
			<section class="py-5 text-center container">
				<div class="row py-lg-5">
					<div class="col-lg-6 col-md-8 mx-auto">
						<h1 class="display-1 fw-bold">{{ .ErrorCode }}</h1>

						<p class="fs-3"> <span class="text-danger">¡Ooops!</span> {{ .ErrorMessage }} </p>

						<p class="lead">
							La página que estás buscando no existe.
						</p>

						<a href="./" class="btn btn-primary">Ir a la página principal</a>
					</div>
				</div>
			</section>
		</main>

		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"
			integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz"
			crossorigin="anonymous"></script>
	</body>

</html>
```