/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 *
 * https://leetcode-cn.com/problems/relative-sort-array/description/
 *
 * algorithms
 * Easy (62.19%)
 * Likes:    13
 * Dislikes: 0
 * Total Accepted:    5.6K
 * Total Submissions: 8.9K
 * Testcase Example:  '[2,3,1,3,2,4,6,7,9,2,19]\n[2,1,4,3,9,6]'
 *
 * 给你两个数组，arr1 和 arr2，
 *
 *
 * arr2 中的元素各不相同
 * arr2 中的每个元素都出现在 arr1 中
 *
 *
 * 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1
 * 的末尾。
 *
 *
 *
 * 示例：
 *
 * 输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
 * 输出：[2,2,2,1,4,3,3,9,6,7,19]
 *
 *
 *
 *
 * 提示：
 *
 *
 * arr1.length, arr2.length <= 1000
 * 0 <= arr1[i], arr2[i] <= 1000
 * arr2 中的元素 arr2[i] 各不相同
 * arr2 中的每个元素 arr2[i] 都出现在 arr1 中
 *
 *
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	statArr := [1001]int{}
	res := []int{}
	// 计数统计arr1
	for _, val := range arr1 {
		statArr[val]++
	}
	// 先基于arr2排序
	for _, val := range arr2 {
		for statArr[val] > 0 {
			statArr[val]--
			res = append(res, val)
		}
	}
	// 再对剩余进行升序排序
	for i := 0; i < 1001; i++ {
		for statArr[i] > 0 {
			res = append(res, i)
			statArr[i]--
		}
	}
	return res
}

// @lc code=end

