// @Author: 2014BDuck
// @Date: 2020/10/17

package lfu

import (
	"github.com/2014BDuck/cache"
	"sort"
)

type entry struct {
	key    string
	value  interface{}
	weight int
	index  int
}

func (e *entry) Len() int {
	return cache.CalcLen(e.value) + 4 + 4
}

type queue []*entry

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].weight < q[j].weight
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *queue) Push(x interface{}) {
	n := len(*q)
	en := x.(*entry)
	en.index = n
	*q = append(*q, en)
}

func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	en := old[n-1]
	old[n-1] = nil
	en.index = -1
	*q = old[0 : n-1]
	return en
}