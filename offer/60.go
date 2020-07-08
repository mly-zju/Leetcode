func twoSum(n int) []float64 {
	// 根据n计算有多少种可能性
	nNum := n * 6 - n + 1
	res := make([]float64, nNum)
	
	// tmpCache存放点数集合
	tmpCache := make([]int, n)
	// 回溯法
	res = dfs(1, n, tmpCache, res)
	// 统一截取4位小数
	// for i := 0; i < nNum; i++ {
	// 	res[i] = trimN(res[i], 5)
	// }
	return res
}

func dfs(round, maxRound int, tmpCache []int, res []float64) []float64 {
	if round > maxRound {
		sum, chance := getSumAndChance(tmpCache)
		// 给相应的点数增加概率
		res[sum-maxRound] += chance
	} else {
		// 每一轮骰子从1到6执行
		for i := 1; i <= 6; i++ {
			tmpCache[round - 1] = i
			res = dfs(round + 1, maxRound, tmpCache, res)	
		}
	}
	return res
}

func getSumAndChance(arr []int) (int, float64) {
	var sum int
	var chance float64 = 1.0
	for _, num := range arr {
		sum += num
		chance = chance * float64(0.16667)
	}
	return sum, chance
}

func trimN(x float64, n int) float64 {
	numArr := []int{}
	xC := int(x)
	xN := int(x * math.Pow(10, float64(n)))
	for i := 0; i < n; i++ {
		tmp := xN % 10
		numArr = append(numArr, tmp)
		xN = xN / 10
	}

	count := 10
	res := float64(xC)
	for i := n - 1; i >= 0; i-- {
		res += float64(numArr[i]) / float64(count)
		count = count * 10
	}	
	return res
}