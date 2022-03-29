package day02

func findPeakElement(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	// right = n-1 的原因，因为下面用到了mid+1的判断，会越界
	// 因为有多个峰值，所以要判断两个情况
	// 1. mid > mid+1 的值，说明峰值在左边
	// 2. mid <= mid+1 的值，说明峰值在右边

	// 寻找峰值，相当于是寻找一个边界，因此应该要选择用左闭右开 [left,right)
	// 但是为什么right = n-1呢？
	// 因为判断的时候，已经提前把mid+1也访问了
	for left < right {
		mid := int(uint(left+right) >> 1)
		if nums[mid] < nums[mid+1] {
			// 如果mid < mid + 1的值，说明峰值肯定在右边，向右缩减搜索范围
			left = mid + 1
		} else {
			// 如果mid >= mid + 1的值，说明峰值肯定在左边，向左缩减搜索范围
			right = mid
		}

	}
	return left
}
