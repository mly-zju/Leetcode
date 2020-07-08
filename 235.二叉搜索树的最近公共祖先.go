/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 *
 * https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (60.70%)
 * Likes:    184
 * Dislikes: 0
 * Total Accepted:    28.5K
 * Total Submissions: 46.7K
 * Testcase Example:  '[6,2,8,0,4,7,9,null,null,3,5]\n2\n8'
 *
 * 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 * 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
 *
 *
 *
 *
 *
 * 示例 1:
 *
 * 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
 * 输出: 6
 * 解释: 节点 2 和节点 8 的最近公共祖先是 6。
 *
 *
 * 示例 2:
 *
 * 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
 * 输出: 2
 * 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
 *
 *
 *
 * 说明:
 *
 *
 * 所有节点的值都是唯一的。
 * p、q 为不同节点且均存在于给定的二叉搜索树中。
 *
 *
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// var target *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 所有节点都是唯一，则只要判别一个在左边一个在右边即可
	maxVal, minVal := 0, 0
	if p.Val > q.Val {
		maxVal = p.Val
		minVal = q.Val
	} else {
		maxVal = q.Val
		minVal = p.Val
	}

	return dfs(root, minVal, maxVal)
}

func dfs(root *TreeNode, minVal, maxVal int) *TreeNode {
	if root == nil {
		return nil
	}

	// 如果某个节点等于自己，那么一定存在
	if root.Val == minVal || root.Val == maxVal {
		return root
	} else if root.Val > minVal && root.Val < maxVal {
		return root
	} else if root.Val > maxVal {
		return dfs(root.Left, minVal, maxVal)
	} else if root.Val < minVal {
		return dfs(root.Right, minVal, maxVal)
	}
	return nil
}



// func inOrder(root, p, q *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	leftNum := inOrder(root.Left, p, q)
// 	rightNum := inOrder(root.Right, p, q)

// 	var maxVal, minVal int
// 	if p.Val > q.Val {
// 		maxVal = p.Val
// 		minVal = q.Val
// 	} else {
// 		maxVal = q.Val
// 		minVal = p.Val
// 	}
	
// 	// 三种情况下当前节点一定是target：
// 	if leftNum == 1 && rightNum == 1 {
// 		// 如果left和right子树各包含一个，说明当前一定是
// 		target = root
// 		return 2
// 	} else if leftNum == 1 && root.Val == maxVal {
// 		// 自己是右节点且左子树包含，说明当前一定是
// 		target = root
// 		return 2
// 	} else if rightNum == 1 && root.Val == minVal {
// 		target = root
// 		return 2
// 	}

// 	// 非以上情况，计算当前包含数目
// 	if root.Val == maxVal || root.Val == minVal {
// 		return leftNum + rightNum + 1
// 	}
// 	return leftNum + rightNum
// }

// @lc code=end

