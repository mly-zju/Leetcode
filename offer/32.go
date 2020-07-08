/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	cache := []*TreeNode{root}
	for levelLen := len(cache); levelLen > 0; {
		for i := 0; i < levelLen; i++ {
			res = append(res, cache[i].Val)
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
	return res
}