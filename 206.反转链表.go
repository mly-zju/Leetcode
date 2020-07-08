/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (64.49%)
 * Likes:    625
 * Dislikes: 0
 * Total Accepted:    113K
 * Total Submissions: 173.2K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
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
func reverseList(head *ListNode) *ListNode {
	// 1. 迭代
	// if head == nil {
	// 	return head
	// }
	// var prev, cur, next *ListNode = nil, head, head.Next
	// for next != nil {
	// 	cur.Next = prev
	// 	prev = cur
	// 	cur = next
	// 	next = next.Next
	// }
	// // 结束循环后再处理一次
	// cur.Next = prev
	// return cur

	// 2. 递归
	return dfs(nil, head)
}

func dfs(prev, cur *ListNode) *ListNode {
	// cur为nil，递归结束
	if cur == nil {
		return prev
	}

	next := cur.Next
	cur.Next = prev
	prev = cur
	cur = next
	return dfs(prev, cur)
}


// @lc code=end

