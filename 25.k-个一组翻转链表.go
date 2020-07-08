/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (57.18%)
 * Likes:    438
 * Dislikes: 0
 * Total Accepted:    48.3K
 * Total Submissions: 83.9K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 * 
 * k 是一个正整数，它的值小于或等于链表的长度。
 * 
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 * 
 * 
 * 
 * 示例：
 * 
 * 给你这个链表：1->2->3->4->5
 * 
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 * 
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 * 
 * 
 * 
 * 说明：
 * 
 * 
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	// 先计算count
	tmp := head
	for tmp != nil {
		count++
		tmp = tmp.Next
	}

	// 迭代执行
	var nHead *ListNode
	var prev, cur *ListNode = nil, head
	for count >= k {
		nCur := reverseK(cur, k)
		if prev == nil {
			nHead = nCur
			cur = nCur
		} else {
			prev.Next = nCur
			cur = nCur
		}

		// 向前进k步
		for i := 0; i < k; i++ {
			prev = cur
			cur = cur.Next
		}
		count -= k
	}
	return nHead
}

// 反转前k个元素
func reverseK(head *ListNode, k int) *ListNode {
	var prev, cur, next *ListNode = nil, head, head.Next
	// reverseFirst代表第一个反转的节点
	var reverseFirst *ListNode

	i := 0
	for i < k && next != nil {
		cur.Next = prev
		prev = cur
		cur = next
		next = next.Next
		if reverseFirst == nil {
			reverseFirst = prev
		}
		i++
	}

	// 如果还没到k，需要补一下cur(因为长度可能小于k)
	if i != k {
		cur.Next = prev
		prev = cur
		cur = next
		if reverseFirst == nil {
			reverseFirst = prev
		}
	}

	// reverseFirst的next需要设置为当前的cur，且新的head是当前的prev
	reverseFirst.Next = cur
	return prev
}
// @lc code=end

