package day07

// 创建移动策略 x,y对应上 右 下 左
var (
	//x轴
	dx = []int{0, 1, 0, -1}
	//y轴
	dy = []int{1, 0, -1, 0}
)

//广度优先搜索
func FloodFill1(image [][]int, sr int, sc int, newColor int) [][]int {
	currColor := image[sr][sc]
	// 如果目标位置已经是目标颜色
	if currColor == newColor {
		return image
	}

	// 使用二维数组模拟创建队列，每次移动，将需要改色的坐标放入队列
	queue := [][]int{}
	queue = append(queue, []int{sr, sc})
	image[sr][sc] = newColor
	n, m := len(image), len(image[0])

	//一直循环到队列末尾
	for i := 0; i < len(queue); i++ {
		cell := queue[i]
		// 循环移动策略，四个方向
		for j := 0; j < 4; j++ {
			// 存储移动后的坐标
			mx, my := cell[0]+dx[j], cell[1]+dy[j]
			if mx >= 0 && mx < n && my >= 0 && my < m && image[mx][my] == currColor {
				// 如果移动后，坐标没有越界，则将坐标添加到队列末尾
				queue = append(queue, []int{mx, my})
				image[mx][my] = newColor
			}
		}
	}
	return image
}

//深度优先
func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	currColor := image[sr][sc]
	if currColor != newColor {
		dfs(image, sr, sc, currColor, newColor)
	}
	return image
}

func dfs(image [][]int, x, y, color, newColor int) {
	// 如果是需要切换的颜色
	if image[x][y] == color {
		image[x][y] = newColor
		// 按照移动策略移动
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx >= 0 && mx < len(image) && my >= 0 && my < len(image[0]) {
				dfs(image, mx, my, color, newColor)
			}
		}
	}

}
