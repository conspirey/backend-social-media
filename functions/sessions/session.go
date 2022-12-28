package sessions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Session struct {
	EncrKey string
	
	GetReq func(c *http.Request) any
	SetGin  func(c *gin.Context) any
	SetReq func(c *http.Request) any
	Name string
}
func NewSession(Name, EncrKey string) *Session {
	return &Session{
		EncrKey: EncrKey,
		Name: Name,
	}
}
func (s *Session) GetGin(c *gin.Context) any {
	// c.Get()
	
	return ""
}