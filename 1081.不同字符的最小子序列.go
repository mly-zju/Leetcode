package main

/*
 * @lc app=leetcode.cn id=1081 lang=golang
 *
 * [1081] 不同字符的最小子序列
 *
 * https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters/description/
 *
 * algorithms
 * Medium (40.21%)
 * Likes:    9
 * Dislikes: 0
 * Total Accepted:    886
 * Total Submissions: 2.1K
 * Testcase Example:  '"cdadabcc"'
 *
 * 返回字符串 text 中按字典序排列最小的子序列，该子序列包含 text 中所有不同字符一次。
 *
 *
 *
 * 示例 1：
 *
 * 输入："cdadabcc"
 * 输出："adbc"
 *
 *
 * 示例 2：
 *
 * 输入："abcd"
 * 输出："abcd"
 *
 *
 * 示例 3：
 *
 * 输入："ecbacba"
 * 输出："eacb"
 *
 *
 * 示例 4：
 *
 * 输入："leetcode"
 * 输出："letcod"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= text.length <= 1000
 * text 由小写英文字母组成
 *
 *
 */
// @lc code=start
type Mystack []byte

func (this *Mystack) push(a byte) {
	(*this) = append((*this), a)
}

func (this *Mystack) isEmpty() bool {
	return len((*this)) == 0
}

func (this *Mystack) pop() byte {
	length := len(*this)
	tmp := (*this)[length-1]
	(*this) = (*this)[0 : length-1]
	return tmp
}

func (this *Mystack) top() byte {
	length := len(*this)
	tmp := (*this)[length-1]
	return tmp
}

func smallestSubsequence(text string) string {
	// 有条件的单调栈。最前面的越小越好，但是条件是越大的那个在后面还会出现。因此先做统计。
	// 此外每个字符只需要出现一次，因此还要去重。
	charMap := map[byte]int{}
	existMap := map[byte]bool{}
	length := len(text)
	for i := 0; i < length; i++ {
		charMap[text[i]]++
	}

	stack := Mystack{}
	for i := 0; i < length; i++ {
		curChar := text[i]
		charMap[curChar]--
		// 如果之前已经存在，则不进行尝试了
		if existMap[curChar] {
			continue
		}
		for !stack.isEmpty() && curChar < stack.top() && charMap[stack.top()] > 0 {
			tmp := stack.pop()
			existMap[tmp] = false
		}
		stack.push(curChar)
		existMap[curChar] = true
	}
	return string(stack)
}

// func main() {
// 	// a := "cbaacabcaaccaacababa"
// 	// abc is right
// 	// a := "baababaaaaababbbbbbaababaababa"
// 	// ab is right
// 	fmt.Println(smallestSubsequence(a))
// }

// @lc code=end
