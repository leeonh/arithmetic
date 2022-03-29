package day04

/*
给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。

注意：如果对空文本输入退格字符，文本继续为空。



示例 1：

输入：s = "ab#c", t = "ad#c"
输出：true
解释：s 和 t 都会变成 "ac"。
示例 2：

输入：s = "ab##", t = "c#d#"
输出：true
解释：s 和 t 都会变成 ""。
示例 3：

输入：s = "a#c", t = "b"
输出：false
解释：s 会变成 "c"，但 t 仍然是 "b"。

*/

func BackspaceCompare(s string, t string) bool {
	// 两个指针，分别指向两个字符串的最后一个字符，从最后开始扫描，一个个对比，遇到一个#，就n计数，指针要从遇到不是#的字符的时候，向左移动n次
	// 只要两个字符串当前循环的字符不等#的时候，就拿出来对比
	n1, n2 := len(s)-1, len(t)-1
	skipS, skipT := 0, 0
	for n1 >= 0 || n2 >= 0 {
		// 遍历s字符串
		for n1 >= 0 {
			if s[n1] == '#' {
				skipS++
				n1--
			} else if skipS > 0 {
				// 如果计数大于0，则要继续向左移动
				n1--
				skipS--
			} else {
				// 跳出循环进行比对
				break
			}
		}

		// 遍历t字符串
		for n2 >= 0 {
			if t[n2] == '#' {
				skipT++
				n2--
			} else if skipT > 0 {
				// 如果计数大于0，则要继续向左移动
				n2--
				skipT--
			} else {
				// 跳出循环进行比对
				break
			}
		}
		//如果两个字符串都还没遍历完，对比两个字符，如果不相等，直接返回false
		if n1 >= 0 && n2 >= 0 {
			if s[n1] != t[n2] {
				return false
			}
		} else if n1 >= 0 || n2 >= 0 {
			// 只要其中一个遍历完了，就说明肯定不会相等了
			return false
		}
		n1--
		n2--

	}
	return true
}
