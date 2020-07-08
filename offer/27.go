/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	// 1. 递归求解
	// return dfs(root)

	// 2. 层次遍历迭代求解
	if root == nil {
		return root
	}
	cache := []*TreeNode{root}
	for levelLen := len(cache); levelLen != 0; {
		for i := 0; i < levelLen; i++ {
			cache[i].Left, cache[i].Right = cache[i].Right, cache[i].Left
			if cache[i].Left != nil {
				cache = append(cache, cache[i].Left)
			}
			if cache[i].Right != nil {
				cache = append(cache, cache[i].Right)
			}
		}
		cache = cache[levelLen:]
		levelLen = len(cache)
	}
	return root
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 交换左右子树
	root.Left, root.Right = root.Right, root.Left
	root.Left = dfs(root.Left)
	root.Right = dfs(root.Right)
	return root
}