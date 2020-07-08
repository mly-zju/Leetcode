func movingCount(m int, n int, k int) int {
	// 回溯法
	if m == 0 || n == 0 {
		return 0
	}

	count := 1
	statMap := make(map[string]bool)
	statMap[getKey(0, 0)] = true
	return dfs(0, 0, m, n, k, count, statMap)
}

func dfs(i, j, m, n, k, count int, statMap map[string]bool) int {
	// 上
	if i - 1 >= 0 && !statMap[getKey(i-1, j)] && getSum(i-1, j) <= k {
		count++
		statMap[getKey(i-1, j)] = true
		count = dfs(i-1, j, m, n, k, count, statMap)
	}
	// 下
	if i + 1 < m && !statMap[getKey(i+1, j)] && getSum(i+1, j) <= k {
		count++
		statMap[getKey(i+1, j)] = true
		count = dfs(i+1, j, m, n, k, count, statMap)
	}
	// 左
	if j - 1 >= 0 && !statMap[getKey(i, j-1)] && getSum(i, j-1) <= k {
		count++
		statMap[getKey(i, j-1)] = true
		count = dfs(i, j-1, m, n, k, count, statMap)
	}
	// 右
	if j + 1 < n && !statMap[getKey(i, j+1)] && getSum(i, j+1) <= k {
		count++
		statMap[getKey(i, j+1)] = true
		count = dfs(i, j+1, m, n, k, count, statMap)
	}
	return count
}

func getKey(i, j int) string {
	return strconv.Itoa(i) + "-" + strconv.Itoa(j)
}

func getSum(i, j int) int {
	sum := 0
	for i > 0 {
		sum += i % 10
		i = i / 10
	}
	for j > 0 {
		sum += j % 10
		j = j / 10
	}
	return sum
}