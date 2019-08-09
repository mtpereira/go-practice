// Package bfs implements a Breadth-First Search algorithm, as described on the third edition of "Introduction to Algorithms" by Thomas H. Cormen et al.
package search

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mtpereira/go-practice/graphs"
)

func TestSearch(t *testing.T) {
	emptyGraph, _ := graphs.New("")
	tests := map[string]struct {
		graph   graph
		alg     algorithm
		start   uint64
		want    graph
		wantErr bool
	}{
		"empty graph": {
			graph: emptyGraph,
			alg:   func(id []uint64) []uint64 { return []uint64{} },
			start: 0,
			want:  emptyGraph,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Search(tt.graph, tt.alg, tt.start)
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error, got none")
			}
			diff := cmp.Diff(tt.want, got, cmp.AllowUnexported(graphs.Graph{}))
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
