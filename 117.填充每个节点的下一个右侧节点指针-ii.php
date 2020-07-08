/*
 * @lc app=leetcode.cn id=117 lang=php
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 *
 * https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/description/
 *
 * algorithms
 * Medium (40.85%)
 * Likes:    79
 * Dislikes: 0
 * Total Accepted:    9.9K
 * Total Submissions: 23.4K
 * Testcase Example:  '{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":null,"right":null,"val":4},"next":null,"right":{"$id":"4","left":null,"next":null,"right":null,"val":5},"val":2},"next":null,"right":{"$id":"5","left":null,"next":null,"right":{"$id":"6","left":null,"next":null,"right":null,"val":7},"val":3},"val":1}'
 *
 * 给定一个二叉树
 * 
 * struct Node {
 * ⁠ int val;
 * ⁠ Node *left;
 * ⁠ Node *right;
 * ⁠ Node *next;
 * }
 * 
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
 * 
 * 初始状态下，所有 next 指针都被设置为 NULL。
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 
 * 
 * 输入：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":null,"right":null,"val":4},"next":null,"right":{"$id":"4","left":null,"next":null,"right":null,"val":5},"val":2},"next":null,"right":{"$id":"5","left":null,"next":null,"right":{"$id":"6","left":null,"next":null,"right":null,"val":7},"val":3},"val":1}
 * 
 * 
 * 输出：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":{"$id":"4","left":null,"next":{"$id":"5","left":null,"next":null,"right":null,"val":7},"right":null,"val":5},"right":null,"val":4},"next":{"$id":"6","left":null,"next":null,"right":{"$ref":"5"},"val":3},"right":{"$ref":"4"},"val":2},"next":null,"right":{"$ref":"6"},"val":1}
 * 
 * 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 你只能使用常量级额外空间。
 * 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
 * 
 * 
 */

// @lc code=start
/*
// Definition for a Node.
class Node {
    public $val;
    public $left;
    public $right;
    public $next;

    @param Integer $val 
    @param Node $left 
    @param Node $right 
    @param Node $next 
    function __construct($val, $left, $right, $next) {
        $this->val = $val;
        $this->left = $left;
        $this->right = $right;
        $this->next = $next;
    }
}
*/
class Solution {

    /**
     * @param Node $root
     * @return Node
     */
    function connect($root) {
        // 基于next遍历
        if ($root == NULL || ($root->left == NULL && $root->right == NULL)) {
            return $root;
        }

        $header = NULL;
        // 第一层只需要更新left
        if ($root->left != NULL) {
            if ($root->right != NULL) {
                $root->left->next = $root->right;
            }
            $header = $root->left;
        } else {
            $header = $root->right;
        }
        $cur = $header;
        $prev = NULL;

        while ($header != NULL) {
            // 查找当前层下一个子树不为null的节点
            while ($cur != NULL && $cur->left == NULL && $cur->right == NULL) {
                $cur = $cur->next;
            }
            if ($cur != NULL) {
                $flag = 0;
                if ($cur->left != NULL) {
                    // 如果左子树存在
                    if ($cur->right != NULL) {
                        $cur->left->next = $cur->right;
                    }
                } else {
                    $flag = 1;
                }

                // prev执行连接
                if ($prev != NULL) {
                    if ($prev->right != NULL) {
                        $prev->right->next = $flag == 0 ? $cur->left : $cur->right;
                    } else {
                        $prev->left->next = $flag == 0 ? $cur->left : $cur->right;
                    }
                }
                $prev = $cur;
                $cur = $cur->next;
            } else {
                // 如果没找到，说明本层结束，查找下一层的开头
                while ($header != NULL && $header->left == NULL && $header->right == NULL) {
                    $header = $header->next;
                }
                if ($header != NULL) {
                    $header = $header->left != NULL ? $header->left : $header->right;
                    $cur = $header;
                    $prev = NULL;
                }
            }
        }
        return $root;
    }
}
// @lc code=end

