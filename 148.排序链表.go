/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 *
 * https://leetcode-cn.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (62.07%)
 * Likes:    324
 * Dislikes: 0
 * Total Accepted:    30.2K
 * Total Submissions: 48.2K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 *
 * 示例 1:
 *
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 *
 *
 * 示例 2:
 *
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
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
func sortList(head *ListNode) *ListNode {
	// nlogn, 快排不是双链表不现实，归并排序其实空间复杂度超了O(1), 只有堆排序才符合条件。
	// 不过这里就先用归并排序玩一下
	// 截止条件：为空或者只有一个元素的时候
	if head == nil || head.Next == nil {
		return head
	}
	// 寻找中点
	s, f := head, head
	var prevS *ListNode
	for f != nil {
		f = f.Next
		if f != nil {
			f = f.Next
			prevS = s
			s = s.Next
		}
	}
	// s认为是中点
	// 记得还要把中点前一个节点next置为nil，分成两个链表
	prevS.Next = nil

	return mergeList(sortList(head), sortList(s))
}


func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	// 区分有一个为空的情况
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	// 1. 新建一个链表，这个很简单
	// var nHead, nTmp *ListNode
	// for l1 != nil && l2 != nil {
	// 	var curNode *ListNode
	// 	if l1.Val <= l2.Val {
	// 		curNode = l1
	// 		l1 = l1.Next
	// 	} else {
	// 		curNode = l2
	// 		l2 = l2.Next
	// 	}

	// 	if nHead == nil {
	// 		nHead = curNode
	// 		nTmp = curNode
	// 	} else {
	// 		nTmp.Next = curNode
	// 		nTmp = curNode
	// 	}
	// }

	// // 分别继续执行
	// if l1 != nil {
	// 	nTmp.Next = l1
	// }
	// if l2 != nil {
	// 	nTmp.Next = l2
	// }
	// return nHead

	// 2. 原地算法, 以l1为基准
	var nHead, prev *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			if nHead == nil {
				nHead = l1
			}
			prev = l1
			l1 = l1.Next
		} else {
			if nHead == nil {
				nHead = l2
				prev = l2
				l2 = l2.Next
				prev.Next = l1
			} else {
				// 将l2插入prev和l1之间
				ol2 := l2
				l2 = l2.Next
				prev.Next = ol2
				ol2.Next = l1
				prev = ol2
			}
		}
	}

	if l1 == nil && l2 != nil {
		prev.Next = l2
	}
	return nHead
}

// @lc code=end

