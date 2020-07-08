/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 *
 * https://leetcode-cn.com/problems/sort-colors/description/
 *
 * algorithms
 * Medium (52.90%)
 * Likes:    266
 * Dislikes: 0
 * Total Accepted:    43K
 * Total Submissions: 80K
 * Testcase Example:  '[2,0,2,1,1,0]'
 *
 * 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
 *
 * 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
 *
 * 注意:
 * 不能使用代码库中的排序函数来解决这道题。
 *
 * 示例:
 *
 * 输入: [2,0,2,1,1,0]
 * 输出: [0,0,1,1,2,2]
 *
 * 进阶：
 *
 *
 * 一个直观的解决方案是使用计数排序的两趟扫描算法。
 * 首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
 * 你能想出一个仅使用常数空间的一趟扫描算法吗？
 *
 *
 */

// @lc code=start
func sortColors(nums []int) {
	// 原地排序。最简单一遍扫描统计，二遍扫描修改。
	// 第二种：一遍扫描，基于快排的partition2算法的发散版本，需要大的放在尾部小的放在头部, 且i和r不能相遇
	length := len(nums)
	if length == 0 {
		return
	}
	l, r, i := 0, length - 1, 0
	for i <= r {
		if nums[i] == 0 {
			// 遇到0，放到左边, 且新的nums[i]一定等于1(因为如果等于2，会换到右边), 因此i直接++
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
			// 这里i就不能++了，因为不确定新的nums[i]属于哪个范畴
		} else {
			i++
		}
	}















	// length := len(nums)
	// l, r := 0, length - 1
	// for i := 0; i <= r; {
	// 	if nums[i] < 1 {
	// 		nums[l], nums[i] = nums[i], nums[l]
	// 		l++
	// 		// 这里i需要++, 因为左边一定都是1，不可能是2，2都换到右边了
	// 		i++
	// 	} else if nums[i] > 1 {
	// 		nums[r], nums[i] = nums[i], nums[r]
	// 		r--
	// 	} else {
	// 		i++
	// 	}
	// }
}

// @lc code=end

