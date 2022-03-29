package day02

func Rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(arr []int) {
	n := len(arr)
	for i, j := 0, n-1; i <= j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
