/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	length := len(preorder)
	if length == 0 {
		return nil
	}

	var mid int
	// 寻找根节点的重点
	for index, val := range inorder {
		if val == preorder[0] {
			mid = index
			break
		}
	}
	return &TreeNode{
		Val: inorder[mid],
		Left: buildTree(preorder[1:mid+1], inorder[0:mid]),
		Right: buildTree(preorder[mid+1:], inorder[mid+1:]),
	}
}