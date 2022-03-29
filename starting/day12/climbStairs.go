package day12

/**
假设你正在爬楼梯。需要 n阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？


示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
*/

func climbStairs(n int) int {
	// 通过枚举，先找出题目的规律
	// 第一阶1种，第二阶2种，第三阶3种，第四阶5种，第五阶8种
	// 得出结论，每一阶的方法种数，等于前两阶的总数相加 f(n) = fn-1 + fn
	// 可以将第零阶也看做一种方法，f2 = f0 + f1
	// 分别用3个指针循环迭代到后面
	zero := 1
	first := 1
	var second int
	for i := 1; i < n; i++ {
		// second指针的值等于前两个指针的值相加
		second = zero + first
		// zero指针移动到first指针的位置
		zero = first
		// first指针移动到second指针的位置
		first = second
	}
	return first
}
