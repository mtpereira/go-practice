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
	Id     uint64
	Colour Colour
	Edges  []uint64
}

// NewNode returns a pointer to a new Node with the given ID.
func NewNode(id uint64) *Node { return &Node{Id: id} }

// IncrementColour sets the Colour of the Node one further closer to black. If it is black already, it does nothing.
func (n *Node) IncrementColour() {
	if n.Colour != Black {
		n.Colour++
	}
}

// AddEdge appends a Node to the list of Edges.
func (n *Node) AddEdge(id uint64) { n.Edges = append(n.Edges, id) }
