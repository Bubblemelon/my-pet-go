package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Plant properties
//
// store keys in JSON as lower case
type Plant struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Kind        string `json:"kind"`
	Amount      string `json:"amount"`
}

// ToJSON takes as its receiver "p", a variable to Plant.
// Takes no parameters and returnes a slice of type byte (a byte slice)
//
// Converts Plant variable properties into JSON format
func (p Plant) ToJSON() []byte {

	// returns the JSON encoding of p as a byte slice
	pJSON, err := json.Marshal(p)

	// if error occurs in json.Marshall
	if err != nil {
		log.Fatal(err)
		// unwinds the stack goroutine when reaches an unrecoverable state
		panic(err)
	}

	return pJSON
}

// FromJSON takes the parameter of a byte slice
// and returns a variable of a Plant type
//
// Converts JSON format to Plant type
func FromJSON(jsonData []byte) Plant {

	// empty Plant
	plant := Plant{}

	// converts jsonData to store into the address of empty plant variable
	//
	// works with {"key": 'value"} space between
	// does not need to be {"key":'value"}
	// and can have a mixture of both of the above
	//
	// keys can also be mixed upper and lowercase
	err := json.Unmarshal(jsonData, &plant)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return plant
}

// Should request from redis
// instead of using test data named testPlant

// PlantHandler handles requests for the plant data api
func PlantHandler(w http.ResponseWriter, r *http.Request) {

	// test data named testPlant
	testPlant := &Plant{Name: "Brinjal", Description: "Purple longish fruit.", Kind: "Fruit", Amount: "5"}

	jsonTestPlants := testPlant.ToJSON()

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonTestPlants)
}
