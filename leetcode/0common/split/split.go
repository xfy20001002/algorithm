package split

func Split(s string) []string {
	res := make([]string, 0)
	record := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = append(res, s[record:i])
			record = i + 1
		}
	}
	res = append(res, s[record:])
	return res
}
