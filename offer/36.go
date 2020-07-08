package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func changeTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// 迭代执行中序遍历，需要记录前驱节点
	var prev, head *TreeNode
	cache := []*TreeNode{}
	for root != nil || len(cache) != 0 {
		for root != nil {
			cache = append(cache, root)
			root = root.Left
		}
		newLen := len(cache)
		root = cache[newLen-1]
		cache = cache[0 : newLen-1]

		if prev == nil {
			// 记录head和prev，记得head最后还要修改左指针到最后一个节点
			head = root
			prev = root
		} else {
			root.Left = prev
			prev.Right = root
			prev = root
		}
		root = root.Right
	}
	// 最后的prev就是最后一个节点
	head.Left = prev
	prev.Right = head
	return head
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	root = changeTree(root)
	for i := 0; i < 5; i++ {
		fmt.Println("---------------")
		fmt.Println(root.Val, root.Left.Val, root.Right.Val)
		root = root.Right
	}
}
