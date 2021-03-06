/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
 *
 * https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/description/
 *
 * algorithms
 * Easy (46.73%)
 * Likes:    37
 * Dislikes: 0
 * Total Accepted:    4.5K
 * Total Submissions: 7.8K
 * Testcase Example:  '[1,0,1,0,1,0,1]'
 *
 * 给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。例如，如果路径为 0 -> 1 -> 1
 * -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。
 * 
 * 对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。
 * 
 * 以 10^9 + 7 为模，返回这些数字之和。
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 
 * 输入：[1,0,1,0,1,0,1]
 * 输出：22
 * 解释：(100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中的结点数介于 1 和 1000 之间。
 * node.val 为 0 或 1 。
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
func sumRootToLeaf(root *TreeNode) int {
	routeArr := dfs(root)
	fmt.Println(routeArr)
	sum := 0
	for _, val := range routeArr {
		sum += getDeci(val)
	}
	return sum
}

func dfs(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res 
	}
	newStr := strconv.Itoa(root.Val)
	leftArr := dfs(root.Left)
	for _, str := range leftArr {
		res = append(res, newStr + str)
	}
	rightArr := dfs(root.Right)
	for _, str := range rightArr {
		res = append(res, newStr + str)
	}

	if len(res) == 0 {
		res = append(res, newStr)
	}

	return res
}

func getDeci(str string) int {
	res := 0
	for _, ch := range str {
		if ch == '0' {
			res = res * 2 + 0
		} else {
			res = res * 2 + 1
		}
	}
	return res
}
// @lc code=end

