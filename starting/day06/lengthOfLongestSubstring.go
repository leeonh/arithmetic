package day06

/**map方法*/
func lengthOfLongestSubstring1(s string) int {
	ret := map[byte]int{}
	n := len(s)
	r, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			//在map中移除l的数据
			ret[s[i-1]] = 0
		}
		// 右指针一直循环到重复的字符才停止
		for r+1 < n && ret[s[r+1]] == 0 {
			ret[s[r+1]] = 1
			r++
		}
		ans = max(ans, r-i+1)
	}
	return ans
}

func LengthOfLongestSubstring(s string) int {
	n := len(s)
	m := make([]int, 128)
	r, l := 0, 0
	var num int
	for r = 0; r < n; r++ {
		// 如果不存在，说明没有重复字符，左指针保持不动，如果 m[s[r]] 存在，说明字符是重复的，将滑动窗口移动到当前重复字符的位置，重新开始累计无重复字符长度
		l = max(l, m[s[r]])
		// 计算滑动窗口的左指针与右指针之间无重复字符的数量，与上一次的数据作对比，取最大值
		num = max(num, r-l+1)
		// 记录当前字符位置的下一个位置，配合上面如果发现重复字符了，m[s[r]]拿出来的值才是当前循环r(右指针)的值
		m[s[r]] = r + 1
	}
	return num
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
