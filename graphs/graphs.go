package graphs

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

const (
	white Colour = iota
	grey
	black
)

// Colour represents the exploration status of a node.
type Colour uint8

// Graph represents an undirected graph.
type Graph struct {
	nodes map[uint64]*Node
}

// Node represents a node of a graph.
type Node struct {
	id     uint64
	colour Colour
	edges  []*Node
}

// New returns a Graph instance, populated according to the input string.
// Each input line represents an edge where U and V are Nodes of the graph:
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
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	nodes := make(map[uint64]*Node)
	var current *Node

	for scanner.Scan() {
		ids := strings.Fields(scanner.Text())
		nodeID, err := strconv.ParseUint(ids[0], 10, 64)
		if err != nil {
			return nil, errors.New("non-digit id on node")
		}
		current = &Node{id: nodeID}
		nodes[nodeID] = current

		if len(ids) == 2 {
			edgeID, err := strconv.ParseUint(ids[1], 10, 64)
			if err != nil {
				return nil, errors.New("non-digit id on node")
			}
			edge := &Node{id: edgeID}
			current.edges = append(current.edges, edge)
			nodes[edgeID] = edge
		}
	}
	return nodes, nil
}

// Colour a Node with either grey or black. Cannot transition back to an earlier colour.
func (n *Node) Colour() {
	if n.colour != black {
		n.colour++
	}
}

// Visit a Node identified by its id. It returns the list of Nodes connected to it and returns an error if the Node does not exist.
func (g *Graph) Visit(id uint64) ([]uint64, error) {
	if g.nodes[id] == nil {
		return nil, errors.New("node does not exist")
	}

	ret := []uint64{}
	for _, e := range g.nodes[id].edges {
		ret = append(ret, e.id)
	}
	return ret, nil
}
