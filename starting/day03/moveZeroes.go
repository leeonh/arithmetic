package day03

func MoveZeroes(nums []int) {
	for i, f := 0, 0; f < len(nums); f++ {
		if nums[f] != 0 {
			nums[i], nums[f] = nums[f], nums[i]
			i++
		}
	}
}
