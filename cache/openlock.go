package cache

// https://leetcode-cn.com/problems/open-the-lock/
func openLock(deadends []string, target string) int {
	var step int
	var size int
	var queue []string
	visited := map[string]bool{}
	for _, d := range deadends {
		visited[d] = true
	}
	queue = append(queue, "0000")
	for len(queue) > 0 {
		size = len(queue)
		for i := 0; i < size; i++ {
			tmp := queue[0]
			// 出队
			queue = queue[1:]
			if visited[tmp] {
				continue
			}
			visited[tmp] = true
			if tmp == target {
				return step
			}
			for j := 0; j < 4; j++ {
				up := moveUp(tmp, j)
				if !visited[up] {
					queue = append(queue, up)
				}
				down := moveDown(tmp, j)
				if !visited[down] {
					queue = append(queue, down)
				}
			}
		}
		step++
	}
	return -1
}

func moveUp(str string, j int) string {
	arr := []byte(str)
	if arr[j] == '9' {
		arr[j] = '0'
	} else {
		arr[j] += 1
	}

	return string(arr)
}

func moveDown(str string, j int) string {
	arr := []byte(str)
	if arr[j] == '0' {
		arr[j] = '9'
	} else {
		arr[j] -= 1
	}

	return string(arr)
}

//https://leetcode-cn.com/problems/fibonacci-number/
func fib(n int) int {
	if n == 0 || n == 1{
		return n
	}
	a := 0
	b := 1
	c := 0
	for i := 0; i < n-1; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func Test(list []string, target string) int {
	return openLock(list, target)
}

func Test2(n int) int {
	return fib(n)
}