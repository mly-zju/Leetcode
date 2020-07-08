/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

/**
 * Encodes a tree to a single string.
 *
 * @param {TreeNode} root
 * @return {string}
 */
var serialize = function(root) {
    // 层次遍历的序列化
    let res = '';
    if (root == null) {
        return res;
    }
    let cache = [root];
    while (cache.length != 0) {
        let levelLen = cache.length;
        for (let i = 0; i < levelLen; i++) {
            if (cache[i] == null) {
                res += ' ' + 'null';
            } else {
                res += ' ' + cache[i].val;
            }

            if (cache[i] != null) {
                cache.push(cache[i].left);
                cache.push(cache[i].right);
            }
        }
        cache = cache.slice(levelLen);
    }
    return res;
};

/**
 * Decodes your encoded data to tree.
 *
 * @param {string} data
 * @return {TreeNode}
 */
var deserialize = function(data) {
    if (data == '') {
        return null;
    }

    let nodeArr = data.split(' '); 
    let root = new TreeNode(+nodeArr[0]);
    let cache = [root];    

    let i = 1;
    while (i < nodeArr.length) {
        let levelLen = cache.length;
        for (let j = 0; j < levelLen; j++) {
            if (nodeArr[i] == 'null') {
                // 如果节点是null，当前左子树为null且不推入cache
                cache[j].left = null;
            } else {
                cache[j].left = new TreeNode(+nodeArr[i]);
                cache.push(cache[j].left);
            }
            i++;
            if (nodeArr[i] == 'null') {
                cache[j].right = null;
            } else {
                cache[j].right = new TreeNode(+nodeArr[i]);
                cache.push(cache[j].right);
            }
            i++;
        }
        cache = cache.slice(levelLen);
    }
    return root;
};

/**
 * Your functions will be called as such:
 * deserialize(serialize(root));
 */