package mincoins

import "sort"

func mincoins(cents int, coins []int) int {
	if cents < 1 {
		return 0
	}

	// guarantee coins are in desceinding order
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	// calculate min number of coins with all sets of coins that result from removing the largest coin, one by one
	var results []int
	for i := 0; i < len(coins); i++ {
		results = append(results, mincoinsAux(cents, 0, coins[i:]))
	}

	// return the smallest result
	sort.Sort(sort.IntSlice(results))
	return results[0]
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
