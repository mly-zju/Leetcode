func replaceSpace(s string) string {
	res := ""
	for i, length := 0, len(s); i < length; i++ {
		if s[i] == ' ' {
			res += "%20"
		} else {
			res += string(s[i])
		}
	}
	return res
}