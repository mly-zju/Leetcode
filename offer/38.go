// 回溯法
func permutation(s string) []string {
	sArr := []byte(s)
	res := []string{}
	length := len(sArr)
	if length == 0 {
		return res
	}
	visitArr := make([]bool, length)
	repeatMap := make(map[string]bool)
	// 执行回溯
	return dfs(sArr, visitArr, "", res, repeatMap)
}

func dfs(sArr []byte, visitArr []bool, tmpStr string, res []string, repeatMap map[string]bool) []string {
	if len(tmpStr) == len(sArr) {
		// 注意还需要去重
		if repeatMap[tmpStr] == true {
			return res
		}
		res = append(res, tmpStr)
		repeatMap[tmpStr] = true
		return res
	} else {
		for index, ch := range sArr {
			if !visitArr[index] {
				visitArr[index] = true
				newStr := tmpStr + string(ch)
				res = dfs(sArr, visitArr, newStr, res, repeatMap)
				visitArr[index] = false
			}
		}
		return res
	}
}