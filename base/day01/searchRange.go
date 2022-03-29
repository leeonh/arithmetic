package day01

/**
34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？


示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]
*/
func SearchRange(nums []int, target int) []int {
	// 数组长度
	n := len(nums)
	// 二分查找有几个常见的场景，例如寻找一个数，寻找左边界，寻找右边界
	// 此题的答案刚好就是，一个寻找左边界，一个寻找右边界
	// 为什么循环的结束条件是left < right而不是left<=right?这个其实没有任何区别，只是习惯问题，用哪个，最终只是细节方面分别做一些不同的调整
	// 因为 right=len(nums)，right这个索引是越界了的，所以这是一个左闭右开的区间[left,right)，当rigth=left的时候，正好区间[left,right)相当于[left,left)的结果为空，结束循环
	// 为什么此题要用左闭右开区间[left,right)而不用左闭右闭[left,right]来做呢？
	// 根据题目的定义，0 <= nums.length，也就是说数组是会为空的，如果用闭区间，那么你的
	// 首先找出第一个答案的位置，相当于就是寻找目标值的区间范围的左边界[left,right)，左边闭区间，右边开区间
	// 例如 [5,7,7,7,7,8,9,10], target=7，那就相当于找值为7的区间的左边界，也就是索引1

	// 搜索的思路：
	// 当nums[mid]<target的时候，left=mid+1，向右缩减搜索区间，在[mid + 1,right)继续搜索
	// 关键步骤，当nums[mid]==target的时候，不要立即返回mid，而是缩小搜索区间的上界right，在[left,mid)继续搜索
	// 当nums[mid]>target的时候，缩小搜索区间的上界right，在[left,mid)继续搜索
	// 相当于不停的向左搜索，直到将索引锁定到左边界为止，也就是当right=mid，并且mid刚好计算出来也等于left => [left,left)结束循环
	// left或者right就是左边界，你可以随便拿其中一个作为答案，我这里用了left
	// 那为什么是r=mid 而不是r = mid - 1 呢
	// 因为这是一个左闭右开的区间[left,rigth)，所以当搜索到mid的时候，就会将区间分割为两个，分别是[left,mid) 或者 [mid + 1,right)，其实[left,mid-1] 和 [left,mid)区间内的值是一样的
	left := searchR(n, func(mid int) bool {
		return nums[mid] >= target
	})

	// 当然还得判断，这次循环结束后，这个指针的位置的值，是否就是答案，不是的话直接返回(-1,-1)
	// 为什么要判断left == n呢？这个判断的意思就是，判断左边界的索引是否越界了，例如[],target=1，这种情况搜查左边界的时候，right=len(nums)=0,初始化left=0
	if left == n || nums[left] != target {
		return []int{-1, -1}
	}

	// 查找右边界，和左边界相似，但是关键步骤相反
	// 核心的地方就是nums[mid]<=target 相当于 nums[mid]>target的时候，left=mid+1
	// 同样的道理，nums[mid]<=target的时候，不要立即返回mid，而是不断的缩小下界left，在[mid+1,right)区间继续搜索，直到将索引锁定为右边界位置
	// 但是这里的情况不一样
	// 例如[7,7,10,11,12], target=7
	// 第一次mid=2，nums[2] > 7，right=mid,区间变成[0,2)，也就是[7,7,10)
	// 第二次mid=1，nums[1] >= 7，left = mid+1，区间变成[10),此时left=right，退出循环
	// 也就是说，每次更新left都是mid+1，所以循环结束后，nums[left]是一定不等于target了，而nums[left-1]可能等于target，为什么这里我能断定nums[left-1]一定会等于target呢？
	// 这就是因为，上面已经把左边界搜索出来了，如果没搜索出来，就不会进入到这里搜索右边界，也就是说，当前目标值target是绝对存在的，因此右边界的取值至少都会等于左边界的索引，right>=left
	right := searchR(n, func(mid int) bool {
		return nums[mid] > target
	})
	return []int{left, right - 1}
}
func searchR(n int, f func(int) bool) int {
	left, right := 0, n
	for left < right {
		mid := int(uint(left+right) >> 1)
		if f(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
	// 这里再贴一个left<=right的，区别就在于
	// right=n-1，然后循环结束的条件变成了left<=right，右边界缩减的时候，right=mid-1
	//left, right := 0, n-1
	//for left <= right {
	//	mid := int(uint(left+right) >> 1)
	//	if f(mid) {
	//		right = mid - 1
	//	} else {
	//		left = mid + 1
	//	}
	//}
	//return left

}
