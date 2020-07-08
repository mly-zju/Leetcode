/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	// 反转
	l, r := 0, len(res) - 1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}