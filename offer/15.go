func hammingWeight(n int) int {
	res := 0
	for n > 0 {
		tmp := n % 2
		if tmp == 1 {
			res++
		}
		n = n / 2
	}
	return res
}
