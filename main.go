package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//fmt.Fprintf(w, r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send me to <a href=\"mailto:oscar.taracena@gmail.com\">oscar.taracena@gmail.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1> We could not find the page you are looking for :(</h1><p>please email us if you keep getting sent to an invalid page</p>")
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerFunc)
	r.HandleFunc("/contact", handlerFunc)
	http.ListenAndServe(":3000", r)
}
