package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

type Users struct {
	Users []User `json:users`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	jsonFilePath := "users.json"
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Errorf("Error happened during the openning json file!")
	}

	jsonByte, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Errorf("Error happened during the reading of the json file!")
	}

	var mapStruct map[string]interface{}
	json.Unmarshal(jsonByte, &mapStruct)

	for k,v := range mapStruct {
		fmt.Printf("The element is: %v\n",k)
		fmt.Printf("The element value is: %v\n",v)
	}
	

}