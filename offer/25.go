/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 1. 新建一个链表，这个很简单

	// 区分有一个为空的情况
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var nHead, nTmp *ListNode
	for l1 != nil && l2 != nil {
		curVal := 0
		if l1.Val < l2.Val {
			curVal = l1.Val
			l1 = l1.Next
		} else {
			curVal = l2.Val
			l2 = l2.Next
		}

		// 区分头部
		if nHead == nil {
			nHead = &ListNode{
				Val: curVal,
			}
			nTmp = nHead
		} else {
			nTmp.Next = &ListNode{
				Val: curVal,
			}
			nTmp = nTmp.Next
		}
	}

	// 分别继续执行
	for l1 != nil {
		nTmp.Next = &ListNode{
			Val: l1.Val,
		}
		nTmp = nTmp.Next
		l1 = l1.Next
	}
	for l2 != nil {
		nTmp.Next = &ListNode{
			Val: l2.Val,
		}
		nTmp = nTmp.Next
		l2 = l2.Next
	}
	return nHead

	// 2. 原地修改算法
	// if l1 == nil {
	// 	return l2
	// } else if l2 == nil {
	// 	return l1
	// }
	
	// // 以l1为基准
	// var prev, head *ListNode =  nil, l1
	// for l1 != nil && l2 != nil {
	// 	if l2.Val < l1.Val {
	// 		// l2小于l1的时候，插入到l1当前的前一位
	// 		newNode := l2
	// 		l2 = l2.Next
	// 		if prev == nil {
	// 			prev = newNode
	// 			prev.Next = l1
	// 			head = prev
	// 		} else {
	// 			prev.Next = newNode
	// 			prev = newNode
	// 			prev.Next = l1
	// 		}
	// 	} else {
	// 		// l2大于l1的时候，l1只需要继续前进
	// 		prev = l1
	// 		l1 = l1.Next
	// 	}
	// }
	// // 如果是l1到底了而l2没到底，则把剩余l2全部放在l1后面
	// if l1 == nil && l2 != nil {
	// 	prev.Next = l2
	// }
	// return head
}