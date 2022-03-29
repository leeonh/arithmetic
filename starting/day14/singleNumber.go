package day14

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例2:

输入: [4,1,2,1,2]
输出: 4

*/
func singleNumber(nums []int) int {
	// 位进制运算，用异或
	// 首先异或有一个特性，就是不同的时候为1，相同的时候为0
	// 这个特性放到到这道题里面，当出现相同数字的时候，二进制就是直接抵消掉，直接变成0，直到最后剩下一个只出现过一次的数
	// 就拿4，1，2，1，2这个举例
	// 4的二进制：100，1的二进制：001
	// 异或(^)的结果为 101，也就是5，接着和2的二进制：010做异或
	// 结果为111，也就是7，接着和1的二进制：001做异或，此时最右的1就抵消了，变成了110，也就是6
	// 接着和2的二进制010，异或后，100，再次抵消变为4了

	// 我们用一个0去异或完整的接收第一个数
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}
