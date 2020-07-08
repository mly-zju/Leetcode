/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 *
 * https://leetcode-cn.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (41.97%)
 * Likes:    300
 * Dislikes: 0
 * Total Accepted:    52.5K
 * Total Submissions: 125.2K
 * Testcase Example:  '"2"\n"3"'
 *
 * 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
 *
 * 示例 1:
 *
 * 输入: num1 = "2", num2 = "3"
 * 输出: "6"
 *
 * 示例 2:
 *
 * 输入: num1 = "123", num2 = "456"
 * 输出: "56088"
 *
 * 说明：
 *
 *
 * num1 和 num2 的长度小于110。
 * num1 和 num2 只包含数字 0-9。
 * num1 和 num2 均不以零开头，除非是数字 0 本身。
 * 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 *
 *
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	// 单字符串乘法 + 字符串加法结合
	singleMultiArr := []string{}
	len2 := len(num2)

	// 计算每一个单字符乘法结果，注意每一次都要多加0
	postFix := ""
	for i := len2 - 1; i >= 0; i-- {
		singleMultiArr = append(singleMultiArr, singleMulti(num1, string(num2[i]))+postFix)
		postFix += "0"
	}

	// 循环执行加法操作
	res := "0"
	for _, str := range singleMultiArr {
		res = strAdd(res, str)
	}

	// 处理前缀0
	byteArr := []byte(res)
	preZeroNum := 0
	for _, by := range byteArr {
		if by == '0' {
			preZeroNum++
		} else {
			break
		}
	}
	byteArr = byteArr[preZeroNum:]
	if len(byteArr) == 0 {
		return "0"
	}
	return string(byteArr)
}

func singleMulti(num1 string, num2 string) string {
	len1 := len(num1)
	cache := []string{}
	n2, _ := strconv.Atoi(num2)
	flag := 0
	for i := len1 - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		tmp := n1*n2 + flag
		flag = 0
		if tmp >= 10 {
			flag = tmp / 10
			tmp = tmp % 10
		}
		cache = append(cache, strconv.Itoa(tmp))
	}
	// 最后还要处理一下flag
	if flag > 0 {
		cache = append(cache, strconv.Itoa(flag))
	}

	// 组装
	res := ""
	for _, ch := range cache {
		res = ch + res
	}
	return res
}

func strAdd(a, b string) string {
	lena, lenb := len(a), len(b)
	i, j := lena-1, lenb-1
	cache := []string{}
	flag := 0
	for i >= 0 && j >= 0 {
		tmp := int(a[i]-'0') + int(b[j]-'0') + flag
		flag = 0
		if tmp >= 10 {
			flag = 1
			tmp = tmp % 10
		}
		cache = append(cache, strconv.Itoa(tmp))
		i--
		j--
	}

	for i >= 0 {
		tmp := int(a[i]-'0') + flag
		flag = 0
		if tmp >= 10 {
			flag = 1
			tmp = tmp % 10
		}
		cache = append(cache, strconv.Itoa(tmp))
		i--
	}
	for j >= 0 {
		tmp := int(b[j]-'0') + flag
		flag = 0
		if tmp >= 10 {
			flag = 1
			tmp = tmp % 10
		}
		cache = append(cache, strconv.Itoa(tmp))
		j--
	}

	// 最后还要处理flag
	if flag == 1 {
		cache = append(cache, "1")
	}

	res := ""
	for _, str := range cache {
		res = str + res
	}
	return res
}

// @lc code=end
