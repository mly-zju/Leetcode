import "strconv"

/*
 * @lc app=leetcode.cn id=606 lang=golang
 *
 * [606] 根据二叉树创建字符串
 *
 * https://leetcode-cn.com/problems/construct-string-from-binary-tree/description/
 *
 * algorithms
 * Easy (51.41%)
 * Likes:    83
 * Dislikes: 0
 * Total Accepted:    6.4K
 * Total Submissions: 12.5K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。
 *
 * 空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。
 *
 * 示例 1:
 *
 *
 * 输入: 二叉树: [1,2,3,4]
 * ⁠      1
 * ⁠    /   \
 * ⁠   2     3
 * ⁠  /
 * ⁠ 4
 *
 * 输出: "1(2(4))(3)"
 *
 * 解释: 原本将是“1(2(4)())(3())”，
 * 在你省略所有不必要的空括号对之后，
 * 它将是“1(2(4))(3)”。
 *
 *
 * 示例 2:
 *
 *
 * 输入: 二叉树: [1,2,3,null,4]
 * ⁠      1
 * ⁠    /   \
 * ⁠   2     3
 * ⁠    \
 * ⁠     4
 *
 * 输出: "1(2()(4))(3)"
 *
 * 解释: 和第一个示例相似，
 * 除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。
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
func tree2str(t *TreeNode) string {
	// 1. 递归
	// return preOrder(t)
}

// func preOrder(t *TreeNode) string {
// 	res := ""
// 	if t == nil {
// 		return res
// 	}
// 	res += strconv.Itoa(t.Val)
// 	// 左右都空的时候直接省略。否则左边始终不能省略，右边空就省略。
// 	leftRes := preOrder(t.Left)
// 	rightRes := preOrder(t.Right)
// 	if leftRes == "" && rightRes == "" {
// 		// 左右都空时候直接省略
// 		return res
// 	}
// 	// 否则左边用于不能省略
// 	res += "(" + leftRes + ")"
// 	// 右边为空就省略
// 	if rightRes != "" {
// 		res += "(" + rightRes + ")"
// 	}

// 	return res
// }

// @lc code=end

