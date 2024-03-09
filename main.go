package main // this is how go divides

import (
	// "fmt"
	"net/http"

	"github.com/erumtw/go-webapp/handlers"
	"github.com/gorilla/mux"
)

// compile languages , golang here
// interpret langauages

// web app is a set of pages(urls)
// url
// / -> root

// every url is connected to a function(handler)
func main() {
	// router is making choice based on the url
	// which jandler to call
	r := mux.NewRouter()

	// connect the url to the handler
	r.HandleFunc("/", handlers.IndexPage)
	r.HandleFunc("/user/{username}", handlers.UserPage)

	// start ther server
	http.ListenAndServe(":3000", r)
}
