/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	// 1. 递归法, 主体是左右子节点
	if root == nil {
		return true
	}
	// return dfs(root.Left, root.Right)

	// 2. 迭代法，层次遍历
	cache := []*TreeNode{root}
	for levelLen := len(cache); levelLen != 0; {
		// cache里面天然的是每一层的节点，先判断是否对称
		l, r := 0, levelLen - 1
		for l < r {
			if cache[l] == nil && cache[r] == nil {
				l++
				r--
				continue
			} else if cache[l] != nil && cache[r] != nil && cache[l].Val == cache[r].Val {
				l++
				r--
				continue
			}
			return false
		}

		// 插入子节点
		for i := 0; i < levelLen; i++ {
			if cache[i] != nil {
				cache = append(cache, cache[i].Left)
				cache = append(cache, cache[i].Right)
			}
		}
		cache = cache[levelLen:]
		levelLen = len(cache)
	}
	return true
}

func dfs(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}
	return dfs(t1.Left, t2.Right) && dfs(t1.Right, t2.Left)
}