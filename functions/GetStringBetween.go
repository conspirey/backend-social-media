package functions

import (
	"reflect"
	"strings"
)

func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	return str[s : s+e]
}
func GetStructField(strucT any, FieldName string) (reflect.StructField, bool) {
	//fmt.Println(reflect.TypeOf(strucT))
	return reflect.TypeOf(strucT).Elem().FieldByName(FieldName)
}

// field, ok := reflect.TypeOf(user).Elem().FieldByName("Name")
func GetStructTag(f reflect.StructField) string {
	return string(f.Tag)
}
func GetStrJSONtag(trucT any, FieldName string) string {
	field, _ := reflect.TypeOf(trucT).Elem().FieldByName(FieldName)
	return GetStringInBetween(GetStructTag(field), "json:\"", "\"")
}
