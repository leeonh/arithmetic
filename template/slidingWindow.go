package template

/*
滑动窗口模板
伪代码:

int left = 0;
int right = 0;

while (right < s.size()) {
	// 滑动窗口右边界向右移动扩张
	window.add(s[right]);
	right++;
	while(valid) {
		// 滑动窗口左边界向右移动缩小
		window.remove(s[left])
		left++
	}
}

滑动窗口大体有两种

一种是固定长度的滑动窗口
另一种是不固定长度的

*/

/**
固定长度的滑动窗口

适用于字符串固定长度对比

例如截取字符串里面某一部分，出来和目标字符串作对比
*/
func slidingWindowFixed(s string, p string) []int {
	// 获取两个字符串的长度
	n, m := len(s), len(p)
	// 定义答案
	var ans []int
	// 新建两个数组，用来记录字符出现的次数
	var sCnt, pCnt [26]int
	// 将p字符串里出现的字符次数放进去数组
	for i := 0; i < m; i++ {
		pCnt[p[i]-'a'] += 1
	}
	// 定义两个指针指向开头
	left, right := 0, 0
	for right < n {
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
		right++
	}
	return ans
}

/*
不定长滑动窗口

例如用来寻找 乘积小于K的子数组
*/
func slidingWindow(nums []int, k int) int {
	n := len(nums)
	// 排除特殊情况，保证循环时候，left不会越界
	// 例如[1,2,3],k=0，p >= k这个条件就会变成无限循环，因为p永远大于0
	// 还有[1,1,1],k=1，这个也是，会变成无限循环，直到left超界
	if k <= 1 {
		return 0
	}
	// 定义一个变量，存储left到right的所有值的乘积
	p := 1
	// 定义滑动窗口两个指针放到开头
	left, right := 0, 0
	// 定义结果，连续数组的数量
	ans := 0
	for right < n {
		// right指针移动的时候一直累计乘积
		p *= nums[right]
		// 判断乘积大于目标值的时候，也就是p>=k的时候，左指针要往右指针靠拢，直到left和right直接数字的乘积小于k为止
		for p >= k {
			// left一直向right靠拢，每移动一次，乘积就除以nums[left]，除了以后就相当于left右移后left到right之间所有数字的乘积了
			p /= nums[left]
			// left指针右移
			left++
		}
		ans += right - left + 1
		right++
	}
	return ans
}
