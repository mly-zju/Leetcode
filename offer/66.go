func constructArr(a []int) []int {
	res := []int{}
	length := len(a)
	if length == 0 {
		return res
	}
	// 基于缓存数组, prev[i]代表a[0] * a[1] ... * a[i], post[i]代表a[len-1] * a[len-2] .. * a[i]
	// 则res[i] = prev[i-1] * post[i+1]
	prev := make([]int, length)
	post := make([]int, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			prev[i] = a[i]
		} else {
			prev[i] = prev[i-1] * a[i]
		}
	}
	for i := length - 1; i >= 0; i-- {
		if i == length - 1 {
			post[i] = a[i]
		} else {
			post[i] = post[i+1] * a[i]
		}
	}
	for i := 0; i < length; i++ {
		if i == 0 {
			res = append(res, post[i + 1])
		} else if i == length - 1 {
			res = append(res, prev[i - 1])
		} else {
			res = append(res, prev[i-1] * post[i+1])
		}
	}
	return res
}