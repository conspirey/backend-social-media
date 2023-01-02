package sessions

import (
	"main/functions"
	"main/functions/security"
	"net/http"
	"time"

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

func MiddleWare(name, EncrKey string, MaxAge int, Domain, Path string, HTTPOnly, secure bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		s := &Session{
			Name: name,
			EncrKey: EncrKey,
			Options: &Options{
				MaxAge: MaxAge,
				Path: Path,
				Domain: Domain,
				HttpOnly: HTTPOnly,
				Secure: secure,


			},
			Values: map[string]any{},
		}
		str, err := c.Cookie(s.Name)
		if err == nil {
			strD, err := security.Decrypt(str, EncrKey)
			if err == nil {
				s.Values = functions.StringToValue[map[string]any](strD)
			}
			// s.Values = security.Decrypt()
		}
		c.Set(defKey, *s)
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
func (s *Session) Set(key string, value any) any {
	// c.Get()
	s.Values[key] = value

	s.Written = true
	return ""
}
func (s *Session) Get(key string) any {
	return s.Values[key]
}
/*

*/
func GetDec(cookie, EncrKey string) (string, error) {
	return security.Decrypt(cookie, EncrKey)
}
func (s *Session) ValuesToString() string {
	return functions.MapToJSON(s.Values)
}
func (s *Session) Save(c *gin.Context) error {
	if s.Written {
		encoded, err := security.Encrypt(s.ValuesToString(), s.EncrKey)
		if err != nil {
			return err
		}
		c.Set(defKey, *s)

		c.SetCookie(s.Name, encoded, s.Options.MaxAge, s.Options.Path, s.Options.Domain, s.Options.Secure, s.Options.HttpOnly)
		
		s.Written = false
	}
	return nil
}
func NewCookie(name, value string, options Options) *http.Cookie {
	cookie := newCookieFromOptions(name, value, &options)
	if options.MaxAge > 0 {
		d := time.Duration(options.MaxAge) * time.Second
		cookie.Expires = time.Now().Add(d)
	} else if options.MaxAge < 0 {
		// Set it to the past to expire now.
		cookie.Expires = time.Unix(1, 0)
	}
	return cookie
}




func newCookieFromOptions(name, value string, options *Options) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
		SameSite: options.SameSite,
	}

}