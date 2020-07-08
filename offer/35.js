/**
 * // Definition for a Node.
 * function Node(val,next,random) {
 *    this.val = val;
 *    this.next = next;
 *    this.random = random;
 * };
 */
/**
 * @param {Node} head
 * @return {Node}
 */
// function Node(val, next, random) {
//     this.val = val;
//     this.next = next;
//     this.random = random;
// }

var copyRandomList = function(head) {
    // 链表的深拷贝，关键在于random这个值，不能重复copy
    // 两个数组，一个数组存放每个节点的拷贝，next和random都设置null，另一个数组存放每个数组的random在链表中的位置索引，如果是null就存放-1。之后再把next和random连接起来。
    let nodeArr = [];
    let randomIndexArr = [];
    let cur = head;
    while (cur != null) {
        // 先推入新的节点
        nodeArr.push(new Node(cur.val, null, null));
        // 从头遍历，寻找random的位置, 这里一定要严格相等哦
        let tmp = head;
        let index = 0;
        let flag = false;
        while (tmp != null) {
            if (tmp == cur.random) {
                randomIndexArr.push(index);
                flag = true;
                break;
            }
            tmp = tmp.next;
            index++;
        }
        if (flag == false) {
            randomIndexArr.push(-1);
        }
            
        cur = cur.next;
    }

    // 第二次遍历，执行连接
    let len = nodeArr.length;
    for (let i = 0; i < len; i++) {
        // 连接next
        if (i != len - 1) {
            nodeArr[i].next = nodeArr[i+1];
        }

        // 连接random
        if (randomIndexArr[i] != -1) {
            nodeArr[i].random = nodeArr[randomIndexArr[i]];
        }
    }
    return nodeArr[0];
};