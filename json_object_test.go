package golearnjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age int
	Married bool
}

func TestJsonObject(t *testing.T) {
	var customer = Customer{
		FirstName:  "Akhsan",
		MiddleName: "Bayu",
		LastName:   "Riantama",
		Age: 24,
		Married: false,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}