package handlers

// web app <-> browser, using http to communicate
import (
	"net/http"
)

// connect handler to url with "gorilla web toolkit"
func IndexPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", map[string]any{
		"Name": "Amri",
		"City": "BKK",
	})
}



