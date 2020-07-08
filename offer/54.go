/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
	// 反向中序迭代遍历: 先右后左
	cache := []*TreeNode{}
	count := 0
	var res int
	for root != nil || len(cache) > 0 {
		for root != nil {
			cache = append(cache, root)
			root = root.Right
		}
		newLen := len(cache)
		root = cache[newLen-1]
		cache = cache[0:newLen-1]
		count++
		if count == k {
			res = root.Val
			break
		} else {
			root = root.Left
		}
	}
	return res
}