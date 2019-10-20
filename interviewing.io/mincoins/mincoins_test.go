package mincoins

import "testing"

func Test_mincoins(t *testing.T) {
	tests := map[string]struct {
		cents int
		coins []int
		count int
	}{
		"0 cents, all coins": {
			cents: 0,
			coins: []int{25, 10, 5, 1},
			count: 0,
		},
		"1 cent, all coins": {
			cents: 1,
			coins: []int{25, 10, 5, 1},
			count: 1,
		},
		"2 cents, all coins": {
			cents: 2,
			coins: []int{25, 10, 5, 1},
			count: 2,
		},
		"5 cents, all coins": {
			cents: 5,
			coins: []int{25, 10, 5, 1},
			count: 1,
		},
		"6 cents, all coins": {
			cents: 6,
			coins: []int{25, 10, 5, 1},
			count: 2,
		},
		"31 cents, all coins": {
			cents: 31,
			coins: []int{25, 10, 5, 1},
			count: 3,
		},
		"33 cents, all coins": {
			cents: 33,
			coins: []int{25, 10, 5, 1},
			count: 5,
		},
		"31 cents, no 10 coins": {
			cents: 31,
			coins: []int{25, 5, 1},
			count: 3,
		},
		"33 cents, no 10 coins": {
			cents: 33,
			coins: []int{25, 5, 1},
			count: 5,
		},
		"31 cents, no 5 coins": {
			cents: 31,
			coins: []int{25, 10, 1},
			count: 4,
		},
		"60 cents, no 5 coins": {
			cents: 70,
			coins: []int{25, 10, 1},
			count: 4,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if r := mincoins(tt.cents, tt.coins); r != tt.count {
				t.Fatalf("expected %d, got %d", tt.count, r)
			}
		})
	}
}
