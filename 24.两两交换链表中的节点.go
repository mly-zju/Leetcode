/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (62.31%)
 * Likes:    317
 * Dislikes: 0
 * Total Accepted:    48.4K
 * Total Submissions: 77.3K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
 *
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
func swapPairs(head *ListNode) *ListNode {
	// 1. 非递归解法
	if head == nil || head.Next == nil {
		return head
	}
	var nHead, prev, cur, next *ListNode = nil, nil, head, head.Next
	for next != nil {
		cur.Next = next.Next
		next.Next = cur
		if nHead == nil {
			nHead = next
			prev = cur 
		} else {
			// 重置prev的Next
			prev.Next = next
			prev = cur
		}

		// 移动2位，重置prev,cur和next
		cur = cur.Next
		if cur != nil {
			next = cur.Next
		} else {
			next = nil
		}
	}
	return nHead

	// 2. 递归解法
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	// next := head.Next
	// head.Next = swapPairs(next.Next)
	// next.Next = head
	// return next
}

// @lc code=end

