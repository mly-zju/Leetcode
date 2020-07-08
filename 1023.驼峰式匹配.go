/*
 * @lc app=leetcode.cn id=1023 lang=golang
 *
 * [1023] 驼峰式匹配
 *
 * https://leetcode-cn.com/problems/camelcase-matching/description/
 *
 * algorithms
 * Medium (49.35%)
 * Likes:    10
 * Dislikes: 0
 * Total Accepted:    1.3K
 * Total Submissions: 2.6K
 * Testcase Example:  '["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"]\n"FB"'
 *
 * 如果我们可以将小写字母插入模式串 pattern 得到待查询项 query，那么待查询项与给定模式串匹配。（我们可以在任何位置插入每个字符，也可以插入
 * 0 个字符。）
 *
 * 给定待查询列表 queries，和模式串 pattern，返回由布尔值组成的答案列表 answer。只有在待查项 queries[i] 与模式串
 * pattern 匹配时， answer[i] 才为 true，否则为 false。
 *
 *
 *
 * 示例 1：
 *
 * 输入：queries =
 * ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern =
 * "FB"
 * 输出：[true,false,true,true,false]
 * 示例：
 * "FooBar" 可以这样生成："F" + "oo" + "B" + "ar"。
 * "FootBall" 可以这样生成："F" + "oot" + "B" + "all".
 * "FrameBuffer" 可以这样生成："F" + "rame" + "B" + "uffer".
 *
 * 示例 2：
 *
 * 输入：queries =
 * ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern =
 * "FoBa"
 * 输出：[true,false,true,false,false]
 * 解释：
 * "FooBar" 可以这样生成："Fo" + "o" + "Ba" + "r".
 * "FootBall" 可以这样生成："Fo" + "ot" + "Ba" + "ll".
 *
 *
 * 示例 3：
 *
 * 输出：queries =
 * ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern =
 * "FoBaT"
 * 输入：[false,true,false,false,false]
 * 解释：
 * "FooBarTest" 可以这样生成："Fo" + "o" + "Ba" + "r" + "T" + "est".
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= queries.length <= 100
 * 1 <= queries[i].length <= 100
 * 1 <= pattern.length <= 100
 * 所有字符串都仅由大写和小写英文字母组成。
 *
 *
 */

// @lc code=start
func camelMatch(queries []string, pattern string) []bool {
	res := []bool{}
	for _, str := range queries {
		res = append(res, isSub(str, pattern))
	}
	return res
}

func isSub(a, b string) bool {
	lena, lenb := len(a), len(b)
	ia, ib := 0, 0
	// a要能完全匹配b中字符，且不匹配的不能是大写字母
	for ia < lena && ib < lenb {
		if a[ia] == b[ib] {
			ia++
			ib++
		} else if a[ia] >= 'A' && a[ia] <= 'Z' {
			return false
		} else {
			ia++
		}
	}

	// 分两种情况：ia完毕或者ib完毕，ia完毕要检测ib是否也已经完毕；ib完毕要检测ia后续是否有大写
	if ia == lena {
		return ib == lenb
	}
	for ia < lena {
		if a[ia] >= 'A' && a[ia] <= 'Z' {
			return false
		}
		ia++
	}
	return true
}

// @lc code=end

