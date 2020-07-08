/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (47.80%)
 * Likes:    219
 * Dislikes: 0
 * Total Accepted:    52.5K
 * Total Submissions: 109K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 *
 * 示例 1:
 *
 * 输入: 1->1->2
 * 输出: 1->2
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->2->3->3
 * 输出: 1->2->3
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
func deleteDuplicates(head *ListNode) *ListNode {
	// tmp := head
	// for tmp != nil && tmp.Next != nil {
	// 	// 要注意只有不重复时候才进位
	// 	if tmp.Val == tmp.Next.Val {
	// 		tmp.Next = tmp.Next.Next
	// 	} else {
	// 		tmp = tmp.Next
	// 	}
	// }
	// return head

	tmp := head
	for tmp != nil && tmp.Next != nil {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}

// @lc code=end

