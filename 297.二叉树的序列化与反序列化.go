/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 *
 * https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (41.38%)
 * Likes:    154
 * Dislikes: 0
 * Total Accepted:    17.4K
 * Total Submissions: 38.2K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 * 
 * 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
 * 
 * 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 /
 * 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
 * 
 * 示例: 
 * 
 * 你可以将以下二叉树：
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠    / \
 * ⁠   4   5
 * 
 * 序列化为 "[1,2,3,null,null,4,5]"
 * 
 * 提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode
 * 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
 * 
 * 说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 两种方案：1. 两个数组，中序+前序 2. 一个数组，bfs，但是需要将null节点也记录下来
// 不过1存在一个问题，就是如果存在相同的val，这个方法会失效
// 因此还是优先考虑方法2
type Codec struct {
	list []string
}

func Constructor() Codec {
	return Codec{list: []string{}}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// bfs，将叶子null也需要加进去, 序列化为一维字符串，通过逗号分割
	res := ""
	cache := []*TreeNode{root}
	for len(cache) > 0 {
		levelLen := len(cache)
		for i := 0; i < levelLen; i++ {
			if cache[i] == nil {
				res = res + "null,"
			} else {
				res = res + strconv.Itoa(cache[i].Val) + ","
				// 对于非null节点，无条件推入子树
				cache = append(cache, cache[i].Left)
				cache = append(cache, cache[i].Right)
			}
		}
		cache = cache[levelLen:]
	}
	// 去掉最后的逗号
	res = res[0:len(res)-1]
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
	// 按照逗号分割，获取字符串数组
	nodeArr := strings.Split(data, ",")
	var root *TreeNode = &TreeNode{}
	cache := []*TreeNode{root}

	for i, length := 0, len(nodeArr); i < length; {
		if i == 0 {
			// i=0需要特殊处理
			if nodeArr[i] == "null" {
				return nil
			} else {
				num, _ := strconv.Atoi(nodeArr[i])
				cache[0].Val = num
				i++
			}
		} else {
			levelLen := len(cache)
			for j := 0; j < levelLen; j++ {
				// 先处理左子树
				if nodeArr[i] == "null" {
					cache[j].Left = nil
				} else {
					num, _ := strconv.Atoi(nodeArr[i])
					cache[j].Left = &TreeNode{
						Val: num,
					}
					cache = append(cache, cache[j].Left)
				}
				i++

				// 再处理右子树
				if nodeArr[i] == "null" {
					cache[j].Right = nil
				} else {
					num, _ := strconv.Atoi(nodeArr[i])
					cache[j].Right = &TreeNode{
						Val: num,
					}
					cache = append(cache, cache[j].Right)
				}
				i++
			}

			cache = cache[levelLen:]
		}
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

