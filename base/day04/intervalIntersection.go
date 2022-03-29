package day04

/*
给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。

返回这 两个区间列表的交集 。

形式上，闭区间[a, b]（其中a <= b）表示实数x的集合，而a <= x <= b 。

两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。

示例 1：
输入：firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

示例 2：
输入：firstList = [[1,3],[5,9]], secondList = []
输出：[]

示例 3：
输入：firstList = [], secondList = [[4,8],[10,12]]
输出：[]

示例 4：
输入：firstList = [[1,7]], secondList = [[3,10]]
输出：[[3,7]]

*/

func IntervalIntersection(firstList [][]int, secondList [][]int) (ans [][]int) {
	i, j, m, n := 0, 0, len(firstList), len(secondList)
	// 如果其中一个为空，直接返回[][]int{}
	if m == 0 || n == 0 {
		return ans
	}
	// 因为数组都是排好序的升序数组，所以可以从左到右，分别在两个数组里直接拿一个区间出来比较
	for i < m && j < n {
		// 分别从两边拿出区间出来对比
		// 第一个参数取最大的
		cur0 := getMax(firstList[i][0], secondList[j][0])
		// 第二个参数取最小的
		cur1 := getMin(firstList[i][1], secondList[j][1])

		// 如果第一个参数的最大者<=第二个参数最小者，说明这两个参数就是交集
		if cur0 <= cur1 {
			ans = append(ans, []int{cur0, cur1})
		}

		// 如果两个相同，说明上面已经把这个交集都处理加进去了，直接同时移动到右边，就少遍历一次
		// 例如 [0,2]
		//     [1,2] 交集是[1,2]
		// 因为题目说明了每个数组的每个区间列表都是成对 不相交的
		// 所以不会出现[0,2][2,3]这种
		//
		if firstList[i][1] == secondList[j][1] {
			i++
			j++
			continue
		}
		// 第二个参数最小者的那个数组的区间，移动到下一个区间
		if firstList[i][1] < secondList[j][1] {
			i++
			continue
		}
		j++

	}
	return ans
}
func getMax(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}
func getMin(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
