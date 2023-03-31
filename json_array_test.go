package golearnjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street string
	Country string
	PostalCode string
}

type People struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies []string
	Address []Address
}

func TestJsonArray(t *testing.T) {
	var people = People{
		FirstName: "Akhsan",
		MiddleName: "Bayu",
		LastName: "Riantama",
		Age: 24,
		Married: false,
		Hobbies: []string{"Fishing", "Gaming"},
	}

	bytes, _ := json.Marshal(people)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Akhsan","MiddleName":"Bayu","LastName":"Riantama","Age":24,"Married":false,"Hobbies":["Fishing","Gaming"]}`
	jsonBytes := []byte(jsonString)

	people := &People{}

	err := json.Unmarshal(jsonBytes, people)
	if err != nil {
		panic(err)
	}
	fmt.Println(people)
}

func TestJsonArrayComplex(t *testing.T) {
	people := People{
		FirstName: "Akhsan",
		MiddleName: "Bayu",
		LastName: "Riantama",
		Age: 24,
		Married: false,
		Hobbies: []string{"Fishing", "Coding"},
		Address: []Address{
			{
				Street: "Kopeng KM5",
				Country: "Indonesia",
				PostalCode: "56192",
			},
		},
	}

	bytes, _ := json.Marshal(people)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Akhsan","MiddleName":"Bayu","LastName":"Riantama","Age":24,"Married":false,"Hobbies":["Fishing","Coding"],"Address":[{"Street":"Kopeng KM5","Country":"Indonesia","PostalCode":"56192"}]}`
	jsonBytes := []byte(jsonString)

	people := &People{}

	err := json.Unmarshal(jsonBytes, people)
	if err != nil {
		panic(err)
	}
	fmt.Println(people)
}
