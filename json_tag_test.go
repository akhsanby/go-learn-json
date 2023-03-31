package golearnjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id: "001",
		Name: "Bakso Beku",
		ImageUrl: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"001","name":"Bakso Beku","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)
	
	product := &Product{}
	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
}