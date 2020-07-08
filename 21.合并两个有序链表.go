/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (60.55%)
 * Likes:    1028
 * Dislikes: 0
 * Total Accepted:    260.1K
 * Total Submissions: 417.7K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
 * 
 * 示例：
 * 
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	// 1. 新创建
	// var newHead, cur *ListNode
	// for l1 != nil && l2 != nil {
	// 	var curNode *ListNode
	// 	if l1.Val <= l2.Val {
	// 		curNode = l1
	// 		l1 = l1.Next
	// 	} else {
	// 		curNode = l2
	// 		l2 = l2.Next
	// 	}

	// 	if newHead == nil {
	// 		newHead = curNode
	// 		cur = newHead
	// 	} else {
	// 		cur.Next = curNode
	// 		cur = cur.Next
	// 	}
	// }

	// if l1 != nil {
	// 	cur.Next = l1
	// }
	// if l2 != nil {
	// 	cur.Next = l2
	// }
	// return newHead

	// 2. 原地修改, 以l1为基准
	var newHead, prev *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			if newHead == nil {
				newHead = l1
			}
			prev = l1
			l1 = l1.Next
		} else {
			if newHead == nil {
				newHead = l2
				prev = l2
				l2 = l2.Next
				prev.Next = l1
			} else {
				// 将该l2节点插入prev和l1之间
				ol2 := l2
				l2 = l2.Next
				prev.Next = ol2
				ol2.Next = l1
				prev = ol2
			}
		}
	}

	// 如果l1已空, l2不空，其余都接入l2
	if l1 == nil && l2 != nil {
		prev.Next = l2
	}
	return newHead
}
// @lc code=end

