/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 方法1：先遍历一次计算长度l，再遍历第二次走l-k步返回
	// 方法2：快慢指针，快指针先走k步
	f, s := head, head
	// 快指针先走k步
	for i := 0; i < k; i++ {
		f = f.Next
	}
	for f != nil {
		s = s.Next
		f = f.Next
	}
	return s
}