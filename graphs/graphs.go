package graphs

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type Node interface {
	ID() NodeID
	Edges() []NodeID
	Colour() Colour
	IncrementColour()
	AddEdge(NodeID)
}

// Graph represents an undirected graph.
type Graph struct {
	nodes map[NodeID]Node
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

func readNodes(r io.Reader) (map[NodeID]Node, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	nodes := make(map[NodeID]Node)
	var current Node

	for scanner.Scan() {
		ids := strings.Fields(scanner.Text())
		id, err := strconv.ParseUint(ids[0], 10, 64)
		if err != nil {
			return nil, errors.New("non-digit id on node")
		}
		current = NewNode(NodeID(id))
		nodes[NodeID(id)] = current

		if len(ids) == 2 {
			e, err := strconv.ParseUint(ids[1], 10, 64)
			if err != nil {
				return nil, errors.New("non-digit id on node")
			}
			edge := NewNode(NodeID(e))
			current.AddEdge(edge.ID())
			nodes[NodeID(e)] = Node(edge)
		}
	}
	return nodes, nil
}

// Visit a Node identified by its id. It returns the list of Nodes connected to it and returns an error if the Node does not exist.
func (g *Graph) Visit(id NodeID) ([]NodeID, error) {
	n := g.nodes[id]
	if n == nil {
		return nil, errors.New("node does not exist")
	}

	ret := []NodeID{}
	for _, e := range n.Edges() {
		ret = append(ret, e)
	}
	return ret, nil
}

// Node returns a Node identified by its id. Returns `nil` if the Node does not exist.
func (g *Graph) Node(id NodeID) Node {
	return g.nodes[NodeID(id)]
}

// Len returns the number of Nodes in the Graph.
func (g *Graph) Len() int { return len(g.nodes) }
