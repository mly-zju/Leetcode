/*
 * @lc app=leetcode.cn id=1010 lang=golang
 *
 * [1010] 总持续时间可被 60 整除的歌曲
 *
 * https://leetcode-cn.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/description/
 *
 * algorithms
 * Easy (41.69%)
 * Likes:    53
 * Dislikes: 0
 * Total Accepted:    6.6K
 * Total Submissions: 15.4K
 * Testcase Example:  '[30,20,150,100,40]'
 *
 * 在歌曲列表中，第 i 首歌曲的持续时间为 time[i] 秒。
 * 
 * 返回其总持续时间（以秒为单位）可被 60 整除的歌曲对的数量。形式上，我们希望索引的数字  i < j 且有 (time[i] + time[j]) %
 * 60 == 0。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[30,20,150,100,40]
 * 输出：3
 * 解释：这三对的总持续时间可被 60 整数：
 * (time[0] = 30, time[2] = 150): 总持续时间 180
 * (time[1] = 20, time[3] = 100): 总持续时间 120
 * (time[1] = 20, time[4] = 40): 总持续时间 60
 * 
 * 
 * 示例 2：
 * 
 * 输入：[60,60,60]
 * 输出：3
 * 解释：所有三对的总持续时间都是 120，可以被 60 整数。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= time.length <= 60000
 * 1 <= time[i] <= 500
 * 
 * 
 */

// @lc code=start
func numPairsDivisibleBy60(time []int) int {
	// 暴力法当然可以。为了简化，基于统计hash，这里就可以对数组执行预处理，所有都取余数
	length := len(time)
	if length == 0 {
		return 0
	}
	statMap := map[int]int{}
	for i := 0; i < length; i++ {
		tmp := time[i] % 60
		time[i] = tmp
		statMap[tmp] += 1
	}

	res := 0
	for _, val := range time {
		// 先要减自身
		statMap[val]--
		target := (60 - val) % 60
		if statMap[target] > 0 {
			res += statMap[target]
		}
		statMap[val]++
	}
	return res / 2
}
// @lc code=end

