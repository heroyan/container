package cache

func shipWithinDays(weights []int, D int) int {
	left := getMax1(weights)
	right := getSum(weights)
	for left < right {
		mid := left + (right-left)/2
		if canShipping(weights, mid, D) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func getMax1(piles []int) int {
	max := 0
	for i := 0; i < len(piles); i++ {
		if piles[i] > max {
			max = piles[i]
		}
	}

	return max
}

func getSum(piles []int) int {
	sum := 0
	for i := 0; i < len(piles); i++ {
		sum += piles[i]
	}

	return sum
}

func canShipping(weights []int, cap int, D int) bool {
	//按顺序来放，能放尽量放，看最少多少天能完成
	time := 1
	total := 0
	for i := 0; i < len(weights); i++ {
		total += weights[i]
		if total > cap {
			total = weights[i]
			time += 1
		}
	}

	return time <= D
}