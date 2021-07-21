package jsonutil

import (
	"fmt"
	"reflect"
)

const (
	typeSliceInterface     = "[]interface {}"
	typeMapStringInterface = "map[string]interface {}"
)

func HasKey(element interface{}, key string) bool {
	if reflect.TypeOf(element).String() != typeMapStringInterface {
		return false
	}
	_, ok := element.(map[string]interface{})[key]
	return ok
}

func GetKeyAsJsonObject(el interface{}, key string) (map[string]interface{}, error) {
	if !HasKey(el, key) {
		return nil, fmt.Errorf("%q does not exist in input", key)
	}
	obj := AsJsonObject(el)
	val := obj[key]
	ret := AsJsonObject(val)
	if ret == nil {
		return nil, fmt.Errorf("element %q is not a valid JSON object", key)
	}
	return ret, nil
}

func AsJsonArray(arr interface{}) []interface{} {
	if reflect.TypeOf(arr).String() == typeSliceInterface {
		return arr.([]interface{})
	}
	return nil
}

func AsJsonObject(obj interface{}) map[string]interface{} {
	if obj == nil {
		return nil
	}
	if reflect.TypeOf(obj).String() == typeMapStringInterface {
		return obj.(map[string]interface{})
	}
	return nil
}

func ExtractID(data interface{}) (string, error) {
	return ExtractAsString(data, []string{"id", "tsmMetricKey", "entityId"})
}

func ExtractAsString(data interface{}, keys []string) (string, error) {
	jsonObj := AsJsonObject(data)
	if jsonObj == nil {
		return "", fmt.Errorf("object is not a valid JSON object")
	}

	var el interface{}
	var key string
	for _, k := range keys {
		if val, ok := jsonObj[k]; ok {
			el = val
			key = k
			break
		}
	}
	if el == nil {
		return "", nil
	}
	elType := reflect.TypeOf(el).String()
	if elType != reflect.String.String() {
		return "", fmt.Errorf("expected %q to be string, got %s", key, elType)
	}
	return el.(string), nil
}

func OrderUnawareEquals(a, b interface{}) bool {
	if reflect.DeepEqual(a, b) {
		return true
	}

	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)

	kind := valA.Kind()
	if kind != valB.Kind() {
		return false
	}

	switch kind {
	case reflect.Map:
		return mapsMatch(valA, valB)
	case reflect.Slice:
		return slicesMatchInRandomOrder(valA, valB)
	case reflect.Struct:
		return structsMatch(valA, valB)
	default:
		return false
	}
}

func mapsMatch(valA, valB reflect.Value) bool {
	if valA.Len() != valB.Len() {
		return false
	}
	iter := valA.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value().Interface()
		b := valB.MapIndex(k).Interface()
		if !OrderUnawareEquals(v, b) {
			return false
		}
	}
	return true
}

func structsMatch(valA, valB reflect.Value) bool {
	if valA.NumField() != valB.NumField() {
		return false
	}
	for i := 0; i < valA.NumField(); i++ {
		if !OrderUnawareEquals(valA.Field(i).Interface(), valB.FieldByName(valA.Type().Field(i).Name).Interface()) {
			return false
		}
	}
	return true
}

// does not handle nested objects
func slicesMatchInRandomOrder(valA, valB reflect.Value) bool {
	visited := make(map[interface{}]int)
	for i := 0; i < valA.Len(); i++ {
		visited[valA.Index(i).Interface()]++
	}
	for i := 0; i < valB.Len(); i++ {
		visited[valB.Index(i).Interface()]--
	}
	for _, v := range visited {
		if v != 0 {
			return false
		}
	}
	return true
}
