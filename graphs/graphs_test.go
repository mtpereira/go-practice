package graphs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	node42 = &Node{
		Id:     42,
		Colour: White,
		Edges:  nil,
	}
	node42Grey = &Node{
		Id:     42,
		Colour: Grey,
		Edges:  nil,
	}
	node42Black = &Node{
		Id:     42,
		Colour: Grey,
		Edges:  nil,
	}
	node1 = &Node{
		Id:     1,
		Colour: White,
		Edges:  nil,
	}
	node2 = &Node{
		Id:     2,
		Colour: White,
		Edges:  nil,
	}
	node3 = &Node{
		Id:     3,
		Colour: White,
		Edges:  nil,
	}
	node1_2 = &Node{
		Id:     1,
		Colour: White,
		Edges:  []uint64{node2.Id},
	}
	node2_3 = &Node{
		Id:     2,
		Colour: White,
		Edges:  []uint64{node3.Id},
	}
	node1_2_3 = &Node{
		Id:     1,
		Colour: White,
		Edges:  []uint64{node2.Id, node3.Id},
	}
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    *Graph
		wantErr bool
	}{
		"empty": {
			input: "",
			want:  &Graph{nodes: map[uint64]*Node{}},
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
				map[uint64]*Node{
					uint64(1): node1,
				},
			},
		},
		"two nodes, one edge": {
			input: "1 2\n",
			want: &Graph{
				map[uint64]*Node{
					uint64(1): node1_2,
					uint64(2): node2,
				},
			},
		},
		"three nodes, two edges": {
			input: "1 2\n2 3\n",
			want: &Graph{
				map[uint64]*Node{
					uint64(1): node1_2,
					uint64(2): node2_3,
					uint64(3): node3,
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
	tests := map[string]struct {
		node   *Node
		colour Colour
	}{
		"colour from white to grey": {
			node:   node42,
			colour: Grey,
		},
		"colour from grey to black": {
			node:   node42Grey,
			colour: Black,
		},
		"do nothing when colouring from black": {
			node:   node42Black,
			colour: Black,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			n := &Node{
				Id:     tt.node.Id,
				Colour: tt.node.Colour,
				Edges:  tt.node.Edges,
			}
			n.IncrementColour()
			diff := cmp.Diff(tt.colour, n.Colour, cmp.AllowUnexported(Graph{}, Node{}))
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestGraph_Visit(t *testing.T) {
	tests := map[string]struct {
		graph   *Graph
		input   uint64
		want    []uint64
		wantErr bool
	}{
		"1 node, visit node 2": {
			graph: &Graph{
				nodes: map[uint64]*Node{
					uint64(1): node1,
				},
			},
			input:   42,
			wantErr: true,
		},
		"1 node, visit node 1": {
			graph: &Graph{
				nodes: map[uint64]*Node{
					uint64(1): node1,
				},
			},
			input: 1,
			want:  []uint64{},
		},
		"2 nodes, visit node 1": {
			graph: &Graph{
				map[uint64]*Node{
					uint64(1): node1_2,
					uint64(2): node2,
				},
			},
			input: 1,
			want:  []uint64{2},
		},
		"3 nodes, visit node 1": {
			graph: &Graph{
				map[uint64]*Node{
					uint64(1): node1_2_3,
					uint64(2): node2,
					uint64(3): node3,
				},
			},
			input: 1,
			want:  []uint64{2, 3},
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
	tests := map[string]struct {
		graph *Graph
		input uint64
		want  *Node
	}{
		"get node 42, return nil": {
			graph: &Graph{
				nodes: map[uint64]*Node{
					uint64(1): node1,
				},
			},
			input: 42,
			want:  nil,
		},
		"get node 1": {
			graph: &Graph{
				nodes: map[uint64]*Node{
					uint64(1): node1,
				},
			},
			input: 1,
			want:  node1,
		},
		"get node 2": {
			graph: &Graph{
				map[uint64]*Node{
					uint64(1): node1,
					uint64(2): node2,
				},
			},
			input: 2,
			want:  node2,
		},
		"get node 3": {
			graph: &Graph{
				map[uint64]*Node{
					uint64(1): node1,
					uint64(2): node2,
					uint64(3): node3,
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
	tests := map[string]struct {
		graph *Graph
		want  int
	}{
		"get length of a 1 node graph": {
			graph: &Graph{
				nodes: map[uint64]*Node{
					uint64(1): node1,
				},
			},
			want: 1,
		},
		"get length of a 2 node graph": {
			graph: &Graph{
				map[uint64]*Node{
					uint64(1): node1,
					uint64(2): node2,
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
