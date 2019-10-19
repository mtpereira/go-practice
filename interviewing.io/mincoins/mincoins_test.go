package mincoins

import "testing"

func Test_mincoins(t *testing.T) {
	tests := map[string]struct {
		input  int
		output int
	}{
		"0 cents": {
			input:  0,
			output: 0,
		},
		"1 cent": {
			input:  1,
			output: 1,
		},
		"2 cents": {
			input:  2,
			output: 2,
		},
		"5 cents": {
			input:  5,
			output: 1,
		},
		"6 cents": {
			input:  6,
			output: 2,
		},
		"31 cents": {
			input:  31,
			output: 3,
		},
		"33 cents": {
			input:  33,
			output: 5,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if r := mincoins(tt.input); r != tt.output {
				t.Fatalf("expected %d, got %d", tt.output, r)
			}
		})
	}
}
