package day04

import "math"

/*
给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。

找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

示例 1：
说明：你不能倾斜容器。
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
示例 2：

输入：height = [1,1]
输出：1
*/
func MaxArea(height []int) int {
	n := len(height)
	// 初始化一个最小的容量答案
	ans := math.MinInt64
	// 双指针一个开头一个结尾，往中间扫描
	i, j := 0, n-1
	for i < j {
		// 每次循环计算容量，容量体积 = 横坐标的差 * height最矮的那个
		ans = max(ans, (j-i)*min(height[i], height[j]))
		// height小的那边继续向另一边移动
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
