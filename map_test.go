package golearnjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"001", "name":"Bakso Beku", "image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id": "0002",
		"name": "Sosis Beku",
		"price": 10000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}