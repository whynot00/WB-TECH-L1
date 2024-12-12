package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	fmt.Println(determineType(123))

}

func determineType(value interface{}) string {

	var typeStr string

	switch v := value.(type) {
	case int, string, bool:
		typeStr = fmt.Sprintf("%T", v)
	default:
		typeStr = "other type"

		if isChan(v) {
			typeStr = "channel"
		}
	}

	return typeStr
}

func isChan(value interface{}) bool {
	valType := strings.Split(reflect.TypeOf(value).String(), " ")

	return valType[0] == "chan"

}
