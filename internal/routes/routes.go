package routes

import (
	"net/http"

	"github.com/jersonmartinez/OpenWebinars/internal/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/error", handlers.errorHandler)

	fs := http.FileServer(http.Dir("web/static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
