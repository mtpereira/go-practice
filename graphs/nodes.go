package graphs

// Colour represents the exploration status of a node.
type Colour uint8

// White, Grey and Black are the possible Colours of a Node.
const (
	White Colour = iota
	Grey
	Black
)

// Node represents a node of a graph.
type Node struct {
	id     uint64
	colour Colour
	edges  []uint64
}

// NewNode returns a pointer to a new Node with the given ID.
func NewNode(id uint64) *Node { return &Node{id: id} }

// Colour returns the Node's colour.
func (n *Node) Colour() Colour { return n.colour }

// IncrementColour sets the Colour of the Node one further closer to black. If it is black already, it does nothing.
func (n *Node) IncrementColour() {
	if n.colour != Black {
		n.colour++
	}
}

// AddEdge appends a Node to the list of Edges.
func (n *Node) AddEdge(id uint64) { n.edges = append(n.edges, id) }

