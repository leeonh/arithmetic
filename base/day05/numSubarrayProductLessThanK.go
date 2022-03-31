package day05

/*
乘积小于K的子数组

给定一个正整数数组nums和整数 k 。

请找出该数组内乘积小于k的连续的子数组的个数。



示例 1:

输入: nums = [10,5,2,6], k = 100
输出: 8
解释: 8个乘积小于100的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
示例 2:

输入: nums = [1,2,3], k = 0
输出: 0

*/

// 题目的关键词，连续的子数组
// 说明例如[10,5,2,6]，不会出现[10,2][10,6]这种跳着来的组合，那这样就可以用滑动窗口来做了
func numSubarrayProductLessThanK(nums []int, k int) int {
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
		// 关键地方：
		// right - left + 1 相当于 left和right之间所有的连续子数组组合 例如 [1,2,3,4],k=100, left=0，right=3，相当于总共有[4],[4, 3],[4, 3, 2],[4, 3, 2, 1]
		// 为什么是从右边开始往左边数连续数组？再举个例子，如果从左往右数连续数组的话，那上面这个[1,2,3,4] 就会变成 [1], [1,2], [1,2,3], [1,2,3,4]
		// 那[1,2,3,4]的上一个循环，也就是[1,2,3]，也满足乘积小于k，那么从左往右数的话，[1], [1,2], [1,2,3]，就和上面的重复了
		// 如果是从右往左数
		// [1,2,3,4]，就是[4],[4, 3],[4, 3, 2],[4, 3, 2, 1]，总共4个
		// [1,2,3]，就是[3],[3,2],[3,2,1]，总共3个
		// [1,2]，就是[2],[2,1]，总共2个
		// [1]，就是[1]，总共1个
		// 连续子数组的总数1+2+3+4=10个
		ans += right - left + 1
		right++
	}
	return ans
}