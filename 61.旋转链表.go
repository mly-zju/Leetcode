/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 *
 * https://leetcode-cn.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (39.06%)
 * Likes:    162
 * Dislikes: 0
 * Total Accepted:    30.3K
 * Total Submissions: 77.3K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->4->5->NULL, k = 2
 * 输出: 4->5->1->2->3->NULL
 * 解释:
 * 向右旋转 1 步: 5->1->2->3->4->NULL
 * 向右旋转 2 步: 4->5->1->2->3->NULL
 *
 *
 * 示例 2:
 *
 * 输入: 0->1->2->NULL, k = 4
 * 输出: 2->0->1->NULL
 * 解释:
 * 向右旋转 1 步: 2->0->1->NULL
 * 向右旋转 2 步: 1->2->0->NULL
 * 向右旋转 3 步: 0->1->2->NULL
 * 向右旋转 4 步: 2->0->1->NULL
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	// 先计算length, 并寻找尾部节点
	length := 0
	tmp := head
	var tail *ListNode
	for tmp != nil {
		length++
		tail = tmp
		tmp = tmp.Next
	}

	// 去除无效旋转
	k = k % length
	if k == 0 {
		return head
	}

	cur := head
	moveNum := length - k
	for i := 0; i < moveNum - 1; i++ {
		cur = cur.Next
	}

	// 尾部节点next连向头部，断开节点next置为nil
	newHead := cur.Next
	tail.Next = head
	cur.Next = nil
	return newHead
}

// @lc code=end

