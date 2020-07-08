/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (35.71%)
 * Likes:    3302
 * Dislikes: 0
 * Total Accepted:    244.2K
 * Total Submissions: 679.6K
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	var cur, head *ListNode
	for l1 != nil || l2 != nil {
		// 计算结果
		tmpVal := flag
		flag = 0
		if l1 != nil {
			tmpVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmpVal += l2.Val
			l2 = l2.Next
		}
		if tmpVal >= 10 {
			tmpVal = tmpVal % 10
			flag = 1
		}

		// 构造新节点
		if head == nil {
			head = &ListNode{
				Val: tmpVal,
			}
			cur = head
		} else {
			cur.Next = &ListNode{
				Val: tmpVal,
			}
			cur = cur.Next
		}
	}

	// 看看新的进位
	if flag == 1 {
		cur.Next = &ListNode{
			Val: 1,
		}
	}
	return head
}

// @lc code=end

