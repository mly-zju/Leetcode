/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 *
 * https://leetcode-cn.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (49.58%)
 * Likes:    321
 * Dislikes: 0
 * Total Accepted:    35.9K
 * Total Submissions: 71K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
 * 
 * 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
 * 
 * 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: 2, [[1,0]] 
 * 输出: true
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
 * 
 * 示例 2:
 * 
 * 输入: 2, [[1,0],[0,1]]
 * 输出: false
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
 * 你可以假定输入的先决条件中没有重复的边。
 * 1 <= numCourses <= 10^5
 * 
 * 
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1. 深度优先搜索判断是否是有向无环图
	// 先构造领接表
	adJustMap := map[int][]int{}
	length := len(prerequisites)
	for i := 0; i < length; i++ {
		startNode := prerequisites[i][0]
		prerequisites[i] = prerequisites[i][1:]
		if adJustMap[startNode] == nil {
			adJustMap[startNode] = prerequisites[i]
		} else {
			adJustMap[startNode] = append(adJustMap[startNode], prerequisites[i]...)
		}
	}

	// 判断是否是有向无环图
	// seenMap -1表示被本轮其他节点访问过，1表示已经访问完成, 防止重复访问
	seenMap := map[int]int{}
	for i := 0; i < numCourses; i++ {
		// 每轮如果遇到命中的，说明成环
		res := dfs(adJustMap, seenMap, i)
		if res == false {
			return false
		}
	}
	return true
}

func dfs(adJustMap map[int][]int, seenMap map[int]int, src int) bool {
	if seenMap[src] == 1 {
		return true
	} else if seenMap[src] == -1 {
		return false
	}

	// 标记src被本轮访问过
	seenMap[src] = -1
	// 访问src的子节点
	for _, vChild := range adJustMap[src] {
		res := dfs(adJustMap, seenMap, vChild)
		if res == false {
			return false
		}
	}

	// src访问完成，标记一轮完全完成
	seenMap[src] = 1

	return true
}
// @lc code=end

