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
		colour colour
		edges  []*Node
	}
	type args struct {
		c colour
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				id:     tt.fields.id,
				colour: tt.fields.colour,
				edges:  tt.fields.edges,
			}
			if err := n.Colour(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Node.Colour() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
