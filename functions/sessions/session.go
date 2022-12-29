package sessions

import (
	"fmt"
	"main/functions/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
)
var (
	defKey = "conspirey/session/func"
)
// Options stores configuration for a session or session store.
//
// Fields are a subset of http.Cookie fields.
type Options struct {
	Path   string
	Domain string
	// MaxAge=0 means no Max-Age attribute specified and the cookie will be
	// deleted after the browser session ends.
	// MaxAge<0 means delete cookie immediately.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge   int
	Secure   bool
	HttpOnly bool
	// Defaults to http.SameSiteDefaultMode
	SameSite http.SameSite
}

func MiddleWare(name, EncrKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := &Session{
			Name: name,
			EncrKey: EncrKey,
		}
		c.Set(defKey, s)
		defer context.Clear(c.Request)
		c.Next()
	}
}
type Store struct {

}
func Default(c *gin.Context) Session {
	return c.MustGet(defKey).(Session)
}
type Session struct {
	C *gin.Context
	EncrKey string
	Values map[string]any
	Options *Options
	Name string
	Written bool
}
func (s *Session) Set(c *gin.Context) any {
	// c.Get()


	s.Written = true
	return ""
}
func (s *Session) Get(c *gin.Context) any {
	// c.Get()
	
	return ""
}
/*

*/
func (s *Session) GetDec(c *gin.Context) any {
	// c.Get()
	
	return ""
}
func (s *Session) Save() error {
	if s.Written {
		
		encoded, err :=security.Encrypt(fmt.Sprintf("%v", s.Values), s.EncrKey)
		if err != nil {
			return err
		}
		http.SetCookie(s.C.Writer, NewCookie(s.Name, encoded, ))
		s.Written = false
	}
}
func NewCookie(name, value string, options Options) *http.Cookie {
	return nil
}




