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
	flag := 0
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
		// 查看当前层是否需要反转
		if flag % 2 == 1 {
			l, r := 0, len(levelArr) - 1
			for l < r {
				levelArr[l], levelArr[r] = levelArr[r], levelArr[l]
				l++
				r--
			}
		}
		res = append(res, levelArr)
		flag++
	}
	return res
}