/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

struct TreeNode *res;
struct TreeNode* lowestCommonAncestor(struct TreeNode* root, struct TreeNode* p, struct TreeNode* q) {
	res = NULL;
	dfs(root, p, q);
	return res;
}

int dfs(struct TreeNode* root, struct TreeNode* p, struct TreeNode* q) {
	// 如果是null，没有相等的值
	if (root == NULL) {
		return 0;
	}
	int leftEqual = 0;
	int rightEqual = 0;
	int selfEqual = 0;
	if (root->val == p->val || root->val == q->val) {
		selfEqual = 1;
	}

	leftEqual = dfs(root->left, p, q);
	rightEqual = dfs(root->right, p, q);

	// res不为null的情况下才执行判断
	if (res == NULL) {
		if ((selfEqual == 1 && (leftEqual == 1 || rightEqual == 1)) || (leftEqual == 1 && rightEqual == 1)) {
			res = root;
		}
	}
	return selfEqual + leftEqual + rightEqual;
}