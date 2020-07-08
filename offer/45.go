func minNumber(nums []int) string {
	// 两数比较，按照组合字符串比较，比如'30', '20' 比较，按照'3020' 和'2030'来按位比
	sort.Slice(nums, func(i, j int) bool {
		is, js := getStr(nums[i]), getStr(nums[j])
		str1, str2 := is+js, js+is
		return str1 <= str2
	})
	fmt.Println(nums)
	res := ""
	for _, num := range nums {
		res += getStr(num)
	}
	return res
}

func getStr(a int) string {
	res := []byte{}
	for a > 0 {
		tmp := a % 10
		res = append(res, strconv.Itoa(tmp)[0])
		a = a / 10
	}

	length := len(res)
	if length == 0 {
		return "0"
	}
	// 反转
	l, r := 0, length-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return string(res)
}

