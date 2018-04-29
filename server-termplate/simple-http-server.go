package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bubblemelon/my-pet-go/server-termplate/api"
)

func main() {

	http.HandleFunc("/", root)

	http.HandleFunc("/api/sample", api.SampleHandler)

	// start HTTP server
	http.ListenAndServe(port(), nil)

	fmt.Println("Serving")
}

func port() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return "localhost:" + port
}

// http.ResponseWriter == responds content to client
// pointer to HTTP request received
func root(w http.ResponseWriter, r *http.Request) {

	// respond with HTTP status code OK == HTTP 200
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hello I'm the Root Page! \n")
}
