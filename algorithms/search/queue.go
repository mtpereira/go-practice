package search

import "github.com/mtpereira/go-practice/graphs"

type Queue []*graphs.Node

func (q Queue) Len() int { return len(q) }

func (q Queue) Less(i, j int) bool { return q[i].Id < q[j].Id }

func (q Queue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q *Queue) Push(x interface{}) { *q = append(*q, x.(*graphs.Node)) }

func (q *Queue) Pop() interface{} {
	node := (*q)[len(*q)-1]
	*q = (*q)[0 : len(*q)-1]
	return node
}
