package handlers

import (
	"bytes"
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

	// Head functions
	headContent, err := os.ReadFile("web/templates/head.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	headTemplate, err := template.New("head").Parse(string(headContent))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	var headBuffer bytes.Buffer
	err = headTemplate.Execute(&headBuffer, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data.HeadContent = template.HTML(headBuffer.String())

	// Navbar functions
	navbarContent, err := os.ReadFile("web/templates/navbar.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	navbarTemplate, err := template.New("navbar").Parse(string(navbarContent))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	var navbarBuffer bytes.Buffer
	err = navbarTemplate.Execute(&navbarBuffer, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data.NavbarContent = template.HTML(navbarBuffer.String())

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
