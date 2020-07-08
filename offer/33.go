func verifyPostorder(postorder []int) bool {
	// 后序二叉搜索树中心点是最后一个值, 然后前面的数组里面要找到一个点，能够分出左右子树；然后递归左右子树查看。
	length := len(postorder)
	// 空节点true
	if length < 2 {
		return true
	}
	lastValue := postorder[length - 1]
	// 从前往后寻找第一个大于last的
	// 如果已经找到了最大的，但是后面却遇到了最小的，那就返回false
	firstBig := -1
	for i := 0; i < length - 1; i++ {
		if postorder[i] > lastValue && firstBig == -1 {
			firstBig = i
		} else if postorder[i] < lastValue && firstBig != -1 {
			return false
		}
	}

	if firstBig == -1 || firstBig == 0 {
		// -1，说明全是左子树, 0, 说明全是右子树
		return verifyPostorder(postorder[0:length - 1])
	}

	// 区分了左右子树之后，看左右子树是否分别是合法的搜索二叉树
	return verifyPostorder(postorder[0:firstBig]) && verifyPostorder(postorder[firstBig:length - 1])
}