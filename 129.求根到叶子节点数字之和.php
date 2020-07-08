/*
 * @lc app=leetcode.cn id=129 lang=php
 *
 * [129] 求根到叶子节点数字之和
 *
 * https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/description/
 *
 * algorithms
 * Medium (59.08%)
 * Likes:    91
 * Dislikes: 0
 * Total Accepted:    12.5K
 * Total Submissions: 20.9K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
 * 
 * 例如，从根到叶子节点路径 1->2->3 代表数字 123。
 * 
 * 计算从根到叶子节点生成的所有数字之和。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例 1:
 * 
 * 输入: [1,2,3]
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * 输出: 25
 * 解释:
 * 从根到叶子节点路径 1->2 代表数字 12.
 * 从根到叶子节点路径 1->3 代表数字 13.
 * 因此，数字总和 = 12 + 13 = 25.
 * 
 * 示例 2:
 * 
 * 输入: [4,9,0,5,1]
 * ⁠   4
 * ⁠  / \
 * ⁠ 9   0
 * / \
 * 5   1
 * 输出: 1026
 * 解释:
 * 从根到叶子节点路径 4->9->5 代表数字 495.
 * 从根到叶子节点路径 4->9->1 代表数字 491.
 * 从根到叶子节点路径 4->0 代表数字 40.
 * 因此，数字总和 = 495 + 491 + 40 = 1026.
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {

    public $numArr;

    /**
     * @param TreeNode $root
     * @return Integer
     */
    function sumNumbers($root) {
        $sum = 0;
        if ($root == NULL) {
            return $sum;
        }
        $cache = [];
        // 使用回溯法更简单一些, 从顶向下
        $this->dfs($root, $cache);
        foreach ($this->numArr as $v) {
            $sum += $v;
        }
        return $sum;
    }

    /**
     * @param TreeNode $root
     * @param Array $cache
     */
    function dfs($root, $cache) {
        // 如果叶子节点，则执行返回
        if ($root == NULL) {
            return;
        } else if ($root->left == NULL && $root->right == NULL) {
            $tmpNum = 0;
            foreach ($cache as $v) {
                $tmpNum = $tmpNum * 10 + $v;
            }
            $tmpNum = $tmpNum * 10 + $root->val;
            $this->numArr[] = $tmpNum;
            return;
        }

        $cache[] = $root->val;
        $this->dfs($root->left, $cache);
        $this->dfs($root->right, $cache);
        return;
    }
}
// @lc code=end

