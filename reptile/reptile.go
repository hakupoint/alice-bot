package reptile

import (
	"errors"
	"net/url"
	"time"
)

var (
	NotFoundParse = errors.New("not found url parse rule !")
)

type Reptile struct {
	Base    *url.URL
	Rule    *Ruler
	Deep    int
	Failure []*ReptileFail
	done    func([]*ReptileFail)
	task    func() error
}

// 抓取失败结构体
type ReptileFail struct {
	Time *time.Time
	Url  *url.URL
	msg  error
}

func NewReptile(u interface{}) *Reptile {
	var uri *url.URL
	switch t := u.(type) {
	case string:
		uri, _ = url.Parse(t)
	case *url.URL:
		uri = t
	}
	return &Reptile{
		Base: uri,
	}
}

func (r *Reptile) SetRule(rule *Ruler) {
	r.Rule = rule
}

func (r *Reptile) SetDone(f func([]*ReptileFail)) {
	r.done = f
}

func (r *Reptile) Start() error {
	if r.Rule != nil {
		return NotFoundParse
	}
	return nil
}
func (r *Reptile) Fail(ra *ReptileFail) {

}

func (r *Reptile) NextPage() {
	defer func() {
		r.Deep--
	}()

	if r.Deep == 0 {
		r.done(r.Failure)
		return
	}

}
