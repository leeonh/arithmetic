package day03

import "sort"

func ThreeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	// 先将数组升序排序
	sort.Ints(nums)
	ans := make([][]int, 0)
	// 循环数组
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}
		// 因为答案是要不一样的排列组合，所以排序后，如果有重复的数字都是连续的，所以这里就可以判断，当i的值等于i-1的时候，说明这个在上一轮循环已经用过这个组合了，进行去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针，一个放在i+1,一个放到数组结尾，如果重合则结束循环
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			// 如果相加的结果等于0，直接插入答案数组
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 因为排了序，所以可以直接判断left是否等于left+1，等于说明下次组合的时候肯定会重复使用同一种数字
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 同样，right是否等于right-1，等于说明下次组合的会重复
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				//如果比0小，说明left需要往右边移
				left++
			} else {
				//如果比0大，说明right需要往做边移
				right--
			}
		}
	}
	return ans
}
