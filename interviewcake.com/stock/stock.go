package stock

// GetMaxProfit returns the best profit that could have been made given the stock prices at previous times.
// `stockPrices` indexes represent minutes, starting at 9:30am and their values represent the stock price.
func GetMaxProfit(stockPrices []int) int {
	profit := 0
	for t, p := range stockPrices {
		for _, s := range stockPrices[t+1:] {
			currentProfit := s - p
			if profit < currentProfit {
				profit = currentProfit
			}
		}
	}
	return profit
}
