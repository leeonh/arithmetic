package day01

/**
搜索旋转排序数组
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。



示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1
*/
func Search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// 二分法
	start, end := 0, len(nums)-1
	for start <= end {
		// 获取中间指针
		m := int(uint(start+end) >> 1)
		// 如果中间值等于目标值
		if nums[m] == target {
			return m
		}
		// 如果前半部分是有序的
		if nums[0] <= nums[m] {
			// 如果目标值在这个有序的前半部分里面
			if nums[0] <= target && target < nums[m] {
				// 右指针移动到中间指针的前面
				end = m - 1
			} else {
				//说明不在前半部分里面，将左指针移动到中间指针的后面
				start = m + 1
			}
		} else {
			// 说明前半部分可能是无序的，当然后半部分也可能是无序的
			// 如果目标值在中间指针和右指针之间
			if nums[m] < target && target <= nums[n-1] {
				start = m + 1
			} else {
				//如果不在后半部分
				end = m - 1
			}
		}
	}
	return -1
}
