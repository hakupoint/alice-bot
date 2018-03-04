package reptile

import (
	"sync"
)

type queue struct {
	mu    sync.Mutex
	rule  []*Rule
	Count int
}

func (q *queue) enqueue(r *Rule) {
	defer q.mu.Unlock()
	q.mu.Lock()
	q.rule = append(q.rule, r)
	q.Count++
}

func (q *queue) dequeue() *Rule {
	defer q.mu.Unlock()
	q.mu.Lock()
	r := q.rule[0]
	q.rule = q.rule[1:]
	return r
}
