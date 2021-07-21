package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:users`
}

type User struct {
	Name string `json:"name"`
	Type   string `json:"type"`
    Age    int    `json:"Age"`
    Social Social `json:"social"`
}

type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

type TestBool struct {
	KSettings KSettings `json:"k8s"`
}
type KSettings struct {
	Untouch string `json:"untouch"`
}



func main(){
	jsonFile, err := os.Open("test.json")

	if err != nil {
		fmt.Println("cannot read the json file")
	}

	fmt.Println("Successfully read the json file")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		// fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}


	var result map[string]interface{}
	
    json.Unmarshal([]byte(byteValue), &result)

    fmt.Println(result["users"])

	responseBody := new(bytes.Buffer)
	fmt.Print(responseBody)

	testStr := `{}`

	boolObj := new(TestBool)

	err = json.Unmarshal([]byte(testStr), &boolObj)
	if err != nil {
		fmt.Println("have error for unmarshalling boolean value")
	}

	b1, _ := strconv.ParseBool(boolObj.KSettings.Untouch)
	fmt.Println(b1)

	fmt.Println(boolObj.KSettings.Untouch)

	if boolObj.KSettings.Untouch == "true" {
		fmt.Printf("it is boolean, %v, type is %T \n", boolObj.KSettings.Untouch, boolObj.KSettings.Untouch)
	}
	if boolObj.KSettings.Untouch == ""{
		fmt.Println("user didnt set anything")	
	}
	if b1 {
		fmt.Printf("it is boolean, %v, type is %T", boolObj.KSettings.Untouch, b1)
	}
	




}