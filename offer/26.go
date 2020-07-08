/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		// 空树不是任何的子树
		return false
	}

	return isSame(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSame(A, B *TreeNode) bool {
	// 某个节点，B为nil而A不为nil，算作ok
	if B == nil {
		return true
	} else if A == nil {
		return false
	} else if A.Val != B.Val {
		return false
	}

	// 根节点相同，则继续判别子节点
	return isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
}