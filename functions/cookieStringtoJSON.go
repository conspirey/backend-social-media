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