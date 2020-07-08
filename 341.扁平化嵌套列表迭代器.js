/*
 * @lc app=leetcode.cn id=341 lang=javascript
 *
 * [341] 扁平化嵌套列表迭代器
 *
 * https://leetcode-cn.com/problems/flatten-nested-list-iterator/description/
 *
 * algorithms
 * Medium (60.76%)
 * Likes:    53
 * Dislikes: 0
 * Total Accepted:    3.8K
 * Total Submissions: 6.3K
 * Testcase Example:  '[[1,1],2,[1,1]]'
 *
 * 给定一个嵌套的整型列表。设计一个迭代器，使其能够遍历这个整型列表中的所有整数。
 * 
 * 列表中的项或者为一个整数，或者是另一个列表。
 * 
 * 示例 1:
 * 
 * 输入: [[1,1],2,[1,1]]
 * 输出: [1,1,2,1,1]
 * 解释: 通过重复调用 next 直到 hasNext 返回false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
 * 
 * 示例 2:
 * 
 * 输入: [1,[4,[6]]]
 * 输出: [1,4,6]
 * 解释: 通过重复调用 next 直到 hasNext 返回false，next 返回的元素的顺序应该是: [1,4,6]。
 * 
 * 
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * function NestedInteger() {
 *
 *     Return true if this NestedInteger holds a single integer, rather than a nested list.
 *     @return {boolean}
 *     this.isInteger = function() {
 *         ...
 *     };
 *
 *     Return the single integer that this NestedInteger holds, if it holds a single integer
 *     Return null if this NestedInteger holds a nested list
 *     @return {integer}
 *     this.getInteger = function() {
 *         ...
 *     };
 *
 *     Return the nested list that this NestedInteger holds, if it holds a nested list
 *     Return null if this NestedInteger holds a single integer
 *     @return {NestedInteger[]}
 *     this.getList = function() {
 *         ...
 *     };
 * };
 */
/**
 * @constructor
 * @param {NestedInteger[]} nestedList
 */
var NestedIterator = function(nestedList) {
    this.list = []; 
    // 使用递归先扁平化
    this.resetList(nestedList);
};

NestedIterator.prototype.resetList = function(nestedList) {
    // 1. 递归解法
    // let len = nestedList.length;
    // for (let i = 0; i < len; i++) {
    //     if (nestedList[i].isInteger()) {
    //         this.list.push(nestedList[i].getInteger());
    //     } else {
    //         this.resetList(nestedList[i].getList());
    //     }
    // }

    // 2. 迭代法：基于栈
    let stack = [nestedList];
    while (stack.length > 0) {
        let curNode = stack.pop();
        if (curNode instanceof Array) {
            // 如果curNode是数组，需要先取第一个值
            let tmpNode = curNode.shift();
            if (curNode.length > 0) {
                stack.push(curNode);
            }
            
            if (tmpNode == undefined) {
                continue;
            }

            if (tmpNode.isInteger()) {
                this.list.push(tmpNode.getInteger());
            } else {
                stack.push(tmpNode.getList());
            }

        } else if (curNode != undefined) {
            this.list.push(curNode.getInteger());
        }
    }
}


/**
 * @this NestedIterator
 * @returns {boolean}
 */
NestedIterator.prototype.hasNext = function() {
    let length = this.list.length;
    return length > 0;  
};

/**
 * @this NestedIterator
 * @returns {integer}
 */
NestedIterator.prototype.next = function() {
    return this.list.shift(); 
};

/**
 * Your NestedIterator will be called like this:
 * var i = new NestedIterator(nestedList), a = [];
 * while (i.hasNext()) a.push(i.next());
*/
// @lc code=end

