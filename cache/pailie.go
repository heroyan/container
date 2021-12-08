package cache

import "fmt"

// Pailie 计算全排列
func Pailie(nums []int) {
	track := []int{}
	visited := map[int]bool{}
	huishuo(track, nums, visited)
}

func huishuo(track []int, nums []int, visited map[int]bool)  {
	if len(track) == len(nums) {
		fmt.Println(track)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}
		// 加进去
		visited[nums[i]] = true
		track = append(track, nums[i])
		huishuo(track, nums, visited)
		// 退出来
		track = track[:len(track)-1]
		visited[nums[i]] = false
	}
}
