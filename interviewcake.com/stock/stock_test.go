package stock

import "testing"

func Test_GetMaxProfit(t *testing.T) {
	tests := map[string]struct {
		stockPrices []int
		profit      int
	}{
		"1 price": {
			stockPrices: []int{10},
			profit:      0,
		},
		"2 prices": {
			stockPrices: []int{10, 7},
			profit:      0,
		},
		"example": {
			stockPrices: []int{10, 7, 5, 8, 11, 9},
			profit:      6,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if r := GetMaxProfit(tt.stockPrices); r != tt.profit {
				t.Fatalf("expected %v, got %v", tt.profit, r)
			}
		})
	}
}
