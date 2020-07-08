/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (35.56%)
 * Likes:    565
 * Dislikes: 0
 * Total Accepted:    86.9K
 * Total Submissions: 240.6K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 *
 * 示例：
 *
 * 给定一个链表: 1->2->3->4->5, 和 n = 2.
 *
 * 当删除了倒数第二个节点后，链表变为 1->2->3->5.
 *
 *
 * 说明：
 *
 * 给定的 n 保证是有效的。
 *
 * 进阶：
 *
 * 你能尝试使用一趟扫描实现吗？
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1. 两趟扫描法
	// length := 0
	// tmp := head
	// for tmp != nil {
	// 	length++
	// 	tmp = tmp.Next
	// }

	// gap := length - n
	// tmp = head
	// // 如果是删除头部元素，则区别处理
	// if gap == 0 {
	// 	head = head.Next
	// } else {
	// 	// 寻找前一个元素
	// 	for gap > 1 {
	// 		tmp = tmp.Next
	// 		gap--
	// 	}
	// 	tmp.Next = tmp.Next.Next
	// }
	// return head

	// 2. 一趟扫描法
	s, f := head, head
	for i := 0; i < n; i++ {
		// f先多遍历n位，让二者相隔n
		f = f.Next
	}

	if f == nil {
		// 如果f直接nil，说明删除头部
		head = head.Next
	} else {
		for f.Next != nil {
			// 当f到最后一个元素，s正好到要删除的前一个元素
			f = f.Next
			s = s.Next
		}

		s.Next = s.Next.Next
	}

	return head
}

// @lc code=end

