// 暂时用不到爬虫
package reptile

import (
	"errors"
	"net/url"
	"time"
)

var (
	NotFoundParse = errors.New("not found url parse rule !")
)

const (
	// 爬取的总页数
	TotalPageNumber = 10
)

type Reptile struct {
	Base    *url.URL
	Rule    *Ruler
	Total   int
	Failure []*ReptileFail
	Deep 		int
	done    func([]*ReptileFail)
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
		Deep: 10,
	}
}

func (r *Reptile) SetRule(rule *Ruler) {
	r.Rule = rule
}

func (r *Reptile) SetTotal(i int) {
	r.Total = i
}

func (r *Reptile) SetDone(f func([]*ReptileFail)) {
	r.done = f
}

func (r *Reptile) Start() error {
	if r.Rule != nil {
		return NotFoundParse
	}
	for {
		go func() {

		}()
		break;
	}


	return nil

}
func (r *Reptile) Fail(ra *ReptileFail) {

}
