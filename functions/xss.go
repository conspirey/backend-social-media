package functions

import (
	"regexp"
	"strings"
	"github.com/microcosm-cc/bluemonday"
)

var URLregex *regexp.Regexp = regexp.MustCompile(`(?m)(href| href)=("|'|\x60)(http|https):/(\/www\.|\/|\/.*\.)(discord\.com|github\.com|scratch-for-discord\.com|mongodb\.com|discord\.js\.org|discordjs\.guide|youtube\.com)(\/|).*("|'|\x60)`)
var EmptyHREFregex = regexp.MustCompile(`(?m)((href| href)=("|'|\x60).*?("|'|\x60))`)
func SanitizeHTMLString(htmlString string) string {
	p := bluemonday.UGCPolicy()

	p.AllowStandardURLs()
	//allowed links Only discord, GitHub, d.js, mongodb, or s4d relate
	p.AllowComments()
	p.AllowLists()
	p.AllowTables()
	p.AllowStyles()
	p.AllowStandardAttributes()
	p.AllowStyling()
	p.AllowElements("a")
	p.AllowAttrs("style", "class", "colspan", "datetime", "headers", "hreflang", "media", "start", "tabindex", "title", "translate", "value", "type").OnElements("details", "summary", "title", "ul", "ol", "li", "style", "title", "body", "article", "h1", "h2", "h3", "h4", "h5", "h6", "blockquote", "dd", "div", "dl", "dt", "figcaption", "hr", "p", "pre", "b", "br", "code", "i", "mark", "q", "s", "span", "small", "strong", "sub", "sup", "u", "caption", "col", "colgroup", "table", "tbody", "td", "tfoot", "th", "thead", "tr", "font", "nobr", "shadow", "tt")
	html := p.Sanitize(htmlString)
	for _, match := range EmptyHREFregex.FindAllString(html, -1) {
		if !URLregex.MatchString(match) {
			html = strings.Replace(html, match, "", -1)
		}
	}
	return html
}

