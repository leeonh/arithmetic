package day09

var (
	// 移动的策略，x轴，y轴， 下面对应，上，下，左，右
	move = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
)

func UpdateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	// 使用二维数组模拟创建队列
	queue := [][]int{}
	// 循环矩阵，将所有的0添加进队列
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果当前值为0，则添加入队列
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				// 标记已经访问过了
				mat[i][j] = -1
			}
		}
	}
	// 循环模拟队列弹出，如果队列不为空，就循环弹出(切片)
	for len(queue) > 0 {
		// 将队列里面的坐标拿出来
		cell := queue[0]
		queue = queue[1:]
		i, j := cell[0], cell[1]
		// 循环移动策略，分别移动上 下 左 右
		for d := 0; d < 4; d++ {
			// x轴移动后的坐标位置，因为0距离最近的0就是0，所以可以作为初始的距离，开始累加起来，每次
			x := i + move[d][0]
			// y轴移动后的坐标位置
			y := j + move[d][1]
			// 判断移动后是否超界，是否已经访问过
			if x >= 0 && x < m && y >= 0 && y < n && mat[x][y] == -1 {
				// 因为只移动一个次，距离循环拿出来的坐标点1个单位，所以加一
				mat[x][y] = mat[i][j] + 1
				// 将移动后的坐标，加入队列
				queue = append(queue, []int{x, y})
			}
		}
	}
	return mat
}
