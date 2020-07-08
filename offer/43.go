/*
总体思想就是分类，先求所有数中个位是 1 的个数，再求十位是 1 的个数，再求百位是 1 的个数...

假设 n = xyzdabc，此时我们求千位是 1 的个数，也就是 d 所在的位置。

那么此时有三种情况，

d == 0，那么千位上 1 的个数就是 xyz * 1000
d == 1，那么千位上 1 的个数就是 xyz * 1000 + abc + 1
d > 1，那么千位上 1 的个数就是 xyz * 1000 + 1000

当我们考虑千位是 1 的时候，我们将千位定为 1，也就是 xyz1abc。

对于 xyz 的话，可以取 0,1,2...(xyz-1)，也就是 xyz 种可能。

当 xyz 固定为上边其中的一个数的时候，abc 可以取 0,1,2...999，也就是 1000 种可能。

这样的话，总共就是 xyz*1000 种可能。

注意到，我们前三位只取到了 xyz-1，那么如果取 xyz 呢？

此时就出现了上边的三种情况，取决于 d 的值。

d == 1 的时候，千位刚好是 1，此时 abc 可以取的值就是 0 到 abc ，所以多加了 abc + 1。

d > 1 的时候，d 如果取 1，那么 abc 就可以取 0 到 999，此时就多加了 1000。
*/

func countDigitOne(n int) int {
	res := 0
	cur, right, bitCount := 0, 0, 0
	for n > 0 {
		cur = n % 10
		n = n / 10

		if cur == 0 {
			res += n * int(math.Pow(10, float64(bitCount)))
		} else if cur == 1 {
			res += n * int(math.Pow(10, float64(bitCount))) + right + 1
		} else {
			res += (n + 1) * int(math.Pow(10, float64(bitCount)))
		}

		right = cur * int(math.Pow(10, float64(bitCount))) + right
		bitCount++
	}
	return res
}