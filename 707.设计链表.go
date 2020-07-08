/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 *
 * https://leetcode-cn.com/problems/design-linked-list/description/
 *
 * algorithms
 * Easy (22.96%)
 * Likes:    92
 * Dislikes: 0
 * Total Accepted:    10.3K
 * Total Submissions: 44.7K
 * Testcase Example:  '["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]\n' +
  '[[],[1],[3],[1,2],[1],[1],[1]]'
 *
 * 设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next
 * 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。
 * 
 * 在链表类中实现这些功能：
 * 
 * 
 * get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
 * addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
 * addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
 * addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index
 * 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
 * deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
 * 
 * 
 * 
 * 
 * 示例：
 * 
 * MyLinkedList linkedList = new MyLinkedList();
 * linkedList.addAtHead(1);
 * linkedList.addAtTail(3);
 * linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
 * linkedList.get(1);            //返回2
 * linkedList.deleteAtIndex(1);  //现在链表是1-> 3
 * linkedList.get(1);            //返回3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 所有val值都在 [1, 1000] 之内。
 * 操作次数将在  [1, 1000] 之内。
 * 请不要使用内置的 LinkedList 库。
 * 
 * 
 */
package main
import "fmt"
// @lc code=start
type MyLinkedList struct {
	Head *Node
}

type Node struct {
	Val int
	Next *Node
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	tmpHead := this.Head	
	for i := 0; i < index; i++ {
		if tmpHead == nil {
			return -1
		}
		tmpHead = tmpHead.Next
	}
	if tmpHead == nil {
		return -1
	}
	return tmpHead.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	tmpHead := this.Head
	newNode := &Node{
		Val: val,
		Next: tmpHead,
	}
	this.Head = newNode
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	tmpHead := this.Head
	newNode := &Node{
		Val: val,
	}
	if tmpHead == nil {
		this.Head = newNode
	} else {
		for tmpHead.Next != nil {
			tmpHead = tmpHead.Next
		}
		tmpHead.Next = newNode
	}
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	tmpHead := this.Head
	newNode := &Node{
		Val: val,
	}
	// 区分头部
	if index == 0 {
		newNode.Next = tmpHead
		this.Head = newNode
	} else {
		// 寻找index位置的前置元素, 因此移动index-1次
		for i := 0; i < index - 1; i++ {
			// 前置元素还没到就发现了nil，则不合法
			if tmpHead == nil {
				return
			}
			tmpHead = tmpHead.Next
		}
		if tmpHead == nil {
			return
		}

		// 开始执行插入
		oldNext := tmpHead.Next
		tmpHead.Next = newNode
		newNode.Next = oldNext
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	tmpHead := this.Head
	// 区分头部
	if index == 0 {
		this.Head = tmpHead.Next
	} else {
		// 寻找index位置前置元素
		for i := 0; i < index - 1; i++ {
			// 因为是删除，发现要删除的前置元素后面没有元素了，那么一定不合法
			if tmpHead.Next == nil {
				return
			}
			tmpHead = tmpHead.Next
		}
		if tmpHead.Next == nil {
			return
		}
		tmpHead.Next = tmpHead.Next.Next
	}
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

// func main() {
// 	obj := Constructor()
// 	obj.AddAtHead(2) // 2
// 	obj.DeleteAtIndex(1) // 2 不变
// 	obj.AddAtHead(2) // 2->2
// 	obj.AddAtHead(7) // 7->2->2
// 	obj.AddAtHead(3) // 3->7->2->2
// 	obj.AddAtHead(2) // 2->3->7->2->2
// 	obj.AddAtHead(5) // 5->2->3->7->2->2
// 	obj.AddAtTail(5) // 5->2->3->7->2->2->5
// 	fmt.Println(obj.Get(5)) // 2
// 	obj.DeleteAtIndex(6) // 不变
// 	obj.DeleteAtIndex(4) // 5->3->7->2->5
// }
// @lc code=end

