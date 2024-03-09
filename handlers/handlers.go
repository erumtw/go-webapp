package handlers
// web app <-> browser, using http to communicate
import "net/http"


// connect handler to url with "gorilla web toolkit"
func IndexPage(w http.RespondseWriter, r *http.Request){
	w.Write([]byte("Hello golang\n"))
}