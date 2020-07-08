/*
 * @lc app=leetcode.cn id=116 lang=php
 *
 * [116] 填充每个节点的下一个右侧节点指针
 *
 * https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/description/
 *
 * algorithms
 * Medium (46.75%)
 * Likes:    99
 * Dislikes: 0
 * Total Accepted:    16.1K
 * Total Submissions: 32.9K
 * Testcase Example:  '{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":null,"right":null,"val":4},"next":null,"right":{"$id":"4","left":null,"next":null,"right":null,"val":5},"val":2},"next":null,"right":{"$id":"5","left":{"$id":"6","left":null,"next":null,"right":null,"val":6},"next":null,"right":{"$id":"7","left":null,"next":null,"right":null,"val":7},"val":3},"val":1}'
 *
 * 给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
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
 * 输入：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":null,"right":null,"val":4},"next":null,"right":{"$id":"4","left":null,"next":null,"right":null,"val":5},"val":2},"next":null,"right":{"$id":"5","left":{"$id":"6","left":null,"next":null,"right":null,"val":6},"next":null,"right":{"$id":"7","left":null,"next":null,"right":null,"val":7},"val":3},"val":1}
 * 
 * 
 * 输出：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":{"$id":"4","left":null,"next":{"$id":"5","left":null,"next":{"$id":"6","left":null,"next":null,"right":null,"val":7},"right":null,"val":6},"right":null,"val":5},"right":null,"val":4},"next":{"$id":"7","left":{"$ref":"5"},"next":null,"right":{"$ref":"6"},"val":3},"right":{"$ref":"4"},"val":2},"next":null,"right":{"$ref":"7"},"val":1}
 * 
 * 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
 * 
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
        // 1. 层次遍历
        // $cache = [];
        // if ($root == NULL) {
        //     return $root;
        // }
        // $cache[] = $root;
        // while (count($cache) > 0) {
        //     $levelLen = count($cache);
        //     $levelArr = [];
        //     for ($i = 0; $i < $levelLen; $i++) {
        //         if ($i > 0) {
        //             $cache[$i - 1]->next = $cache[$i];
        //         }
        //         if ($cache[$i]->left != NULL) {
        //             $cache[] = $cache[$i]->left;
        //         }
        //         if ($cache[$i]->right != NULL) {
        //             $cache[] = $cache[$i]->right;
        //         }
        //     }
        //     // 截取cache
        //     $cache = array_slice($cache, $levelLen);
        // }
        // return $root;

        // 2. 基于next遍历
        if ($root == NULL || $root->left == NULL) {
            return $root;
        }

        $root->left->next = $root->right;

        $header = $root->left;
        $cur = $header;
        $prev = NULL;
        while ($header != NULL) {
            // 查找当前层下一个子树不为null的节点
            while ($cur != NULL && $cur->left == NULL) {
                $cur = $cur->next;
            }
            if ($cur != NULL) {
                $cur->left->next = $cur->right;
                if ($prev != NULL) {
                    $prev->right->next = $cur->left;
                }
                $prev = $cur;
                $cur = $cur->next;
            } else {
                // 如果没找到，说明本层结束，查找下一层的开头
                while ($header != NULL & $header->left == NULL) {
                    $header = $header->next;
                }
                if ($header != NULL) {
                    $header = $header->left;
                    $cur = $header;
                    $prev = NULL;
                }
            }
        }
        return $root;
    }
}
// @lc code=end

