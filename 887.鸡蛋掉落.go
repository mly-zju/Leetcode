/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 *
 * https://leetcode-cn.com/problems/super-egg-drop/description/
 *
 * algorithms
 * Hard (21.50%)
 * Likes:    204
 * Dislikes: 0
 * Total Accepted:    9.6K
 * Total Submissions: 44.2K
 * Testcase Example:  '1\n2'
 *
 * 你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。
 * 
 * 每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。
 * 
 * 你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。
 * 
 * 每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。
 * 
 * 你的目标是确切地知道 F 的值是多少。
 * 
 * 无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：K = 1, N = 2
 * 输出：2
 * 解释：
 * 鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
 * 否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
 * 如果它没碎，那么我们肯定知道 F = 2 。
 * 因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
 * 
 * 
 * 示例 2：
 * 
 * 输入：K = 2, N = 6
 * 输出：3
 * 
 * 
 * 示例 3：
 * 
 * 输入：K = 3, N = 14
 * 输出：4
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= K <= 100
 * 1 <= N <= 10000
 * 
 * 
 */

// @lc code=start
func superEggDrop(K int, N int) int {
	// 这里我们不关心是不是从1楼开始扔，而是只关心有多少层楼。
	// dp[k][n]代表k个鸡蛋n层楼情况下的最优解法
	// 这里需要用到记忆化，不是简单的dp table可以解决的, 主要是因为dp[i][j]依赖的前置项可能大于i或者j，不是顺序的, 这个时候备忘录也不一定非要二维数组，完全可以k+n组成一个key来做
	
	// 先定义备忘录
	memo := [][]int{}
	// 初始化, 一开始都置为-1
	for i := 0; i <= K; i++ {
		lineArr := []int{}
		for j := 0; j <= N; j++ {
			lineArr = append(lineArr, -1)
		}
		memo = append(memo, lineArr)
	}

	return getDp(K, N, memo)
}

// 递归+记忆化求dp[k][n]
func getDp(k int, n int, memo [][]int) int {
	// 先看记忆化
	if memo[k][n] != -1 {
		return memo[k][n]
	}

	// 看边界条件
	if n == 0 {
		memo[k][n] = 0
		return 0
	} else if k == 0 {
		// 如果没有鸡蛋，那么返回一个极大值即可
		memo[k][n] = math.MaxInt32
		return math.MaxInt32
	} else if k == 1 {
		// 如果1个鸡蛋, 那么只能从底层向上扔，最坏结果N次
		memo[k][n] = n
		return n
	}

	// 除此之外，dp需要递归求解
	res := math.MaxInt32
	// 这里顺序搜索竟然会超时！因此走二分查找
	// for i := 1; i <= n; i++ {
		// 迭代法，会超时
		// res = getMin(res, getMax(getDp(k-1, i-1, memo), getDp(k, n-i, memo)) + 1)
	// }

	// 二分查找这里也很tricky，dp[k-1][i-1]随着i递增，dp[k][n-i]随着i递减，我们要找getMin, 而内部是getMax, 那么
	// 肯定两者相等时候求出的getMax最小，因此某一个值大了，就朝着让它缩小的地方取
	l, r := 1, n
	var mid int
	for l <= r {
		mid = (l + r) / 2
		// 不碎
		broken := getDp(k, n-mid, memo)
		// 碎
		notBroken := getDp(k-1, mid-1, memo)
		if broken > notBroken {
			// 为了减小broken，需要增大mid
			l = mid + 1
			res = getMin(res, broken + 1)
		} else {
			// 为了增加notBroken, 需要减小mid
			r = mid - 1
			res = getMin(res, notBroken + 1)
		}
	}

	// 备忘
	memo[k][n] = res
	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

