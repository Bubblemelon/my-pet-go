package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/bubblemelon/my-pet-go/plant-keeper/database"
)

// Plant properties
//
// store keys in JSON as lower case
type Plant struct {
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"` //will be omited if empty - optional field
	Kind        string    `json:"kind"`
	Amount      string    `json:"amount"`
	Timestamp   time.Time `json:timestamp`
	Updates     int       `json:updates-number`
}

// ToJSON takes as its receiver "p", a variable to Plant.
// Takes no parameters and returnes a slice of type byte (a byte slice)
//
// Converts Plant variable properties into JSON format
// func (p Plant) ToJSON() []byte {
//
// 	// returns the JSON encoding of p as a byte slice
// 	pJSON, err := json.Marshal(p)
//
// 	// if error occurs in json.Marshall
// 	if err != nil {
// 		log.Fatal(err)
// 		// unwinds the stack goroutine when reaches an unrecoverable state
// 		panic(err)
// 	}
//
// 	return pJSON
// }

// ToJSON takes a Plant variable and returns it in JSON format
func ToJSON(p Plant) []byte {

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

// PlantsHandler handles requests for the All Plant data stored
// Returns all plants in JSON format
func PlantsHandler(w http.ResponseWriter, r *http.Request) {

	// test data named testPlant
	// localTestPlant := &Plant{Name: "Brinjal", Description: "Purple longish fruit.", Kind: "Fruit", Amount: "5"}

	// jsonTestPlants := localTestPlant.ToJSON()

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// w.Write(jsonTestPlants)
}

// PlantSpecificHandler handles requests for A Specific Plant via KEY
// Listens for:
// GET
// POST
func PlantSpecificHandler(w http.ResponseWriter, r *http.Request) {

	switch method := r.Method; method {

	// create a Plant Record
	case http.MethodPost:
		JSONInput, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		// convert input from JSON format
		input := FromJSON(JSONInput)

		// Key does not exist
		if KeyExistance(input.Name) == false {

			// create plant record
			record := CreatePlantRecord(input)

			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.Header().Add("Location", "/api/plant/"+WhiteSpaceRemover(input.Name))
			w.Write(record)

			fmt.Println("Post Request Completed for %v", record)

			w.WriteHeader(http.StatusCreated)
		} else {
			// Key does Exist!

			w.Header().Add("Content-Type", "text/plain")
			fmt.Fprintf(w, "Plant with the same name already exists!")
			w.WriteHeader(http.StatusConflict)

		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}

}

// CreatePlantRecord saves a record of plants properties into Redis
func CreatePlantRecord(p Plant) []byte {

	c := database.RedisConnection()

	//defers execution until function returns
	defer c.Close()

	// non-user inputs
	p.Timestamp = time.Now()
	// initial value
	p.Updates = 0

	// convert to JSON format
	pJSON := ToJSON(p)

	// save JSON blob into Redis
	c.Do("SET", KeyClean(p.Name), pJSON)

	fmt.Printf("Key " + KeyClean(p.Name) + " creation: ")
	fmt.Printf("%t \n", KeyExistance(p.Name))

	return pJSON

}

// WhiteSpaceRemover removes white space or spaces i.e. for p.Name to be used as a KEY
func WhiteSpaceRemover(s string) string {

	// remove whitespace in s to save as key
	// may require a better function
	// https://stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string
	sNoSpaces := strings.Replace(s, " ", "", -1)

	return sNoSpaces
}

//UpperCaseRemover removes all uppercases from a string and returns the string in lowercases
func UpperCaseRemover(s string) string {

	return strings.ToLower(s)
}

// KeyClean removes Uppercase and whitespaces
func KeyClean(s string) string {

	return WhiteSpaceRemover(UpperCaseRemover(s))
}

//KeyExistance checks to see if KEY exists
func KeyExistance(s string) bool {

	c := database.RedisConnection()

	//defers execution until function returns
	defer c.Close()

	// remove spaces from p.Name
	sClean := KeyClean(s)

	// check if key already exists
	exists, err := redis.Bool(c.Do("EXISTS", sClean))

	if err != nil {
		panic(err)
	}

	return exists

}

// init() will be run automatically upon running
//
// creates sample data to store in redis
// func name is in lowercase lowercase since there’s no need to call it from another package
func init() {

	fmt.Printf("Sample data is now being stored into Redis...\n")

	plant := Plant{Name: "Egg Plant", Description: "Purple longish fruit.", Kind: "Fruit", Amount: "5"}

	CreatePlantRecord(plant)

}
