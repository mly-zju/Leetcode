/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (38.83%)
 * Likes:    295
 * Dislikes: 0
 * Total Accepted:    44.5K
 * Total Submissions: 114.2K
 * Testcase Example:  '[1,2]'
 *
 * 请判断一个链表是否为回文链表。
 *
 * 示例 1:
 *
 * 输入: 1->2
 * 输出: false
 *
 * 示例 2:
 *
 * 输入: 1->2->2->1
 * 输出: true
 *
 *
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 当然可以转为数组计算
	// 如果空间复杂度需要O(1), 则需要找中点，反转链表之后进行比较
	// 翻转链表也有两种处理方案：1. 从中点开始反转左边  2. 从中点开始反转右边
	// 显然还是1比较方便，因为反转左边之后，天然的尾部都是nil，否则右边还要判断尾部
	s, f := head, head
	length := 0
	for f != nil {
		length++
		f = f.Next
		if f == nil {
			break
		} else {
			length++
			f = f.Next
		}

		if f != nil {
			s = s.Next
		}
	}

	// 右半边的头部根据length奇偶会有不同
	var rHead *ListNode
	if length % 2 == 0 {
		rHead = s.Next
	} else {
		rHead = s
	}

	// 反转右边链表
	var prev, cur, next *ListNode = nil, rHead, rHead.Next
	for next != nil {
		cur.Next = prev
		prev = cur
		cur = next
		next = next.Next
	}
	cur.Next = prev

	// 比较cur和head, 二者一样长，只判断一个即可
	for cur != rHead {
		if cur.Val != head.Val {
			return false
		}
		cur = cur.Next
		head = head.Next
	}

	// 肯定是一样长的，所以只要比较到了尽头，再比较一次即可
	return cur.Val == head.Val
}

// @lc code=end

