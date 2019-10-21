package popular

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_popularWord(t *testing.T) {
	tests := map[string]struct {
		url   string
		word  []string
		count int
	}{
		"blog posts": {
			url: "https://jsonplaceholder.typicode.com/posts",
			word: []string{
				"aut",
				"eos",
				"est",
				"et",
				"non",
				"omnis",
				"provident",
				"qui",
				"quia",
				"quo",
				"rerum",
				"ut",
				"voluptas",
				"voluptatem",
			},
			count: 10,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ws, c, err := popularWord(tt.url)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(ws, tt.word); diff != "" {
				t.Fatalf(diff)
			}
			if diff := cmp.Diff(c, tt.count); diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
