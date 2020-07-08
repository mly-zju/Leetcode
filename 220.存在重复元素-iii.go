import "sort"

/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 *
 * https://leetcode-cn.com/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Medium (24.81%)
 * Likes:    107
 * Dislikes: 0
 * Total Accepted:    9.7K
 * Total Submissions: 38.3K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * 给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，并且 i 和 j
 * 之间的差的绝对值最大为 ķ。
 *
 * 示例 1:
 *
 * 输入: nums = [1,2,3,1], k = 3, t = 0
 * 输出: true
 *
 * 示例 2:
 *
 * 输入: nums = [1,0,1,1], k = 1, t = 2
 * 输出: true
 *
 * 示例 3:
 *
 * 输入: nums = [1,5,9,1,5,9], k = 2, t = 3
 * 输出: false
 *
 */

// @lc code=start
type TreeNode struct {
	Val int
	Height int
	Left *TreeNode
	Right *TreeNode
}

func InsertAvl(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
			Height: 1,
		}
	}

	if val <= root.Val {
		root.Left = InsertAvl(root.Left, val)
		// 检测是否需要旋转
		leftH, rightH := GetHeight(root.Left), GetHeight(root.Left)
		if leftH - rightH > 1 {
			if val <= root.Left.Val {
				root = rotateLL(root)
			} else {
				root = rotateLR(root)
			}
		}
	} else {
		root.Right = InsertAvl(root.Right, val)
		// 检测旋转
		leftH, rightH := GetHeight(root.Left), GetHeight(root.Left)
		if rightH - leftH > 1 {
			if val > root.Right.Val {
				root = rotateRR(root)
			} else {
				root = rotateRL(root)
			}
		}
	}

	// 更新高度
	root.Height = getMax(GetHeight(root.Left), GetHeight(root.Right)) + 1
	return root
}

func DeleteAvl(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	root = delete(root, val)
	// 尝试执行平衡
	leftH, rightH := GetHeight(root.Left), GetHeight(root.Right)
	if leftH - rightH > 1 {
		// 进一步判断
		subLeftH, subRightH := GetHeight(root.Left.Left), GetHeight(root.Left.Right)
		if subLeftH > subRightH {
			root = rotateLL(root)
		} else {
			root = rotateLR(root)
		}
	} else if rightH - leftH > 1 {
		// 进一步判断
		subLeftH, subRightH := GetHeight(root.Right.Left), GetHeight(root.Right.Right)
		if subRightH > subLeftH {
			root = rotateRR(root)
		} else {
			root = rotateRL(root)
		}
	}

	return root
}

// 内部删除节点方法
func delete(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	if val < root.Val {
		root.Left = delete(root.Left, val)
		}
	} else if val > root.Val {
		root.Right = delete(root.Right, val)
	} else {
		// 删除本身，区分是否有子树
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// 寻找左子树最大值或者右子树最小值
			maxVal := findMax(root.Left)
			root.Val = maxVal
			root.Left = delete(root.Left, maxVal)
		}
	}

	// 更新高度
	root.Height = getMax(GetHeight(root.Left), GetHeight(root.Right)) + 1

	return root
}

func GetHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Height
}

func findMax(root *TreeNode) int {
	if root.Right == nil {
		return root.Val
	}

	return findMax(root.Right)
}

func findMin(root *TreeNode) int {
	if root.Left == nil {
		return root.Val
	}
	return findMin(root.Left)
}

func rotateLL(root *TreeNode) *TreeNode {
	// 左左
	oldLeft := root.Left
	root.Left = oldLeft.Right
	oldLeft.Right = root
	return oldLeft
}

func rotateRR(root *TreeNode) *TreeNode {
	oldRight := root.Right
	root.Right = oldRight.Left
	oldRight.Left = root
	return oldRight
}

func rotateLR(root *TreeNode) *TreeNode {
	root.Left = rotateRR(root.Left)
	return rotateLL(root)
}

func rotateRL(root *TreeNode) *TreeNode {
	root.Right = rotateLL(root.Right)
	return rotateRR(root)
}

// 寻找绝对值差最小的值
func findClosest(root *TreeNode, val int) int {
	if root == nil {
		return val
	}

	// 如果相等，直接返回0
	if root.Val == val {
		return 0
	}

	// 如果不相等，搜索左右子树的值，然后对比找出最小值
	leftVal, rightVal := findClosest(root.Left), findClosest(root.Right)
	midVal := getAbs(root.Val - val)
	return getMin(midVal, getMin(leftVal, rightVal))
}

func getAbs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	// 最直观的用大小为k的滑动窗口，然后两两比较窗口内的绝对值，但是这样会超时
	// 也可以继续基于滑动窗口，对窗口内的元素维护一个二叉搜索树，这样每次排序会比较简单
	// 或者基于桶排序来做，桶的容量就是k
	length := len(nums)
	if length < 2 {
		return false
	}

	// 1. 双循环滑动窗口法
	// for i := 0; i < length - 1; i++ {
	// 	j := i + 1
	// 	jMax := i + k
	// 	if jMax >= length {
	// 		jMax = length - 1
	// 	}

	// 	for ; j <= jMax; j++ {
	// 		if getAbs(nums[i] - nums[j]) <= t {
	// 			return true
	// 		}
	// 	}
	// }
	// return false

	// 2. avl树
	var root *TreeNode
	for i := 0; i < length - 1; i++ {
		root = InsertAvl(root, nums[i])
		if i == k {
			maxVal, minVal := findMax(root), findMin(root)
			if getAbs(maxVal - minVal) 
		}
	}
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

