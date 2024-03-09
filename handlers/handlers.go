package handlers

// web app <-> browser, using http to communicate
import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// connect handler to url with "gorilla web toolkit"
func IndexPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

var users = map[string]map[string]any{
	"erum": {
		"Name": "Erum",
		"City": "Satun",
	},

	"fair": {
		"Name": "Fair",
		"City": "Pitlok",
	},

	"mids": {
		"Name": "Mids",
		"City": "Udon",
	},
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, isFound := users[vars["username"]]

	log.Printf("\nUser: %v, \nisFound: %v", user, isFound)

	if !isFound {
		http.Error(w, "UserNotFound", http.StatusNotFound)
		// http.NotFound(w, r)
		return
	}
	renderTemplate(w, "user.html", user)
}
