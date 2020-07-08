/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (46.78%)
 * Likes:    323
 * Dislikes: 0
 * Total Accepted:    40.7K
 * Total Submissions: 82.1K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 * 
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 * 
 * 示例:
 * 
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 1. 迭代法
	// var prev, cur, next *ListNode = nil, head, head.Next
	// // 用于记录第一个反转节点
	// var reverseHead *ListNode
	// count := 1
	// for count <= n {
	// 	if count >= m && count <= n {
	// 		// 执行反转
	// 		cur.Next = prev
	// 		prev = cur
	// 		cur = next
	// 		if next != nil {
	// 			next = next.Next
	// 		}

	// 		if reverseHead == nil {
	// 			reverseHead = prev
	// 		}
	// 	} else {
	// 		prev = cur
	// 		cur = next
	// 		if next != nil {
	// 			next = next.Next
	// 		}
	// 	}
	// 	count++
	// }

	// // 需要区分第一个反转节点是头部，最后一个是尾部，正常三种情况
	// if reverseHead.Next == nil {
	// 	// 头部反转的情况
	// 	reverseHead.Next = cur
	// 	head = prev
	// } else {
	// 	oldPrev := reverseHead.Next
	// 	reverseHead.Next = cur
	// 	oldPrev.Next = prev
	// }
	// return head

	var prev, cur, next *ListNode = nil, head, head.Next
	// 用于记录第一个翻转的节点
	var reverseFirst *ListNode
	count := 1
	for count <= n {
		if count >= m {
			// 执行反转
			cur.Next = prev
			prev = cur
			cur = next
			if next != nil {
				next = next.Next
			}

			// 记录第一个反转的节点
			if reverseFirst == nil {
				reverseFirst = prev
			}

		} else {
			prev = cur
			cur = next
			next = next.Next
		}
		count++
	}

	// 区分reverseFirst是否是头部
	if reverseFirst.Next == nil {
		reverseFirst.Next = cur
		head = prev
	} else {
		oldPrev := reverseFirst.Next
		reverseFirst.Next = cur
		oldPrev.Next = prev
	}

	return head
}
// @lc code=end

