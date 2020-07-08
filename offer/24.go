/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 1. 递归最简单
	// if head == nil {
	// 	return nil
	// }
	// var prev *ListNode
	// return _reverse(prev, head)

	// 2. 迭代
	if head == nil {
		return nil
	}
	var prev, cur, next *ListNode = nil, head, head.Next
	for next != nil {
		cur.Next = prev
		prev = cur
		cur = next
		next = next.Next
	}
	cur.Next = prev
	return cur
}

func _reverse(prev, next *ListNode) *ListNode {
	if next.Next == nil {
		next.Next = prev
		return next
	}

	newNext := next.Next
	next.Next = prev
	prev = next
	next = newNext
	return _reverse(prev, next)
}