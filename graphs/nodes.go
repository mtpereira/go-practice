package graphs

// Colour represents the exploration status of a node.
type Colour uint8

// White, Grey and Black are the possible Colours of a Node.
const (
	White Colour = iota
	Grey
	Black
)

// NodeID represents the ID of a Node.
type NodeID uint64

// Node represents a node of a graph.
type node struct {
	id     NodeID
	colour Colour
	edges  []NodeID
}

// NewNode returns a pointer to a new Node with the given ID.
func NewNode(id NodeID) *node { return &node{id: id} }

// Colour returns the Node's colour.
func (n *node) Colour() Colour { return n.colour }

// IncrementColour sets the Colour of the Node one further closer to black. If it is black already, it does nothing.
func (n *node) IncrementColour() {
	if n.colour != Black {
		n.colour++
	}
}

// ID returns the Node's id.
func (n *node) ID() NodeID { return n.id }

// Edges returns the Node's edges.
func (n *node) Edges() []NodeID { return n.edges }

// AddEdge appends a Node to the list of Edges.
func (n *node) AddEdge(id NodeID) { n.edges = append(n.edges, id) }
