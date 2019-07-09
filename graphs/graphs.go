package graphs

import (
	"errors"
	"io"
	"strconv"
	"strings"
	"unicode"
)

const (
	white colour = iota
	grey
	black
)

type colour uint8

// Graph represents an undirected graph.
type Graph struct {
	nodes map[uint64]*Node
}

// Node represents a node of a graph.
type Node struct {
	id     uint64
	colour colour
	edges  []*Node
}

// New returns a Graph instance, populated according to the input string.
// The each input line represents an edge where U and V are nodes of the graph:
// `U [V]`
// As an example, the following lines represent the graph found at https://upload.wikimedia.org/wikipedia/commons/thumb/5/5b/6n-graf.svg/220px-6n-graf.svg.png . Notice the newline at the last line.
// ```
// 1 2
// 1 5
// 2 3
// 2 5
// 3 4
// 4 5
// 4 6
//
// ```
func New(input string) (*Graph, error) {
	nodes, err := readNodes(strings.NewReader(input))
	if err != nil {
		return nil, err
	}
	return &Graph{nodes: nodes}, nil
}

func readNodes(r io.Reader) (map[uint64]*Node, error) {
	input := make([]byte, 8)
	for {
		_, err := r.Read(input)
		if err == io.EOF {
			break
		}
	}

	nodes := make(map[uint64]*Node)
	var current *Node
	for _, c := range string(input) {
		switch {
		case unicode.IsDigit(c):
			id, err := strconv.ParseUint(string(c), 10, 64)
			if err != nil {
				return nil, err
			}

			if current == nil {
				current = &Node{id: id}
			} else {
				previous := current
				current = &Node{id: id}
				previous.edges = append(previous.edges, current)
			}

			nodes[id] = current
		// Space character except newline or EOF -> ignore.
		case unicode.IsSpace(c) && c != '\n' || c == 0:
			continue
		// Newline -> create a new node on next iteration.
		case c == '\n':
			current = nil
		default:
			return nil, errors.New("non-digit id on node")
		}
	}
	return nodes, nil
}
