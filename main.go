package main // this is how go divides

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/erumtw/go-webapp/handlers"
	"net/http"
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
	r.HandlerFunc("/", handlers.IndexPage)

	// start ther server 
	http.ListenAndServe(":3000", r)

}

