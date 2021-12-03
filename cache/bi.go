package cache

func minEatingSpeed(piles []int, H int) int {
	// piles 数组的最大值
	max := getMax(piles)
	speed := 1
	for speed < max {
		if canFinish(piles, speed, H) {
			return speed
		}
		speed += 1
	}

	return max
}

func getMax(piles []int) int {
	max := 0
	for i := 0; i < len(piles); i++ {
		if piles[i] > max {
			max = piles[i]
		}
	}

	return max
}

func canFinish(piles []int, speed int, H int) bool {
	time := 0
	for i := 0; i < len(piles); i++ {
		time += timeOf(piles[i], speed)
	}

	return time <= H
}

func timeOf(n int, speed int) int  {
	if n % speed == 0 {
		return n / speed
	}

	return (n / speed) + 1
}