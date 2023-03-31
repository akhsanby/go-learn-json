package golearnjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Human struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestDecode(t *testing.T) {
	var jsonString = `{"FirstName":"Akhsan","MiddleName":"Bayu","LastName":"Riantama","Age":24,"Married":false}`
	var jsonBytes = []byte(jsonString)

	var human = &Human{}

	err := json.Unmarshal(jsonBytes, human)
	if err != nil {
		panic(err)
	}

	fmt.Println(human)
	fmt.Println(human.FirstName)
	fmt.Println(human.MiddleName)
	fmt.Println(human.LastName)
}