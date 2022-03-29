package day11

/**
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]

*/
func Permute(nums []int) [][]int {
	l := len(nums)
	// 存放结果的二维数组
	var result [][]int
	// 存储单一结果的一维数组
	var path []int
	// boolen数组，判断元素是否再上一层递归使用了
	used := make([]bool, l)
	var getPermute func([]int, []bool)
	// 递归函数, start，当前数组下标，单一结果数组，Boolean数组
	getPermute = func(path []int, used []bool) {
		// 递归结束条件path一维数组的大小等于nums数组大小
		if len(path) == l {
			temp := make([]int, l)
			copy(temp, path)
			result = append(result, temp)
		}
		// 循环递归，每一层循环，判断是否再上一层是否已经使用过了
		for i := 0; i < l; i++ {
			// 如果这个数字还没有使用过
			if !used[i] {
				path = append(path, nums[i])
				// 标记使用过了
				used[i] = true
				// 继续递归组合其他没有使用过的数字
				getPermute(path, used)
				// 回溯，当前递归相当于栈，后进后出，取消元素已经使用的标识
				used[i] = false
				// 回退一个结果元素
				path = path[:len(path)-1]
			}
		}
	}
	getPermute(path, used)
	return result

}
