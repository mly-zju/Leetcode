/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
 *
 * https://leetcode-cn.com/problems/path-sum-iii/description/
 *
 * algorithms
 * Easy (52.02%)
 * Likes:    217
 * Dislikes: 0
 * Total Accepted:    13.9K
 * Total Submissions: 26.4K
 * Testcase Example:  '[10,5,-3,3,2,null,11,3,-2,null,1]\n8'
 *
 * 给定一个二叉树，它的每个结点都存放着一个整数值。
 *
 * 找出路径和等于给定数值的路径总数。
 *
 * 路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
 *
 * 二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
 *
 * 示例：
 *
 * root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
 *
 * ⁠     10
 * ⁠    /  \
 * ⁠   5   -3
 * ⁠  / \    \
 * ⁠ 3   2   11
 * ⁠/ \   \
 * 3  -2   1
 *
 * 返回 3。和等于 8 的路径有:
 *
 * 1.  5 -> 3
 * 2.  5 -> 2 -> 1
 * 3.  -3 -> 11
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
var res int
func pathSum(root *TreeNode, sum int) int {
	// 类似这种找所有可能性，可以使用回溯法
	// 与普通回溯法不同的是，这个的结束条件是不一定的
	res = 0
	cache := []int{}
	dfs(root, sum, cache)
	return res
}

func dfs(root *TreeNode, sum int, cache []int) []int {
	if root == nil {
		return cache
	}

	cache = append(cache, root.Val)
	// 在每个节点作为结束节点，cache中存了独一份的回溯到root的路径，计算中奖是否有满足sum的
	length := len(cache)
	tmp := 0
	for i := length - 1; i >= 0; i-- {
		tmp += cache[i]
		if tmp == sum {
			res++
		}
	}
	dfs(root.Left, sum, cache)
	dfs(root.Right, sum, cache)
	// 可能要返回上层，出栈
	cache = cache[0:length - 1]
	return cache
}


// @lc code=end

