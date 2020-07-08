package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertNode(t *TreeNode, num int) *TreeNode {
	if t == nil {
		return &TreeNode{
			Val: num,
		}
	}

	if num <= t.Val {
		t.Left = insertNode(t.Left, num)
	} else {
		t.Right = insertNode(t.Right, num)
	}
	return t
}

func findMin(t *TreeNode) int {
	if t.Left == nil {
		return t.Val
	}
	return findMin(t.Left)
}

func findMax(t *TreeNode) int {
	if t.Right == nil {
		return t.Val
	}
	return findMax(t.Right)
}

func findNode(t *TreeNode, num int) *TreeNode {
	if t == nil {
		return nil
	}
	if t.Val == num {
		return t
	} else if num < t.Val {
		return findNode(t.Left, num)
	} else {
		return findNode(t.Right, num)
	}
}

func deleteNode(t *TreeNode, num int) *TreeNode {
	if t == nil {
		return nil
	}

	if num < t.Val {
		t.Left = deleteNode(t.Left, num)
		return t
	} else if num > t.Val {
		t.Right = deleteNode(t.Right, num)
		return t
	}

	// 否则就是删除自身
	if t.Left == nil && t.Right == nil {
		// 如果自己是叶子节点，直接返回nil即可
		return nil
	} else if t.Right == nil {
		// 如果自己右子树为空，直接左子树顶上即可
		return t.Left
	} else {
		// 否则寻找右子树最小值, 替换自己，并进一步删除自己的右子树
		minVal := findMin(t.Right)
		t.Val = minVal
		t.Right = deleteNode(t.Right, minVal)
		return t
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	// 维护二叉搜索树
	var root *TreeNode
	// 先初始化
	length := len(nums)
	if length < k {
		return []int{}
	}

	res := []int{}
	for i := 0; i < length; i++ {
		if i < k {
			root = insertNode(root, nums[i])
			if i == k-1 {
				res = append(res, findMax(root))
			}
		} else {
			// 删除最前面的
			root = deleteNode(root, nums[i-k])
			root = insertNode(root, nums[i])
			res = append(res, findMax(root))
		}
	}
	return res
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
