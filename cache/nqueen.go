package cache

var visited [][]bool
var result [][]string

func solveNQueens(n int) [][]string {
	visited = make([][]bool, n)
	for idx, _ := range visited {
		visited[idx] = make([]bool, n)
	}

	result = make([][]string, 0)

	backtrace(n, 0)

	return result
}

func backtrace(n int, row int)  {
	if row == n {
		var ret []string
		var str string
		for _, v := range visited {
			str = ""
			for _, vv := range v {
				if vv {
					str += "Q"
				} else {
					str += "."
				}
			}
			ret = append(ret, str)
		}
		result = append(result, ret)
		return
	}
	for i := 0; i < n; i++ {
		// 左上有冲突，则返回
		if !upLeftOk(row, i) {
			continue
		}
		// 右上有冲突，则返回
		if !upRightOk(row, i, n) {
			continue
		}
		// 同列有冲突，则返回
		if !sameColOk(row, i) {
			continue
		}
		// 放
		visited[row][i] = true
		backtrace(n, row+1)
		// 撤
		visited[row][i] = false
	}
}

func sameColOk(row, col int) bool {
	for i := 0; i < row; i++ {
		if visited[i][col] {
			return false
		}
	}

	return true
}

func upLeftOk(row, col int) bool  {
	for row - 1 >= 0 && col - 1 >= 0 {
		if visited[row-1][col-1] {
			return false
		}
		row--
		col--
	}

	return true
}

func upRightOk(row, col, n int) bool  {
	for row - 1 >= 0 && col + 1 < n {
		if visited[row-1][col+1] {
			return false
		}
		row--
		col++
	}

	return true
}

func TestQueen(n int) [][]string {
	return solveNQueens(n)
}