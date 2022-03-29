package day04

func ReverseWords(s string) string {
	n := len(s)
	res := []byte{}
	for r := 0; r <= n-1; {
		l := r
		for r < n && s[r] != ' ' {
			r++
		}
		for p := l; p < r; p++ {
			//r-1是因为r已经停在了空格上面了，往左边移一个单位才是单词结尾
			res = append(res, s[r-1-p+l])
		}
		for r < n && s[r] == ' ' {
			r++
			res = append(res, ' ')
		}
	}
	return string(res)
}
