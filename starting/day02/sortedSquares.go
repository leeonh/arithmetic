package day02

func SortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; {
		if nums[i]*nums[i] < nums[j]*nums[j] {
			res[pos] = nums[j] * nums[j]
			j--
		} else {
			res[pos] = nums[i] * nums[i]
			i++
		}
		pos--
	}
	return res
}
