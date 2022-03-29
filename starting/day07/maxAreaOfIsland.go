package day07

func MaxAreaOfIsland(grid [][]int) int {
	var area = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area = max(areaDfs(grid, i, j), area)
			}

		}
	}
	return area
}
func MaxAreaOfIsland1(grid [][]int) int {
	var area = 0
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				area = max(bfs(grid, i, j, n, m), area)
			}

		}
	}

	return area
}

func bfs(grid [][]int, x, y, n, m int) int {
	grid[x][y] = 2
	var area = 1
	queue := [][]int{}
	queue = append(queue, []int{x, y})
	for i := 0; i < len(queue); i++ {
		cell := queue[i]
		// 循环移动策略，四个方向
		for j := 0; j < 4; j++ {
			// 存储移动后的坐标
			mx, my := cell[0]+dx[j], cell[1]+dy[j]
			if mx >= 0 && mx < n && my >= 0 && my < m && grid[mx][my] == 1 {
				// 如果移动后，坐标没有越界，则将坐标添加到队列末尾
				queue = append(queue, []int{mx, my})
				grid[mx][my] = 2
				area++
			}
		}
	}
	return area
}

// 深度搜索
func areaDfs(grid [][]int, x, y int) int {
	var m = 0
	// 如果是陆地
	if grid[x][y] == 1 {
		grid[x][y] = 2
		m = 1
		// 按照移动策略移动
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			//如果移动后还是陆地，	并且没有走过登陆过，则继续递归深度搜索
			if mx >= 0 && mx < len(grid) && my >= 0 && my < len(grid[0]) && grid[mx][my] == 1 {
				m += areaDfs(grid, mx, my)
			}

		}
	}
	return m

}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
