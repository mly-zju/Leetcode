func reverseWords(s string) string {
	cacheStr := ""
	cache := []string{}
	for _, ch := range s {
		by := byte(ch)
		if by == ' ' && cacheStr != "" {
			cache = append(cache, cacheStr)
			cacheStr = ""
		} else if by != ' ' {
			cacheStr += string(by)
		}
	}
	if cacheStr != "" {
		cache = append(cache, cacheStr)
	}
	
	res := ""
	length := len(cache)
	for i := length - 1; i >= 0; i-- {
		if i != 0 {
			res += cache[i] + " "
		} else {
			res += cache[i]
		}
	}
	return res
}