package api

import (
	"fmt"
	"net/http"
)

// QueryEchoHandler handles API echos and places echos on the webpage
//
// e.g. http://localhost:5050/api/echo?q=hello+there
func QueryEchoHandler(w http.ResponseWriter, r *http.Request) {

	// save the first Query parameter from URL
	// name the parameter as q on URL
	// r.URL.Query()["q"][0]
	query := r.URL.Query().Get("q")

	// if len(query) != 0
	if query == "" {

		query = "Your Query was not specified\n"

		// prints on terminal
		fmt.Println("Query entry is empty")
		// http://localhost:5050/api/echo?q
		// http://localhost:5050/api/echo?q=
		// Will also be empty if query was only spaces

		// THIS IS NOT EMPTY:
		// http://localhost:5050/api/echo?q=""
		// http://localhost:5050/api/echo?q= f

	} else {

		fmt.Println("Query asked: " + query)

	}

	w.Header().Add("Content-Type", "text/plain")

	// prints on browser (i.e. prints on where request was sent from)
	fmt.Fprintf(w, query)
}
