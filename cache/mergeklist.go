package cache

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 小顶堆
import (
	"container/heap"
)

type ListNodeHeap []*ListNode
func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListNodeHeap{}
	for _, node := range lists {
		if node == nil {
			continue
		}
		h.Push(node)
	}
	heap.Init(h)
	dumy := &ListNode{
		Val: 0,
		Next: nil,
	}
	p := dumy
	for h.Len() > 0 {
		x := heap.Pop(h).(*ListNode)
		if x == nil {
			continue
		}
		p.Next = x
		if x.Next != nil {
			heap.Push(h, x.Next)
		}
		p = p.Next
	}

	return dumy.Next
}

func (obj ListNodeHeap) Len() int {
	return len(obj)
}

func (obj ListNodeHeap) Less(i, j int) bool {
	if obj[i] == nil || obj[j] == nil {
		return true
	}
	return obj[i].Val < obj[j].Val
}

func (obj ListNodeHeap) Swap(i, j int) {
	obj[i], obj[j] = obj[j], obj[i]
}

func (h *ListNodeHeap) Pop() interface{} {
	obj := *h
	n := len(obj)
	x := obj[n-1]
	*h = obj[0:n-1]

	return x
}

func (h *ListNodeHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(*ListNode))
}