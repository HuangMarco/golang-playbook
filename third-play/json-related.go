package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"reflect"

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

	fmt.Println("Start to deal with map string interface...")
	
	dealWithMapString()

	encodeAndDecode()

	testCheckBooleanPointer()
}

type Company struct {
	Name string `json:"name"`
	Year float64   `json:"year"`
}

type Car2 struct {
	Band    string `json:"band"`
	Company Company
	Color   string `json:"color"`
	Like    bool   `json:"like"`
}

func dealWithMapString() {

	cars := map[string]interface{}{
		"band": "tanke",
		"company": struct {
			Name string
			Year float64
		}{"tanke-company", 2021},
		"color": "black",
		"like":  true,
	}

	//格式化输出json - especially for map[string]interface{}
	dataByte, error := json.Marshal(cars)
	if error != nil {
		// erros.new("error happened")
		fmt.Println("error happended during the marshalling")
	}

	var carsMapStringInterface map[string]interface{}

	error = json.Unmarshal(dataByte, &carsMapStringInterface)

	for k,v := range carsMapStringInterface {
		switch c := v.(type) {
		case string:
		  fmt.Printf("Item %q is a string, containing %q\n", k, c)
		case float64:
		  fmt.Printf("It looks like item %q is a number, specifically %f\n", k, c)
		default:
		  fmt.Printf("Don't know what type item %q is, but I think it might be %T\n", k, c)
		}
	}

	//Another way to print the json of the map[string]interface{}
	output, error2 := json.Marshal(carsMapStringInterface)
	if error2 != nil {
		fmt.Println("Error happened during the marshalling")
	}

	fmt.Println(string(output))
	fmt.Println("printing json format completed.")

	//Directly loop the cars as the map[string]interface{}
	for k,v := range cars {
		fmt.Println(k,v)
	}

	var car2 Car2

	unmarshalError := json.Unmarshal(dataByte, &car2)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}

	fmt.Printf("The company name: %s\n", car2.Company.Name)
	fmt.Printf("The length of the map[string]interface: %d, so that it is not empty", len(carsMapStringInterface))
	fmt.Println("Clear the Car2 map")

	car2 = Car2{}

	fmt.Println("Clear the map string interface")

	anotherDecode()

	compareTwoJson()

	createInitializeStruct()

	checkBoolean()

	testCheckBooleanPointer()
}

type Junk struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Area string `json:"area"`
	// Password bool `json:"password,omitempty"` //临时忽略password字段
}

func encodeAndDecode(){
	a := Junk{}
	data := `{"id":1,"name":"gg"}}`

	d := json.NewDecoder(strings.NewReader(data))
	d.DisallowUnknownFields()

	if err := d.Decode(&a); err != nil {
		fmt.Println(err)
	}
	if d.More() {
		fmt.Println("extra junk")
	}
	fmt.Println(a)
}

type testobjectSettingsItem struct {
	TestobjectId string `json:"testobjectId"`
	Value testobjectSettingsValue `json:"value"`
}

type testobjectSettingsResponse struct {
	Items []testobjectSettingsItem `json:"items"`
	TotalCount int `json:"totalCount"`
	PageSize int `json:"pageSize"`
}

type testobjectSettingsValue struct {
	Enabl  bool   `json:"enabl"`
	Modex     string `json:"modex"`
	Property string `json:"property"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

func createInitializeStruct (){
	// first way to create&initialize the struct
	var wayOne testobjectSettingsValue
	wayOne.Enabl = true
	wayOne.Modex = "test"
	wayOne.Property = "test-space"
	wayOne.Operator = "operator1"
	wayOne.Value = "test"

	//struct to string via json
	// fmt.Println(wayOne.(string))
	wayOneJson, error := json.Marshal(wayOne)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(string(wayOneJson))

	// second way to create&initialize the struct
	var secondOne *testobjectSettingsValue
	secondOne = new(testobjectSettingsValue)
	secondOne.Enabl = true
	secondOne.Modex = "test"
	secondOne.Property = "test-space"
	secondOne.Operator = "operator1"
	secondOne.Value = "test"
	secondOneOutput, error := json.Marshal(secondOne)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(string(secondOneOutput))

	//Fourth way to create&initialize the struct

	thirdOne := &testobjectSettingsValue{
		Enabl : true,
		Property: "test-space",
		Operator: "operator1",
		Value: "test",
		Modex: "test",
	}
	fourthOne := &testobjectSettingsValue{
		Enabl : true,
		Property: "test-space",
		Operator: "operator1",
		Value: "test",
		Modex: "test",
	}
	fourthOneOutput, error := json.Marshal(fourthOne)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(string(fourthOneOutput))
	fmt.Println(secondOne == fourthOne)
	fmt.Println(thirdOne == fourthOne)




}

func anotherDecode() error {
	var testobjSettingsResp testobjectSettingsResponse
	var testobjSettingsResp2 testobjectSettingsResponse

	// var values    = `{"enabled": true,"mode": "MONITORING_OFF","property": "KUBERNETES_NAMESPACE","operator": "EQUALS","value": "kube-system"}`

	var wholeItems = `{"items":[{"objectId":"vu9U3hXa3q0AAAABACFidWlsdGluOmNvbnRhaW5lci5tb25pdG9yaW5nLXJ1bGUABnRlbmFudAAGdGVuYW50ACRlYjg3MTIzZC1lM2NhLTM2ZTMtYjY1NS04MWQ0ZGY4NmNkYTO-71TeFdrerQ","value":{"enabl":true,"modex":"MONITORING_OFF","property":"KUBERNETES_NAMESPACE","operator":"EQUALS","value":"kube-system"}},{"objectId":"vu9U3hXa3q0AAAABACFidWlsdGluOmNvbnRhaW5lci5tb25pdG9yaW5nLXJ1bGUABnRlbmFudAAGdGVuYW50ACRlYjg3MTIzZC1lM2NhLTM2ZTMtYj1Y1NS04MWQ0ZGY4NmNkYTO-71TeFdrerQ","value":{"enabl":true,"modex":"MONITORING_OFF","property":"KUBERNETES_NAMESPACE","operator":"EQUALS","value":"kube-system"}}],"totalCount":4,"pageSize":100}`

	var wholeItemsByte = []byte(`{"items":[{"objectId":"vu9U3hXa3q0AAAABACFidWlsdGluOmNvbnRhaW5lci5tb25pdG9yaW5nLXJ1bGUABnRlbmFudAAGdGVuYW50ACRlYjg3MTIzZC1lM2NhLTM2ZTMtYjY1NS04MWQ0ZGY4NmNkYTO-71TeFdrerQ","value":{"enabl":true,"modex":"MONITORING_OFF","property":"KUBERNETES_NAMESPACE","operator":"EQUALS","value":"kube-system"}},{"objectId":"vu9U3hXa3q0AAAABACFidWlsdGluOmNvbnRhaW5lci5tb25pdG9yaW5nLXJ1bGUABnRlbmFudAAGdGVuYW50ACRlYjg3MTIzZC1lM2NhLTM2ZTMtYj1Y1NS04MWQ0ZGY4NmNkYTO-71TeFdrerQ","value":{"enabl":true,"modex":"MONITORING_OFF","property":"KUBERNETES_NAMESPACE","operator":"EQUALS","value":"kube-system"}}],"totalCount":4,"pageSize":100}`)
	// decodeErr := json.NewDecoder(strings.NewReader(values)).Decode(&objSettingsResp)
	decodeErr := json.NewDecoder(strings.NewReader(wholeItems)).Decode(&testobjSettingsResp)
	if decodeErr != nil {
		return decodeErr
	}

	anotherDecodeErr := json.NewDecoder(strings.NewReader(string(wholeItemsByte))).Decode(&testobjSettingsResp2)
	if anotherDecodeErr != nil {
		return anotherDecodeErr
	}

	for _, v := range testobjSettingsResp.Items {
		objValue := v.Value
		if objValue.Enabl && objValue.Modex == "MONITORING_OFF" && objValue.Operator == "EQUALS" && objValue.Value == "kube-system" && objValue.Property == "KUBERNETES_NAMESPACE" {
			fmt.Println("The settings has the target setting")
			return nil
		}
	}

	return nil


}

func compareTwoJson(){

	s1 := `{"dog": 5, "cat": 3}`
	s2 := `{"cat":3, "dog": 5}`

	s3 := `{"enabled": true,"mode": "INVALID-MODE","property": "KUBERNETES_NAMESPACE","operator": "EQUALS","value": "kube-system"}`
	s4 := `{"mode": "INVALID-MODE",  "enabled": true,"operator": "EQUALS","property": "KUBERNETES_NAMESPACE","value": "kube-system"}`

	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		fmt.Printf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		fmt.Printf("Error mashalling string 2 :: %s", err.Error())
	}

	isEqualOrNot := reflect.DeepEqual(o1, o2)

	fmt.Printf("The two string is equal or not: %t\n", isEqualOrNot)

	err = json.Unmarshal([]byte(s3), &o1)
	if err != nil {
		fmt.Printf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s4), &o2)
	if err != nil {
		fmt.Printf("Error mashalling string 2 :: %s", err.Error())
	}


	isEqualOrNot = reflect.DeepEqual(o1, o2)

	fmt.Printf("The another two string is equal or not: %t\n", isEqualOrNot)



}

type boolTest struct {
	BoolAttribute bool `json:"flagBool"`
}
func checkBoolean(){
	// testStr := `{"flagBool":"zxfafasdf"}`
	testStr := `{"flagBool":true}`
	
	boolObj := new(boolTest)

	err := json.Unmarshal([]byte(testStr), &boolObj)
	if err != nil {
		fmt.Println("have error for unmarshalling boolean value")
	}
	if boolObj.BoolAttribute {
		fmt.Println("it is a boolean value")
	}
	

}

type booleanPointer struct {
	BooleanAttribute *bool `json:"flagBool"`
}

func testCheckBooleanPointer(){
	// boolObj := new(booleanPointer)
	testBool := true

	checkBooleanPointer(&testBool)
}

func checkBooleanPointer(testBol *bool){
	// testStr := `{"flagBool":"zxfafasdf"}`
	testStr := `{"flagBool":false}`
	
	boolObj := new(booleanPointer)

	err := json.Unmarshal([]byte(testStr), &boolObj)
	if err != nil {
		fmt.Println("have error for unmarshalling boolean value")
	}
	if boolObj == nil {
		fmt.Println("user doesn't set the boolean value")
	}

	if (*boolObj.BooleanAttribute){
		fmt.Printf("user set the boolean value %t", *boolObj.BooleanAttribute)
	}

	if !(*boolObj.BooleanAttribute){
		fmt.Printf("user set the boolean value %t", *boolObj.BooleanAttribute)
	}
	// if testBol == nil {
	// 	fmt.Println("user doesn't set the boolean value")
	// }

	// if (*testBol){
	// 	fmt.Printf("user set the boolean value %t",*testBol)
	// }

	// if !(*testBol){
	// 	fmt.Printf("user set the boolean value %t",*testBol)
	// }
	

}
