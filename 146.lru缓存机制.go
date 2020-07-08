/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (43.74%)
 * Likes:    450
 * Dislikes: 0
 * Total Accepted:    40.4K
 * Total Submissions: 86.9K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 * 
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 * 
 * 进阶:
 * 
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 * 
 * 示例:
 * 
 * LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
 * 
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得密钥 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得密钥 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 * 
 * 
 */

// @lc code=start
type ListNode struct {
    Val int
    Key int
    Next *ListNode
    Prev *ListNode
}

type LRUCache struct {
    NodeMap map[int]*ListNode
    Head *ListNode
    Tail *ListNode
    Cap int 
    Len int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        NodeMap: make(map[int]*ListNode),
        Cap: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    targetNode := this.NodeMap[key]
    if targetNode == nil {
        return -1
    }

    this.move2Head(targetNode)
    return targetNode.Val
}


func (this *LRUCache) Put(key int, value int)  {
    targetNode := this.NodeMap[key]
    if targetNode == nil {
        // 如果不存在，直接放到头部
        newNode := &ListNode{
            Val: value,
            Key: key,
        }
        this.addHead(newNode)

        // 更新nodeMap
        this.NodeMap[key] = newNode
        // 长度增加
        this.Len++

        if this.Len > this.Cap {
            // 如果超出cap，需要移除尾部并删除map
            this.removeTail()
        }
    } else {
        // 如果存在，修改记录并放到头部
        targetNode.Val = value
        this.move2Head(targetNode)
    }
}

func (this *LRUCache) removeTail() {
    delete(this.NodeMap, this.Tail.Key)

    prevNode := this.Tail.Prev
    prevNode.Next = nil
    this.Tail = prevNode
    this.Len--
}

func (this *LRUCache) move2Head(cur *ListNode) {
    prev, next := cur.Prev, cur.Next
    if prev == nil {
        // 如果prev为nil，说明cur本身就是头部，无需修改
        return
    }

    prev.Next = next
    if next != nil {
        next.Prev = prev
    } else {
        // 如果next为nil说明是尾部
        this.Tail = prev
    }

    // 更新head
    this.addHead(cur)
}

func (this *LRUCache) addHead(newNode *ListNode) {
    oldHead := this.Head
    this.Head = newNode
    this.Head.Next = oldHead
    this.Head.Prev = nil
    if oldHead == nil {
        // 如果老版head是nil，说明之前没有节点，需要更新尾部
        this.Tail = this.Head
    } else {
        // 如果不是空链表，更新原来的Prev
        oldHead.Prev = this.Head
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

