package search

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mtpereira/go-practice/graphs"
)

func TestQueue_Len(t *testing.T) {
	tests := map[string]struct {
		q    Queue
		want int
	}{
		"empty": {
			q:    Queue{},
			want: 0,
		},
		"1 node": {
			q:    Queue{&graphs.Node{Id: 1}},
			want: 1,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.q.Len()
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestQueue_Less(t *testing.T) {
	tests := map[string]struct {
		q    Queue
		i    int
		j    int
		want bool
	}{
		"equals": {
			q:    Queue{&graphs.Node{Id: 1}, &graphs.Node{Id: 1}},
			i:    0,
			j:    1,
			want: false,
		},
		"lesser": {
			q:    Queue{&graphs.Node{Id: 1}, &graphs.Node{Id: 2}},
			i:    0,
			j:    1,
			want: true,
		},
		"greater": {
			q:    Queue{&graphs.Node{Id: 1}, &graphs.Node{Id: 2}},
			i:    1,
			j:    0,
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.q.Less(tt.i, tt.j)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestQueue_Swap(t *testing.T) {
	tests := map[string]struct {
		q    Queue
		i    int
		j    int
		want Queue
	}{
		"equals": {
			q:    Queue{&graphs.Node{Id: 1}, &graphs.Node{Id: 1}},
			i:    0,
			j:    1,
			want: false,
		},
		"lesser": {
			q:    Queue{&graphs.Node{Id: 1}, &graphs.Node{Id: 2}},
			i:    0,
			j:    1,
			want: true,
		},
		"greater": {
			q:    Queue{&graphs.Node{Id: 1}, &graphs.Node{Id: 2}},
			i:    1,
			j:    0,
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.q.Less(tt.i, tt.j)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		q    *Queue
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Push(tt.args.x)
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
