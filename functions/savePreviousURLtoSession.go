package functions

import (
	"fmt"

	"github.com/gin-contrib/sessions"
)

func SaveURL(session sessions.Session, path string) {
	session.Set("prevurl", path);
	err := session.Save()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

}