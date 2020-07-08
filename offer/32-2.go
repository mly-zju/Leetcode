/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	cache := []*TreeNode{root}
	for levelLen := len(cache); levelLen > 0; {
		levelArr := []int{}
		for i := 0; i < levelLen; i++ {
			levelArr = append(levelArr, cache[i].Val)
			if cache[i].Left != nil {
				cache = append(cache, cache[i].Left)
			}
			if cache[i].Right != nil {
				cache = append(cache, cache[i].Right)
			}
		}
		cache = cache[levelLen:]
		levelLen = len(cache)
		res = append(res, levelArr)
	}
	return res
}