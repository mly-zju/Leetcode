/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 分别计算二者长度，找到差值，然后快慢指针直到第一个相等为止
	l1, l2 := 0, 0
	curA, curB := headA, headB
	for curA != nil {
		l1++
		curA = curA.Next
	}
	for curB != nil {
		l2++
		curB = curB.Next
	}
	
	// 计算差值
	curA, curB = headA, headB
	if l1 > l2 {
		for i := 0; i < l1 - l2; i++ {
			curA = curA.Next
		}
	} else if l2 > l1 {
		for i := 0; i < l2 - l1; i++ {
			curB = curB.Next
		}
	}
	// 开始共同前进
	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		} else {
			curA = curA.Next
			curB = curB.Next
		}
	}
	return nil
}