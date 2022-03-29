package day06

func CheckInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	// 创建两个对比数组
	count1, count2 := [26]int{}, [26]int{}
	for i, v := range s1 {
		// 将s1每个字符出现的次数统计到count1
		count1[v-'a']++

		// 截取s2字符串s1长度的字符作为滑动窗口，将滑动窗口内的字符放进去count2
		count2[s2[i]-'a']++
	}

	//如果两个数组相等，则返回true
	if count1 == count2 {
		return true
	}

	// i-n就是滑动窗口左指针左边界的位置，i=n就是右指针滑动窗口的右边界，开头和结尾向右循环
	for i := n; i < m; i++ {
		//右边界每次向右移动，对应字符在数组里面就加一
		count2[s2[i]-'a']++
		//左边界每次向右移动，对应字符在数组里面就减一，剔除一次该字符出现次数
		count2[s2[i-n]-'a']--
		if count1 == count2 {
			return true
		}
	}
	return false

}
