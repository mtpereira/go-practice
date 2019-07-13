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
			want:  &Graph{nodes: map[uint64]*Node{}},
		},
		"non-digit id on node": {
			input:   "a 1\n",
			wantErr: true,
		},
		"one node, no edges": {
			input: "1\n",
			want: &Graph{
				map[uint64]*Node{
					uint64(1): &Node{
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
				map[uint64]*Node{
					uint64(1): &Node{
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
					uint64(2): &Node{
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
				map[uint64]*Node{
					uint64(1): &Node{
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
					uint64(2): &Node{
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
					uint64(3): &Node{
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
				t.Errorf("Node.Colour() error = %v, wantErr %v", err, tt.wantErr)
			}
			diff := cmp.Diff(tt.want, got, cmp.AllowUnexported(Graph{}, Node{}))
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestNode_Colour(t *testing.T) {
	type fields struct {
		id     uint64
		colour Colour
		edges  []*Node
	}
	tests := map[string]struct {
		node    *Node
		colour  Colour
		wantErr bool
	}{
		"colour from white to grey": {
			node: &Node{
				id:     42,
				colour: white,
				edges:  nil,
			},
			colour: grey,
			wantErr: true,
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
			n.Colour()
			diff := cmp.Diff(tt.colour, n.colour, cmp.AllowUnexported(Graph{}, Node{}))
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
