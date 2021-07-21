package test

import(
	"errors"
	"jsonutil"
	"json"
)

func useJsonUtil(existedStr interface{}, newValue string) (outputBool bool, err error ){

	strAsJSON := jsonutil.AsJsonObject(existedStr)
	if strAsJSON == nil {
		return false, errors.New("not a valid JSON object")
	}
	value, ok := strAsJSON["firstKeyItem"]
	if !ok {
		return false, nil
	}
	for _, item := range jsonutil.AsJsonArray(value) {
		// cast the value because each setting is map[string]interface
		itemsJSON := jsonutil.AsJsonObject(item)
		itemValue, ok := itemsJSON["valueOfItems"]
		if !ok {
			continue
		}
		var newVal interface{}
		err = json.Unmarshal([]byte(newValue), &newVal)
		if err != nil {
			return false, err
		}
		// compare the input setting with the returned setting from Dynatrace
		equal := jsonutil.OrderUnawareEquals(newVal, itemValue)
		if equal {
			return true, nil
		}
	}
	return false, nil
}
