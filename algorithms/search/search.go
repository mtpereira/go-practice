// Package search implements a search algorithms on top of graphs.
package search

import (
	"github.com/pkg/errors"

	"github.com/mtpereira/go-practice/graphs"
)

type graph interface {
	Node(uint64) *graphs.Node
	Visit(uint64) ([]uint64, error)
	Len() int
}

type algorithm func([]uint64) []uint64

// TODO: Add a struct for visiting nodes, allowing for colouring, setting distance and previous node.

// Search explores a Graph `g` using the algorithm `alg`, returning all discoverable nodes in a Graph format, starting from node `start`.
func Search(g graph, alg algorithm, start uint64) (graph, error) {
	distances := map[uint64]int{}
	distances[start] = 0
	queue := make(chan uint64)
	queue <- start

	d := 0
	for {
		select {
		case i := <-queue:
			edges, err := g.Visit(i)
			if err != nil {
				return nil, errors.Wrap(err, "error when visiting node")
			}
			d++

			for _, e := range edges {
				n := g.Node(e)
				if n.Colour == graphs.White {
					n.IncrementColour()
				}
				distances[e] = d
				queue <- e
			}
		default:
			var r graph
			r, err := graphs.New("")
			if err != nil {
				return nil, errors.Wrap(err, "error building resulting graph")
			}

			return r, nil
		}
	}
}

// BFS is an implementation of the Breadth-First Search algorithm, as described on the third edition of "Introduction to Algorithms" by Thomas H. Cormen et al.
func BFS() {}
