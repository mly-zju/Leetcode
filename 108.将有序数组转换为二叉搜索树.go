import "math"

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (66.83%)
 * Likes:    265
 * Dislikes: 0
 * Total Accepted:    36.4K
 * Total Submissions: 53.8K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
 *
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 *
 * 示例:
 *
 * 给定有序数组: [-10,-3,0,5,9],
 *
 * 一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	return addBinary(nums, 0, length-1)
}

func addBinary(nums []int, l, r int) *TreeNode {
	// 题目要求高度平衡，所以每次根节点都取中点
	if l > r {
		return nil
	} else if l == r {
		return &TreeNode{
			Val: nums[l],
		}
	}

	mid := (l+r) / 2
	return &TreeNode{
		Val: nums[mid],
		Left: addBinary(nums, l, mid - 1),
		Right: addBinary(nums, mid + 1, r),
	}
}

// @lc code=end

