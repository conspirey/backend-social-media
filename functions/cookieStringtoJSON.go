package functions

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func ConvertCookieValueToJSON(cookie http.Cookie) {
	cookieJSON, err := json.Marshal(cookie)
	if err != nil {
		fmt.Println(err)
		return
	}

	cookieJSONString := string(cookieJSON)
	fmt.Println(cookieJSONString)
}
func Testing() {
	cookie := http.Cookie{
		Name: "mrredo",
		Value: "value name",
	}
	ConvertCookieValueToJSON(cookie);
}

func StringToValue[T any](val string) T {
	sec := new(T)
    if err := json.Unmarshal([]byte(val), &sec); err != nil {
        panic(err)
    }
	return *sec

}
//Type needs to be map[T]T for this to work
func MapToJSON[T any](val T) string {
    jsonStr, err := json.Marshal(val)
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    }
	return string(jsonStr)

}