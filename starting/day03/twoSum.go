package day03

func TwoSum(numbers []int, target int) []int {
	var n = len(numbers)
	for i, j := 0, n-1; i < j; {
		m := (i + j) >> 1
		if numbers[i]+numbers[m] > target {
			j = m - 1
		} else if numbers[m]+numbers[j] < target {
			i = m + 1
		} else if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{-1, -1}

}
