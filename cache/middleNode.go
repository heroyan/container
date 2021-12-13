package cache

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}


func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	flag1 := true
	flag2 := true
	for p1 != p2 {
		if p1 == nil {
			if flag1 {
				p1 = headB
				flag1 = false
			} else {
				return nil
			}
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			if flag2 {
				p2 = headA
				flag2 = false
			} else {
				return nil
			}
		} else {
			p2 = p2.Next
		}
	}

	return p1
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}

	return p1
}