/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	// 这个只删除一个节点，还需要考虑删除多个节点的情况，方法通用的
	var newHead, prev *ListNode
	// 先处理头部
	for head.Val == val {
		head = head.Next
	}

	// 初始化新的头部
	newHead = head
	prev = head
	head = head.Next
	for head != nil {
		if head.Val == val {
			head = head.Next
			prev.Next = head
		} else {
			prev = head
			head = head.Next
		}
	}
	return newHead
}