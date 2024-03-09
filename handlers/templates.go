package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, filename string, data map[string]any) {

	//find all templates in the templates dir
	templates, err := filepath.Glob("templates/*.html")

	// err !empyty -> something wrong about file path
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// read all temjplates from the disk
	t, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//send the template to client
	err = t.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
