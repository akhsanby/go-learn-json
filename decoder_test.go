package golearnjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestStreamingEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")

	encoder := json.NewEncoder(writer)
	
	customer := &Customer{
		FirstName: "Akhsan",
		MiddleName: "Bayu",
		LastName: "Riantama",
	}

	encoder.Encode(customer)
}