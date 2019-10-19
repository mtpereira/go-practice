package mincoins

func mincoins(cents int) int {
	if cents < 1 {
		return 0
	}
	return mincoinsAux(cents, 0, []int{25, 10, 5, 1})
}

func mincoinsAux(cents int, count int, coins []int) int {
	// still got coins to compare to
	if len(coins) > 0 {
		coin := coins[0]
		// the current coin can be used at least once, add it then skip it
		if cents >= coin {
			remaining := cents % coin
			numberCoins := cents / coin
			return mincoinsAux(remaining, count+numberCoins, coins[1:])
		}
		// current coin isn't good for change, try again with the next coin
		return mincoinsAux(cents, count, coins[1:])
	}
	return count
}
