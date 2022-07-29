package SignIn

/*
1184. 公交站间的距离

环形公交路线上有n个站，按次序从0到n - 1进行编号。我们已知每一对相邻公交站之间的距离，distance[i]表示编号为i的车站和编号为(i + 1) % n的车站之间的距离。

环线上的公交车都可以按顺时针和逆时针的方向行驶。

返回乘客从出发点start到目的地destination之间的最短距离。

*/
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	// 总共有两条路线，一条是顺时针走的，一条是逆时针走
	// 但是有可能出现起始点大于终点位置的情况，这种情况其实可以当做将两个点交换一下，效果是一样的，因为只会有两条路线，但是在这样切换了之后，方便我们遍历
	if start > destination {
		start, destination = destination, start
	}

	sum1 := 0
	sum2 := 0
	for i, d := range distance {
		// 如果指针在start和destination范围内，说明要累计顺时针走的路径长度
		if start <= i && i < destination {
			sum1 += d
		} else {
			// 如果在start和destination范围外，说明是在走逆时针的路，也就是start到0，再从0到destination的路径长度
			sum2 += d
		}
	}

	return getMin(sum1, sum2)

}
func getMin(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
