package graphs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    *Graph
		wantErr bool
	}{
		"empty": {
			input: "",
			want:  &Graph{nodes: map[NodeID]*Node{}},
		},
		"non-digit id on node": {
			input:   "a 1\n",
			wantErr: true,
		},
		"non-digit id on edge": {
			input:   "1 a\n",
			wantErr: true,
		},
		"one node, no edges": {
			input: "1\n",
			want: &Graph{
				map[NodeID]*Node{
					NodeID(1): &Node{
						id:     1,
						colour: white,
						edges:  nil,
					},
				},
			},
		},
		"two nodes, one edge": {
			input: "1 2\n",
			want: &Graph{
				map[NodeID]*Node{
					NodeID(1): &Node{
						id:     1,
						colour: white,
						edges: []*Node{
							&Node{
								id:     2,
								colour: white,
								edges:  nil,
							},
						},
					},
					NodeID(2): &Node{
						id:     2,
						colour: white,
						edges:  nil,
					},
				},
			},
		},
		"three nodes, two edges": {
			input: "1 2\n2 3\n",
			want: &Graph{
				map[NodeID]*Node{
					NodeID(1): &Node{
						id:     1,
						colour: white,
						edges: []*Node{
							&Node{
								id:     2,
								colour: white,
								edges:  nil,
							},
						},
					},
					NodeID(2): &Node{
						id:     2,
						colour: white,
						edges: []*Node{
							&Node{
								id:     3,
								colour: white,
								edges:  nil,
							},
						},
					},
					NodeID(3): &Node{
						id:     3,
						colour: white,
						edges:  nil,
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := New(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error, got none")
			}
			diff := cmp.Diff(tt.want, got, cmp.AllowUnexported(Graph{}, Node{}))
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestNode_IncrementColour(t *testing.T) {
	type fields struct {
		id     NodeID
		colour Colour
		edges  []*Node
	}
	tests := map[string]struct {
		node   *Node
		colour Colour
	}{
		"colour from white to grey": {
			node: &Node{
				id:     42,
				colour: white,
				edges:  nil,
			},
			colour: grey,
		},
		"colour from grey to black": {
			node: &Node{
				id:     42,
				colour: grey,
				edges:  nil,
			},
			colour: black,
		},
		"do nothing when colouring from black": {
			node: &Node{
				id:     42,
				colour: black,
				edges:  nil,
			},
			colour: black,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			n := &Node{
				id:     tt.node.id,
				colour: tt.node.colour,
				edges:  tt.node.edges,
			}
			n.IncrementColour()
			diff := cmp.Diff(tt.colour, n.colour, cmp.AllowUnexported(Graph{}, Node{}))
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestGraph_Visit(t *testing.T) {
	tests := map[string]struct {
		graph   *Graph
		input   NodeID
		want    []NodeID
		wantErr bool
	}{
		"1 node, visit node 2": {
			graph: &Graph{
				nodes: map[NodeID]*Node{
					NodeID(1): &Node{
						id:     1,
						colour: white,
						edges:  nil,
					},
				},
			},
			input:   42,
			wantErr: true,
		},
		"1 node, visit node 1": {
			graph: &Graph{
				nodes: map[NodeID]*Node{
					NodeID(1): &Node{
						id:     1,
						colour: white,
						edges:  nil,
					},
				},
			},
			input: 1,
			want:  []NodeID{},
		},
		"2 nodes, visit node 1": {
			graph: &Graph{
				map[NodeID]*Node{
					NodeID(1): &Node{
						id:     1,
						colour: white,
						edges: []*Node{
							&Node{
								id:     2,
								colour: white,
								edges:  nil,
							},
						},
					},
					NodeID(2): &Node{
						id:     2,
						colour: white,
						edges:  nil,
					},
				},
			},
			input: 1,
			want:  []NodeID{2},
		},
		"3 nodes, visit node 1": {
			graph: &Graph{
				map[NodeID]*Node{
					NodeID(1): &Node{
						id:     1,
						colour: white,
						edges: []*Node{
							&Node{
								id:     2,
								colour: white,
								edges:  nil,
							},
							&Node{
								id:     3,
								colour: white,
								edges:  nil,
							},
						},
					},
					NodeID(2): &Node{
						id:     2,
						colour: white,
						edges:  nil,
					},
					NodeID(3): &Node{
						id:     3,
						colour: white,
						edges:  nil,
					},
				},
			},
			input: 1,
			want:  []NodeID{2, 3},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.graph.Visit(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error, got none")
			}
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestGraph_Node(t *testing.T) {
	node1 := &Node{
		id:     1,
		colour: white,
		edges:  nil,
	}
	node2 := &Node{
		id:     2,
		colour: white,
		edges:  nil,
	}
	node3 := &Node{
		id:     1,
		colour: white,
		edges:  []*Node{node1, node2},
	}
	tests := map[string]struct {
		graph *Graph
		input NodeID
		want  *Node
	}{
		"get node 42, return nil": {
			graph: &Graph{
				nodes: map[NodeID]*Node{
					NodeID(1): node1,
				},
			},
			input: 42,
			want:  nil,
		},
		"get node 1": {
			graph: &Graph{
				nodes: map[NodeID]*Node{
					NodeID(1): node1,
				},
			},
			input: 1,
			want:  node1,
		},
		"get node 2": {
			graph: &Graph{
				map[NodeID]*Node{
					NodeID(1): node1,
					NodeID(2): node2,
				},
			},
			input: 2,
			want:  node2,
		},
		"get node 3": {
			graph: &Graph{
				map[NodeID]*Node{
					NodeID(1): node1,
					NodeID(2): node2,
					NodeID(3): node3,
				},
			},
			input: 3,
			want:  node3,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.graph.Node(tt.input)
			diff := cmp.Diff(tt.want, got, cmp.AllowUnexported(Graph{}, Node{}))
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestGraph_Len(t *testing.T) {
	node1 := &Node{
		id:     1,
		colour: white,
		edges:  nil,
	}
	node2 := &Node{
		id:     2,
		colour: white,
		edges:  nil,
	}
	tests := map[string]struct {
		graph *Graph
		want  int
	}{
		"get length of a 1 node graph": {
			graph: &Graph{
				nodes: map[NodeID]*Node{
					NodeID(1): node1,
				},
			},
			want: 1,
		},
		"get length of a 2 node graph": {
			graph: &Graph{
				map[NodeID]*Node{
					NodeID(1): node1,
					NodeID(2): node2,
				},
			},
			want: 2,
		},
		"get length of an empty graph": {
			graph: &Graph{},
			want:  0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.graph.Len()
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
func TestNode_Colour(t *testing.T) {
	type fields struct {
		id     NodeID
		colour Colour
		edges  []*Node
	}
	tests := map[string]struct {
		node   *Node
		colour Colour
	}{
		"get colour from a node": {
			node: &Node{
				id:     42,
				colour: white,
				edges:  nil,
			},
			colour: white,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			n := &Node{
				id:     tt.node.id,
				colour: tt.node.colour,
				edges:  tt.node.edges,
			}
			n.Colour()
			diff := cmp.Diff(tt.colour, n.colour, cmp.AllowUnexported(Graph{}, Node{}))
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestNode_ID(t *testing.T) {
	type fields struct {
		id     NodeID
		colour Colour
		edges  []*Node
	}
	tests := map[string]struct {
		node  *Node
		edges []*Node
	}{
		"get ID from a Node": {
			node: &Node{
				id:     1,
				colour: white,
				edges: []*Node{
					&Node{
						id:     2,
						colour: white,
						edges:  nil,
					},
					&Node{
						id:     3,
						colour: white,
						edges:  nil,
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			n := &Node{
				id:     tt.node.id,
				colour: tt.node.colour,
				edges:  tt.node.edges,
			}
			diff := cmp.Diff(tt.node.id, n.ID(), cmp.AllowUnexported(Graph{}, Node{}))
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestNode_Edges(t *testing.T) {
	type fields struct {
		id     NodeID
		colour Colour
		edges  []*Node
	}
	tests := map[string]struct {
		node  *Node
		edges []*Node
	}{
		"get the edges of a Node": {
			node: &Node{
				id:     1,
				colour: white,
				edges: []*Node{
					&Node{
						id:     2,
						colour: white,
						edges:  nil,
					},
					&Node{
						id:     3,
						colour: white,
						edges:  nil,
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			n := &Node{
				id:     tt.node.id,
				colour: tt.node.colour,
				edges:  tt.node.edges,
			}
			n.Edges()
			diff := cmp.Diff(tt.node.edges, n.Edges(), cmp.AllowUnexported(Graph{}, Node{}))
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
