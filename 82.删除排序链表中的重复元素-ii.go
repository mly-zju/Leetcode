/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (43.79%)
 * Likes:    159
 * Dislikes: 0
 * Total Accepted:    22.3K
 * Total Submissions: 50.2K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->1->2->3
 * 输出: 2->3
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
	if head == nil {
		return head
	}

	// 快慢指针+基准指针
	s, f, cur := head, head, head
	// 还要记录prev，作为最后置nil
	var prev *ListNode

	for f != nil {
		if f.Val == s.Val {
			f = f.Next
		} else {
			// 如果不相等，看看s和f之间是否相隔超过1，如果是说明有重复的
			if s.Next == f {
				cur.Val = s.Val
				prev = cur
				cur = cur.Next
			}
			s = f
		}
	}
	// 最终结束后还要再看一下s是否位于最后一位
	if s.Next == nil {
		cur.Val = s.Val
		prev = cur
		cur = cur.Next
	}

	// 检测是否为空
	if prev == nil {
		return prev
	}
	prev.Next = nil
	return head
}

// @lc code=end

