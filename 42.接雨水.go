/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (49.55%)
 * Likes:    977
 * Dislikes: 0
 * Total Accepted:    69.1K
 * Total Submissions: 139.4K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 
 * 
 * 
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 * 
 * 示例:
 * 
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 * 
 */

// @lc code=start
func trap(height []int) int {
	// 每个柱子，找下一个高度大于等于它的柱子，组成一个区间，看最小值，然后扫描中间柱子和最小值差别，加上即可。
	// 但是需要注意的是要扫描2遍：第一遍从前向后，第二遍从后向前，这样不会漏掉, 但是：两个柱子相等的情况会重复，因此第二遍向后扫描只扫描大于的，而不是大于等于的
	// length := len(height)
	// if length < 3 {
	// 	return 0
	// }

	// res := 0
	// for i := 0; i < length; {
	// 	curH := height[i]
	// 	if curH == 0 {
	// 		i++
	// 	} else {
	// 		// 寻找下一个大于等于的柱子
	// 		nexti := -1
	// 		for j := i + 1; j < length; j++ {
	// 			if height[j] >= curH {
	// 				nexti = j
	// 				break
	// 			}
	// 		}

	// 		// 如果没找到，直接停止扫描
	// 		if nexti == -1 {
	// 			break
	// 		} else {
	// 			minH := getMin(curH, height[nexti])
	// 			for k := i + 1; k < nexti; k++ {
	// 				res += minH - height[k]
	// 			}
	// 			i = nexti
	// 		}
	// 	}
	// }

	// // 在从后向前扫描一次
	// for i := length - 1; i >= 0; {
	// 	curH := height[i]
	// 	if curH == 0 {
	// 		i--
	// 	} else {
	// 		// 寻找下一个大于等于的柱子
	// 		nexti := length
	// 		for j := i - 1; j >= 0; j-- {
	// 			if height[j] > curH {
	// 				nexti = j
	// 				break
	// 			}
	// 		}

	// 		// 如果没找到，则寻找前面
	// 		if nexti == length {
	// 			break
	// 		} else {
	// 			minH := getMin(curH, height[nexti])
	// 			for k := i - 1; k > nexti; k-- {
	// 				res += minH - height[k]
	// 			}
	// 			i = nexti
	// 		}
	// 	}
	// }
	// return res

	// 2. 解法2：维护每一个i的左右最大值，则i位置能够盛放的雨水就是min(max(i-1), max(i+1)) - height(i)
	length := len(height)
	if length < 3 {
		return 0
	}

	maxl, maxr := height[0], findMaxHeight(height, 2)
	res := 0
	// 开头和结尾无法放雨水，排除
	for i := 1; i < length - 1; i++ {
		// 更新两端max
		if height[i] > maxl {
			maxl = height[i]
		}
		if height[i] == maxr {
			maxr = findMaxHeight(height, i + 1)
		}
		tmp := getMin(maxl, maxr) - height[i]
		if tmp > 0 {
			res += tmp
		}
	}
	return res
}

func findMaxHeight(height []int, start int) int {
	max := height[start]
	for i, length := start, len(height); i < length; i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	return max
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

