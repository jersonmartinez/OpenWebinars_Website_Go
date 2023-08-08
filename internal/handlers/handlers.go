package handlers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/jersonmartinez/OpenWebinars_Website_Go/internal/models"
)

func renderTemplate(w http.ResponseWriter, tmplFile string, data interface{}) {
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title:   "OpenWebinars",
		Author:  "Jerson Martínez",
		Welcome: "Este es un curso que impactará en tus proyectos webs actuales y futuros.",
	}

	page := r.URL.Path[1:]

	if page == "" {
		page = "index.html"
	}

	tmplFile := "web/templates/" + page

	if _, err := os.Stat(tmplFile); err != nil {
		tmplFile = "web/templates/error.html"

		data.ErrorCode = http.StatusNotFound
		data.ErrorMessage = "Página no encontrada"
	}

	renderTemplate(w, tmplFile, data)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title:        "¡Página no encontrada!",
		ErrorCode:    http.StatusInternalServerError,
		ErrorMessage: "Error interno de servidor",
	}

	renderTemplate(w, "web/templates/error.html", data)
}
