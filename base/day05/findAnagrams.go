package day05

/*
题目：找到字符串中所有字母异位词

给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。



示例1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
*/

func FindAnagrams(s string, p string) (ans []int) {
	// 获取两个字符串的长度
	n, m := len(s), len(p)
	if n < m {
		return []int{}
	}
	// 新建两个数组，用来记录字符出现的次数
	var sCnt, pCnt [26]int
	// 将p字符串里出现的字符次数放进去数组
	for i := 0; i < m; i++ {
		pCnt[p[i]-'a'] += 1
		sCnt[s[i]-'a'] += 1
	}

	// 切片长度要+1，例如s=abab,p=ab，则切片后就是[a,b]了，但是后面还有一组ab，就会漏掉，所以要+1
	//for i := 0; i < len(s[:n-m])+1; i++ {
	//	// 比较滑块内的字符串，也就是判断两个数组是否相等，如果相等，说明字符串一样
	//	if sCnt == pCnt {
	//		ans = append(ans, i)
	//	}
	//
	//	// 判断滑块是否刚好超界
	//	if i+m == n {
	//		break
	//	}
	//	sCnt[s[i]-'a'] -= 1
	//	sCnt[s[i+m]-'a'] += 1
	//
	//}

	// 当然也可以提前判断，第一次把两个字符串扫描进数组的时候就判断一次
	// 相当于判断索引0的时候，滑动窗口内的字符串是否和目标相等
	if sCnt == pCnt {
		ans = append(ans, 0)
	}
	// 这样就可以变成，滑块先移动一步，再进行对比了，也不用担心超界
	for i := 0; i < len(s[:n-m]); i++ {
		sCnt[s[i]-'a'] -= 1
		sCnt[s[i+m]-'a'] += 1
		// 比较滑块内的字符串，也就是判断两个数组是否相等，如果相等，说明字符串一样
		if sCnt == pCnt {
			// 因为滑块是先移动了再对比的，所以相当于比对的是滑块左边界+1后的字符串
			ans = append(ans, i+1)
		}

	}

	return
}

/*滑动窗口加双指针优化*/
func FindAnagrams1(s string, p string) (ans []int) {
	// 获取两个字符串的长度
	n, m := len(s), len(p)
	// 新建两个数组，用来记录字符出现的次数
	var sCnt, pCnt [26]int
	// 将p字符串里出现的字符次数放进去数组
	for i := 0; i < m; i++ {
		pCnt[p[i]-'a'] += 1
	}
	// 使用两个指针，分别left就是滑动窗口的左边界，right就是滑动窗口的右边界，循环结束条件就是，右边界碰到了字符串s的最后一个字符，因为再往后就超界了
	// 首先右边界一直向右移动，同时将扫描的字符加入数组，当左边界和右边界的距离等于目标字符串长度的时候，就将字符串进行比对
	for left, right := 0, 0; right < n; right++ {
		// right每次+1后，判断left和right之间的距离等于目标字符串的长度
		sCnt[s[right]-'a']++
		if right-left+1 > m {
			// 如果左边界和右边界的距离大于目标值长度
			// 先将当前left位置的字符去掉，再移动左边界
			// 这样就能保证滑动窗口的长度等于目标值长度，从而进入到下面的对比判断里面
			sCnt[s[left]-'a']--
			left++
		}
		if right-left+1 == m {
			// 如果区间长度和目标字符串长度一致，就对比两个字符是否相等
			if sCnt == pCnt {
				// 相等说明字符都一样，是异位词
				ans = append(ans, left)
			}
		}
	}

	return
}
