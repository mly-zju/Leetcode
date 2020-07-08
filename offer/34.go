/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	// 回溯法
	res := [][]int{}
	if root == nil {
		return res
	}
	
	tmpCache := []int{}
	return dfs(root, sum, tmpCache, res)
}

func dfs(root *TreeNode, sum int, tmpCache []int, res [][]int) [][]int {
	// 叶子节点作为结束
	if root.Left == nil && root.Right == nil {
		tmpCache = append(tmpCache, root.Val)
		if getSum(tmpCache) == sum {
			newArr := make([]int, len(tmpCache))
			copy(newArr, tmpCache)
			res = append(res, newArr)
		}
		return res
	} else {
		oldLen := len(tmpCache)
		tmpCache = append(tmpCache, root.Val)
		if root.Left != nil {
			res = dfs(root.Left, sum, tmpCache, res)
		}
		if root.Right != nil {
			res = dfs(root.Right, sum, tmpCache, res)
		}
		// 推出，防止对上层造成影响
		tmpCache = tmpCache[0:oldLen]
		return res
	}
}

func getSum(arr []int) int {
	res := 0
	for _, val := range arr {
		res += val
	}
	return res
}