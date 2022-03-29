package day01

/**
编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值。该矩阵具有如下特性：
每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。


示例 1：

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true
示例 2：

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false

*/

func searchMatrix(matrix [][]int, target int) bool {
	// 每次循环对当前行数组进行一次判断，判断当前行最后一位是否大于target，是的话，二分查找当前行
	// 否则，继续循环到下一层，重复上面的判断操作
	n := len(matrix)
	for i := 0; i < n; i++ {
		nums := matrix[i]
		// 如果目标值在当前层
		if target <= nums[len(nums)-1] {
			left, right := 0, len(nums)-1
			for left <= right {
				mid := int(uint(left+right) >> 1)
				if nums[mid] == target {
					return true
				} else if nums[mid] < target {
					// 如果mid的值比target小，说明mid要继续向右变大，才能可能找到目标值,也就是target在以mid分割的右边区间里面，向右缩小搜索区间
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}
	return false
}
