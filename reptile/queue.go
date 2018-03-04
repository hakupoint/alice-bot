package reptile

import (
	"sync"
)

type queue struct {
	mu    sync.Mutex
	rule  []*Ruler
	Count int
}

func (q *queue) enqueue(r *Ruler) {
	defer q.mu.Unlock()
	q.mu.Lock()
	q.rule = append(q.rule, r)
	q.Count++
}

func (q *queue) dequeue() *Ruler {
	defer q.mu.Unlock()
	q.mu.Lock()
	r := q.rule[0]
	q.rule = q.rule[1:]
	return r
}
