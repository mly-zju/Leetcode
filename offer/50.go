type StatInfo struct {
	firstIndex int
	num int
}

func firstUniqChar(s string) byte {
	// 执行统计
	statMap := make(map[byte]*StatInfo)
	for i, length := 0, len(s); i < length; i++ {
		if statMap[s[i]] == nil {
			statMap[s[i]] = &StatInfo{
				firstIndex: i,
				num: 1,
			}
		} else {
			statMap[s[i]].num++
		}
	}

	// 遍历寻找
	res := byte(' ')
	minIndex := len(s)
	for key, info := range statMap {
		if info.num == 1 && info.firstIndex < minIndex {
			minIndex = info.firstIndex
			res = key
		}
	}
	return res
}