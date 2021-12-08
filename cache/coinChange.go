package cache

import "fmt"

func coinChange(coins []int, amount int) int {
	visited := map[int]int{}
	return dp(coins, amount, visited)
}

func dp(coins []int, amount int, visited map[int]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	vv, ok := visited[amount]
	if ok {
		return vv
	}
	maxint := int(^uint(0) >> 1)  // ^uint(0) >> 1 = maxint
	res := maxint
	for _, v := range coins {
		sub := dp(coins, amount - v, visited)
		if sub == -1 {
			// 子问题无解，放弃
			continue
		}
		res = min(res, sub + 1)
	}

	if res == maxint {
		visited[amount] = -1
	} else {
		visited[amount] = res
	}

	return visited[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func coinChange2(coins []int, amount int) int {
	var dp []int
	for i := 0; i < amount+1; i++ {
		dp = append(dp, amount+1)
	}
	dp[0] = 0
	for i := 0; i < amount+1; i++ {
		for _, v := range coins {
			if i < v {
				continue
			}
			dp[i] = min(dp[i], dp[i-v] + 1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func TestCoin()  {
	fmt.Println(coinChange2([]int{1,2,5}, 11))
	fmt.Println(coinChange2([]int{1,2,5}, 15))
	fmt.Println(coinChange2([]int{2}, 3))
	fmt.Println(coinChange2([]int{1}, 2))
}