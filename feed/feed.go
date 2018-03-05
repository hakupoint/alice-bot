package feed
import (
	"time"
	"text/template"
)

const (
	ParseError = "Parsing failed !"
	
)
type Feed struct {
	MessageTamp *template.Template
	BaseUrl     string
	createTime  time.Time
}

func NewFeed(url string, temp *template.Template) *Feed{
	return &Feed{
		BaseUrl: url,
		MessageTamp: temp,
		createTime: time.Now(),
	}
}
func (f *Feed) Pull(fn func(s string) error) {

}

func (f *Feed) string() {
	
}