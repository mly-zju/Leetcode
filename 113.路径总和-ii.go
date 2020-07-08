/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 *
 * https://leetcode-cn.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (57.09%)
 * Likes:    138
 * Dislikes: 0
 * Total Accepted:    20K
 * Total Submissions: 34.8K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \    / \
 * ⁠       7    2  5   1
 *
 *
 * 返回:
 *
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
 * ]
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
var res [][]int
func pathSum(root *TreeNode, sum int) [][]int {
	// dfs + 回溯法即可
	res = [][]int{}
	if root == nil {
		return res
	}

	dfs(root, sum, []int{})
	return res
}

func dfs(root *TreeNode, sum int, cache []int) {
	// 为nil的直接返回
	if root == nil {
		return
	}

	// 到达叶子节点就是结束条件
	if root.Left == nil && root.Right == nil {
		cache = append(cache, root.Val)
		// 结束条件下，测试如果sum正确，推入最终结果
		if isEqual(cache, sum) {
			newEle := make([]int, len(cache))
			copy(newEle, cache)
			res = append(res, newEle)
		}
		cache = cache[0:len(cache) - 1]
	} else {
		cache = append(cache, root.Val)
		dfs(root.Left, sum, cache)
		dfs(root.Right, sum, cache)
		// 回到父元素前推出cache
		cache = cache[0:len(cache) - 1]
	}
}

func isEqual(arr []int, target int) bool {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum == target
}

// @lc code=end

