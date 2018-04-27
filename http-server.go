package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	InteruptHandler()

	// tells the http package to handle all requests to the web root ("/") with AllRequestHandler.
	http.HandleFunc("/", RootRequestHandler)

	http.HandleFunc("/echo", QueryEchoHandler)

	fmt.Printf("HTTP Server Awakens\n")

	// Starts HTTP server, specifying that it should listen on the port number on any interface
	// This function will block until the program is terminated
	//
	// log.Fatal will log unexpected errors
	log.Fatal(http.ListenAndServe(Port(), nil))

}

// Port looks for $PORT env on OS
// and then sets it as the port number that the HTTP server will serve on
func Port() string {

	// on the terminal: export $PORT=number_you_want
	// https://stackoverflow.com/questions/1158091/defining-a-variable-with-or-without-export
	portNumber := os.Getenv("PORT")

	if len(portNumber) == 0 {
		fmt.Printf("$PORT env is empty/not specified!\n")

		portNumber = "8080"
	}

	fmt.Printf("Serving on " + "localhost:" + portNumber + "\n")

	// Specifying localhost in the port number will avoid the OS from asking if you trust this connection
	return "localhost:" + portNumber
}

// InteruptHandler info can be found in the link below.
// https://stackoverflow.com/questions/11268943/is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in-a-defe
func InteruptHandler() {

	//short variable declaration
	// var i = 1
	// v := 1

	// create channel c
	c := make(chan os.Signal, 2)

	// notify signal to c
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// goroutine
	// https://blog.golang.org/pipelines
	go func() {

		// send c to goroutine
		<-c

		// \r remvoves ^C being shown on the terminal
		fmt.Println("\rI'm shutting down...")

		fmt.Println("Bye bye!")

		os.Exit(0)
	}()

}

// RootRequestHandler handles requests for the "/" page
//
// This function is of the type: http.HandlerFunc
//
// The w value assembles the HTTP server's response; by writing to it, sends data to the HTTP client.
// r is a data structure that represents the client HTTP request
func RootRequestHandler(w http.ResponseWriter, r *http.Request) {

	// respond with HTTP status code OK == HTTP 200
	w.WriteHeader(http.StatusOK)

	// shows on terminal
	fmt.Println("Request Succeeded")

	// Shows on index page
	fmt.Fprintf(w, "Written in GO")
}

// QueryEchoHandler handles API echos and places echos on the webpage
func QueryEchoHandler(w http.ResponseWriter, r *http.Request) {

	// save the first Query parameter from URL
	// name the parameter as q on URL
	query := r.URL.Query()["q"][0]

	// shows on terminal
	if len(query) != 0 {

		fmt.Println("Query asked: " + query)

	} else {

		fmt.Println("Empty Query Entery")
	}

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, query)
}
