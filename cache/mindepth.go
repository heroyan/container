package cache

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	var size int
	var depth int
	var tmp *TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		depth++
		size = len(queue)
		for i := 0; i < size; i++ {
			tmp = queue[0]
			if tmp.Left == nil && tmp.Right == nil {
				return depth
			}
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
			// 相当于出队
			queue = queue[1:]
		}
	}

	return depth
}