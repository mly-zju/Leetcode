func sumNums(n int) int {
	a := 0
	for i := 1; i <= n; i++ {
		a = sum(a, i)
	}
	return a
}

func sum(a int, b int) int {
	for b != 0 {
		tmp := (a << 1) & (b << 1)
		a = a ^ b
		b = tmp
	}
	return a
}
