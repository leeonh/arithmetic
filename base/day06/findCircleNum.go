package day06

/*
省份数量

有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。

返回矩阵中 省份 的数量。



示例 1：


输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
输出：2
示例 2：


输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
输出：3

提示：

1 <= n <= 200
n == isConnected.length
n == isConnected[i].length
isConnected[i][j] 为 1 或 0
isConnected[i][i] == 1
isConnected[i][j] == isConnected[j][i]

*/

func findCircleNum(isConnected [][]int) int {
	// 定义省份数量
	ans := 0
	// 这道睇不一样的地方就是，数组里面的值并不是城市，而是存放着两座城市直接是否是直连的判断，所以岛屿数量那种方法不适用
	// 根据题目的意思可得，坐标(i,j)，就是代表着2个城市(i,j)，如果i==j，说明就是城市自己
	// 数组里面每一个值，都相当于城市之间是否相连的判断
	// 例如：[1,1,0]
	//		[1,1,0]
	//		[0,0,1]
	// 第一个数字1的坐标是(0,0)，说明0和0这两个城市是直连的，其实就是0城市自己
	// 第二个数字1(0,1)，说明0城市和1城市直连
	// 题目的提示里面：isConnected[i][j] == isConnected[j][i]，说明这个二维数组，列和排是相同的
	// 由此可得，上面的例子，总共就3个城市拿出来作比较，分别是0城市、1城市和2城市，也就是数组的长度len(isConnected)
	n, m := len(isConnected), len(isConnected[0])

	// 定义一个boolean数组，用来判断城市是否访问过
	visit := make([]bool, len(isConnected))
	// 创建一个队列
	for i := 0; i < n; i++ {
		// 如果当前城市i还没访问过
		if !visit[i] {
			// 增加省份
			ans++
			// 将当前城市加入队列
			queue := []int{i}
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]
				// 将当前城市标记为已经访问过了
				visit[from] = true
				// 循环这一层，处理这一层所有城市的直连关系判断
				// 例如[1,1,0]
				// 相当于当前0城市 和 0，1，2这三个城市是否存在直连的关系的判断
				for j := 0; j < m; j++ {
					// 如果这两个是直连的，并且城市还没访问过，则加入队列
					// 相当于将直连的城市全部标记为已经访问
					// 因为这些城市都是互相直连或者间接连的，联合起来组成的只能算一个省
					if isConnected[from][j] == 1 && !visit[j] {
						queue = append(queue, j)
						// 整个过程复现一下：
						// 例如第一轮将0加入队列，那么就开始循环第0行，也就是[1,1,0]
						// 然后省份+1
						// 坐标(0,0)=1，第一次循环肯定是从0开始的，所以0这个城市已经访问过了，不加入队列
						// 坐标(0,1)=1，1城市还没访问过，将1加入队列
						// 队列循环来到第二轮，上一轮最后将1加入了队列，所以这一轮循环第1行，也就是[1,1,0]
						// 首先 visit[from] = true 将1城市标记为已访问
						// 坐标(1,0)=1，上一轮0城市已经标记为访问过了，所以不加入队列
						// 坐标(1,1)=1，1城市已经访问过了，当前轮没有新增城市加入队列，所以队列循环结束

						// 最外层的循环来到最后一行，i=2，2城市还没访问过，加入队列，同时省份+1
						// 开始队列循环，来到第2行，也就是[0,0,1]，并将2城市标记为已访问
						// (2,0)=0，不相连
						// (2,1)=0，不相连
						// (2,2)=1，相连，但是2城市已经访问过了，循环结束
					}
				}
			}
		}
	}
	return ans
}