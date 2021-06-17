package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/valyala/fastjson"
)

type Car struct {
	Name    string `json:"Name"`
	Band    string `json:"Band"`
	Company string `json:"Company"`
	Email   string `valid:"email"`
}

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

//判断是否是有效的string
func isJsonString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil
}

//判断是否是有效的json字符串
func isJson(str string) bool {
	var js map[string]interface{}
	// var js json.RawMessage 不能使用这种类型，会将string也误认为是合法的json
	return json.Unmarshal([]byte(str), &js) == nil
}

func main() {
	var tests = []string{
		`"FirstString"`,
		`secondString`,
		`"keyString":"VALUE"`,
		`{"ttt":"afasdf"}`,
	}

	for _, t := range tests {
		fmt.Printf("isJsonString(%s) = %v\n", t, isJsonString(t))
		fmt.Printf("isJson(%s) = %v\n", t, isJson(t))
		fmt.Printf("isJson(%s) checked by the govalidator = %v\n", t, govalidator.IsJSON(t))
	}

	testCar := Car{
		Name:    "benz cla",
		Band:    "benz",
		Company: "Benz",
		Email:   "tttt@#23.com",
	}

	//使用govalidator检查其email
	//govalidator: https://github.com/asaskevich/govalidator
	fmt.Printf("The email is valid or not: %v\n", govalidator.IsEmail(testCar.Email))

	//Use fastjson to validate the string
	data := []byte(`{"foo": [1.23,{"bar":33,"baz":null}]}`)
	fmt.Printf("exists of the foo in data: %v\n", fastjson.Exists(data, "foo"))

	sampleJsonStr := `{"foo": [1.23,{"bar":33,"baz":null}]}`
	err := fastjson.Validate(sampleJsonStr)
	if err != nil {
		// see https://github.com/bahlo/go-styleguide for error guideline
		fmt.Errorf("The sample json is not a valid json %w", err)
	}

	// var byteValue []byte
	// jsonFilePath := `"/Users/i323691/work_dir/training/go-training/golang-playbook/third-play/users.json"`
	jsonFilePath := "users.json"
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Errorf("Cannot open the target json file!")
	}

	//使用ioutil来读取本地存储的json文件
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Errorf("The ioutil read json file has error!")
	}

	var users Users
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Errorf("Error happened during the unmarshalling!")
	}

	fmt.Printf("The User list's length: %v\n", len(users.Users))

	for _, item := range users.Users {
		fmt.Printf("Print users: %v\n", item.Name)
	}

	fmt.Println()

}

type Company struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

type Car2 struct {
	Band    string `json:"band"`
	Company Company
	Color   string `json:"color"`
	Like    bool   `json:"like"`
}

func dealWithMapString() {

	//格式化输出json - especially for map[string]interface{}
	cars := map[string]interface{}{
		"band": "tanke",
		"company": struct {
			name string
			year float64
		}{"tanke-company", 2021},
		"color": "black",
		"like":  true,
	}

	dataByte, error := json.Marshal(cars)

	// error := json.Unmarshal(cars, )

	fileName := "map-string-interface-cars.json"

}
