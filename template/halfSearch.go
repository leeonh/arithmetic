package template

/**
二分搜索模板
关键词：二分法 二分 二分搜索 左边界 右边界

*/

/**
搜索边界的二分法
*/
func halfSearchBoundary(nums []int, target int) int {
	// 这种适合搜索左右边界
	left, right := 0, len(nums)
	for left < right {
		mid := int(uint(left+right) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

/**
搜索准确值的二分法
*/
func halfSearchValue(nums []int, target int) int {
	// 这种适合找到特定的目标值
	// 区别就在于
	// right=n-1，然后循环结束的条件变成了left<=right，右边界缩减的时候，right=mid-1
	left, right := 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
