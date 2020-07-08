/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 *
 * https://leetcode-cn.com/problems/intersection-of-two-linked-lists/description/
 *
 * algorithms
 * Easy (48.35%)
 * Likes:    437
 * Dislikes: 0
 * Total Accepted:    52K
 * Total Submissions: 105K
 * Testcase Example:  '8\n[4,1,8,4,5]\n[5,0,1,8,4,5]\n2\n3'
 *
 * 编写一个程序，找到两个单链表相交的起始节点。
 *
 * 如下面的两个链表：
 *
 *
 *
 * 在节点 c1 开始相交。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2,
 * skipB = 3
 * 输出：Reference of the node with value = 8
 * 输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为
 * [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
 *
 *
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB
 * = 1
 * 输出：Reference of the node with value = 2
 * 输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为
 * [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
 *
 *
 *
 *
 * 示例 3：
 *
 *
 *
 * 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
 * 输出：null
 * 输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为
 * 0，而 skipA 和 skipB 可以是任意值。
 * 解释：这两个链表不相交，因此返回 null。
 *
 *
 *
 *
 * 注意：
 *
 *
 * 如果两个链表没有交点，返回 null.
 * 在返回结果后，两个链表仍须保持原有的结构。
 * 可假定整个链表结构中没有循环。
 * 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 1. 哈希法
	// nodeMap := make(map[*ListNode]bool)
	// tmp := headA
	// // 先遍历headA
	// for tmp != nil {
	// 	nodeMap[tmp] = true
	// 	tmp = tmp.Next
	// }
	// // 再遍历headB
	// tmp = headB
	// for tmp != nil {
	// 	if nodeMap[tmp] {
	// 		return tmp
	// 	}
	// 	tmp = tmp.Next
	// }
	// return nil

	// 2. 双指针法
	lena, lenb := 0, 0
	tmp := headA
	for tmp != nil {
		lena++
		tmp = tmp.Next
	}
	tmp = headB
	for tmp != nil {
		lenb++
		tmp = tmp.Next
	}

	curA, curB := headA, headB
	if lena > lenb {
		for i := 0; i < lena - lenb; i++ {
			curA = curA.Next
		}
	} else if lenb > lena {
		for i := 0; i < lenb - lena; i++ {
			curB = curB.Next
		}
	}

	// 开始寻找
	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}

// @lc code=end

