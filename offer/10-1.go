func fib(n int) int {
	// 1. 递归: 会超时，不采用
	// if n <= 1 {
	// 	return n
	// } 
	// return fib(n-1) + fib(n-2)

	// 2. 迭代
	cache := []int{}
	for i := 0; i <= n; i++ {
		if i <= 1 {
			cache = append(cache, i)
		} else {
			cache = append(cache, (cache[i-1] + cache[i-2]) % (1e9 + 7))
		}
	}
	return cache[len(cache) - 1]
}