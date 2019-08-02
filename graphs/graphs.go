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

// NodeID represents the ID of a Node.
type NodeID uint64

// Graph represents an undirected graph.
type Graph struct {
	nodes map[NodeID]*Node
}

// Node represents a node of a graph.
type Node struct {
	id     NodeID
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

func readNodes(r io.Reader) (map[NodeID]*Node, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	nodes := make(map[NodeID]*Node)
	var current *Node

	for scanner.Scan() {
		ids := strings.Fields(scanner.Text())
		id, err := strconv.ParseUint(ids[0], 10, 64)
		if err != nil {
			return nil, errors.New("non-digit id on node")
		}
		current = &Node{id: NodeID(id)}
		nodes[NodeID(id)] = current

		if len(ids) == 2 {
			e, err := strconv.ParseUint(ids[1], 10, 64)
			if err != nil {
				return nil, errors.New("non-digit id on node")
			}
			edge := &Node{id: NodeID(e)}
			current.edges = append(current.edges, edge)
			nodes[NodeID(e)] = edge
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
	for _, e := range n.edges {
		ret = append(ret, e.id)
	}
	return ret, nil
}

// Node returns a Node identified by its id. Returns `nil` if the Node does not exist.
func (g *Graph) Node(id NodeID) *Node {
	return g.nodes[NodeID(id)]
}

// Len returns the number of Nodes in the Graph.
func (g *Graph) Len() int { return len(g.nodes) }

// Colour returns the Node's colour.
func (n *Node) Colour() Colour { return n.colour }

// IncrementColour sets the Colour of the Node one further closer to black. If it is black already, it does nothing.
func (n *Node) IncrementColour() {
	if n.colour != black {
		n.colour++
	}
}

// ID returns the Node's id.
func (n *Node) ID() NodeID { return n.id }

// Edges returns the Node's edges.
func (n *Node) Edges() []*Node { return n.edges }
