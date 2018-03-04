package reptile

import (
	"net/url"
)

type Ruler interface {
	Parse(url url.URL) error
}

type Reptiler interface {
	Start() error
	Done()
	NextPage()
	Fail()
	task()
}
