package day04

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
