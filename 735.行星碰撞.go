/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] 行星碰撞
 *
 * https://leetcode-cn.com/problems/asteroid-collision/description/
 *
 * algorithms
 * Medium (34.59%)
 * Likes:    41
 * Dislikes: 0
 * Total Accepted:    3.1K
 * Total Submissions: 8.7K
 * Testcase Example:  '[5,10,-5]'
 *
 * 给定一个整数数组 asteroids，表示在同一行的行星。
 *
 * 对于数组中的每一个元素，其绝对值表示行星的大小，正负表示行星的移动方向（正表示向右移动，负表示向左移动）。每一颗行星以相同的速度移动。
 *
 *
 * 找出碰撞后剩下的所有行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。
 *
 * 示例 1:
 *
 *
 * 输入:
 * asteroids = [5, 10, -5]
 * 输出: [5, 10]
 * 解释:
 * 10 和 -5 碰撞后只剩下 10。 5 和 10 永远不会发生碰撞。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * asteroids = [8, -8]
 * 输出: []
 * 解释:
 * 8 和 -8 碰撞后，两者都发生爆炸。
 *
 *
 * 示例 3:
 *
 *
 * 输入:
 * asteroids = [10, 2, -5]
 * 输出: [10]
 * 解释:
 * 2 和 -5 发生碰撞后剩下 -5。10 和 -5 发生碰撞后剩下 10。
 *
 *
 * 示例 4:
 *
 *
 * 输入:
 * asteroids = [-2, -1, 1, 2]
 * 输出: [-2, -1, 1, 2]
 * 解释:
 * -2 和 -1 向左移动，而 1 和 2 向右移动。
 * 由于移动方向相同的行星不会发生碰撞，所以最终没有行星发生碰撞。
 *
 *
 * 说明:
 *
 *
 * 数组 asteroids 的长度不超过 10000。
 * 每一颗行星的大小都是非零整数，范围是 [-1000, 1000] 。
 *
 *
 */

// @lc code=start
type Mystack []int

func (this *Mystack) push(a int) {
	(*this) = append((*this), a)
}

func (this *Mystack) isEmpty() bool {
	return len((*this)) == 0
}

func (this *Mystack) pop() int {
	length := len(*this)
	tmp := (*this)[length-1]
	(*this) = (*this)[0 : length-1]
	return tmp
}

func (this *Mystack) top() int {
	length := len(*this)
	tmp := (*this)[length-1]
	return tmp
}

func asteroidCollision(asteroids []int) []int {
	stack := Mystack{}
	for _, val := range asteroids {
		for !stack.isEmpty() && willHit(stack.top(), val) {
			topVal := stack.pop()
			if getAbs(topVal) > getAbs(val) {
				val = topVal
			} else if getAbs(topVal) == getAbs(val) {
				// 如果相等，原地爆炸，退出本轮检测
				val = 0
				break
			}
		}
		// 本轮检测结束，推入val
		if val != 0 {
			stack.push(val)
		}
	}
	// return []int(stack)
	return stack
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func willHit(a, b int) bool {
	if a > 0 && b < 0 {
		return true
	}
	return false
}

// @lc code=end

