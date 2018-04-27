package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJSON(t *testing.T) {

	// create a plant variable of Plant struct type
	plant := Plant{Name: "Ginger", Description: "A yellow root herb.", Kind: "Tuberous", Amount: "1"}

	plantJSON := plant.ToJSON()

	// t error, the expected string value to match with, plantJSON, error message
	assert.Equal(t, `{"name":"Ginger","description":"A yellow root herb.","kind":"Tuberous","amount":"1"}`, string(plantJSON), "JSON Marshalling Error")
}

func TestFromJSON(t *testing.T) {

	plantJSON := []byte(`{"name":"Ginger","description": "A yellow root herb.","KinD": "Tuberous","Amount": "1"}`)

	plant := FromJSON(plantJSON)

	assert.Equal(t, Plant{Name: "Ginger", Description: "A yellow root herb.", Kind: "Tuberous", Amount: "1"}, plant, "JSON UN-Marshalling Error")
}
