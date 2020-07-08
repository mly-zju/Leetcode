/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (72.54%)
 * Likes:    77
 * Dislikes: 0
 * Total Accepted:    25.1K
 * Total Submissions: 34.3K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的前序遍历。
 * 
 * 例如，给定一个 3叉树 :
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 返回其前序遍历: [1,3,5,6,2,4]。
 * 
 * 
 * 
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	// 1. 递归
	// res := []int{}
	// if root == nil {
	// 	return res
	// }
	// res = dfs(root, res)
	// return res

	// 2. 迭代, 需要visitMap记录每个节点的子节点访问到第几个索引了
	res := []int{}
	if root == nil {
		return res
	}
	visitMap := make(map[*Node]int)
	cache := []*Node{}
	for root != nil || len(cache) != 0 {
		for root != nil {
			// root需要从上一次索引的下一个开始
			prevIndex := visitMap[root]
			// 只有第一次访问root时候才推入res
			if prevIndex == 0 {
				res = append(res, root.Val)
			}
			visitMap[root]++
			// 只有还有剩余子节点没遍历的时候才需要再次推入cache
			if prevIndex < len(root.Children) {
				cache = append(cache, root)
				root = root.Children[prevIndex]
			} else {
				break
			}
		}
		length := len(cache)
		if length == 0 {
			break
		}
		root = cache[length-1]
		cache = cache[0:length-1]
		if visitMap[root] != len(root.Children) {
			prevIndex := visitMap[root]
			visitMap[root]++
			cache = append(cache, root)
			root = root.Children[prevIndex]
		} else {
			root = nil
		}
	}
	return res
}

func dfs(root *Node, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	for _, node := range root.Children {
		res = dfs(node, res)
	}
	return res
}
// @lc code=end

