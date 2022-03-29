package day02

import "math"

func findMin(nums []int) int {
	// 不管旋转多少次，数组都会在某一个下标处分割为两个升序数组
	n := len(nums)

	left, right := 0, n
	min := math.MaxInt64

	for left < right {
		mid := int(uint(left+right) >> 1)
		// 如果mid往左的区间是升序的，那直接取left，就是最小值
		if nums[0] <= nums[mid] {
			min = getMin(min, nums[left])
			// 然后去右区间看看，有没有更小的
			left = mid + 1
		} else {
			// 如果左区间是无序的，说明右区间肯定是有序的，那么此时mid的值，必定是最小值
			min = getMin(min, nums[mid])
			// 然后去左区间看看，有没有更小的
			right = mid
		}
	}
	return min
}
func getMin(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
