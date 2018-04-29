package api

import (
	"fmt"
	"net/http"
	"strconv"
)

var callcount = 0

// SampleHandler handles specific request types
//
// curl -i -X $REQUEST http://localhost:$PORT/api/sample
func SampleHandler(w http.ResponseWriter, r *http.Request) {

	switch method := r.Method; method {

	case http.MethodGet:
		fmt.Println("GET requested")

		// return this if some logic is not desired
		// w.WriteHeader(http.StatusInternalServerError)

		w.Header().Add("Content-Type", "application/json; charset=utf-8")

		w.WriteHeader(http.StatusOK)

		//  Should be printing out JSON
		w.Write([]byte("GET requested"))

	case http.MethodPost:

		fmt.Println("POST requested")

		callcount++

		w.Header().Add("Location", "/api/sample/call"+strconv.Itoa(callcount))
		w.WriteHeader(http.StatusCreated)

		// if failed do:
		// w.WriteHeader(http.StatusConflict)

		w.Write([]byte("POST requested"))

	case http.MethodPut:
		fmt.Println("PUT requested")

		// When update is successful then:
		// w.WriteHeader(http.StatusOK)

		w.Write([]byte("PUT requested"))

	//DELETE
	case http.MethodDelete:
		fmt.Println("DELETE requested")

		// w.WriteHeader(http.StatusOK)

		w.Write([]byte("DELETE requested"))

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}

}
