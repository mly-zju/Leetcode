func myPow(x float64, n int) float64 {
	// 1. 傻傻的迭代计算，可能会超时
	// 2. myPow(x, n) = myPow(x, n/2) * myPow(x, n/2), 或者如果n是奇数，还要再乘以x
	// 注意处理负数

	if n < 0 {
		x = 1 / x
		n = -n
	}

	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}

	base := myPow(x, n / 2)
	if n % 2 == 1 {
		return base * base * x
	}
	return base * base
}