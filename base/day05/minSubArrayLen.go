package day05

import "math"

/*
长度最小的子数组

给定一个含有n个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组[4,3]是该条件下的长度最小的子数组。
示例 2：

输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0


*/

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	// 定义一个变量，存储left到right的所有值的累加
	p := 0
	// 定义滑动窗口两个指针放到开头
	left, right := 0, 0
	// 定义结果，连续数组的长度
	ans := math.MaxInt64
	for right < n {
		// right指针移动的时候一直累加
		p += nums[right]
		// 判断乘积大于目标值的时候，也就是p>=k的时候，左指针要往右指针靠拢，直到left和right直接数字的乘积小于k为止
		for p >= target {
			// 比较最小的数组路径，数组长度等于right-left+1
			ans = min(ans, right-left+1)
			// left一直向right靠拢，每移动一次，就减去nums[left]
			p -= nums[left]
			// left指针右移
			left++
		}
		// 小优化，如果ans=1，可以直接跳出循环了，1是最佳长度
		if ans == 1 {
			break
		}
		right++
	}
	// 如果循环里面没有比初始答案更小的，说明不存在累加>=target
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
