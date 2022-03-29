package day11

func LetterCasePermutation1(s string) []string {
	// 将字符串转换成字节
	chs := []byte(s)
	// 获取字符串长度
	l := len(s)
	// 存储结果的二维数组
	var result []string
	// 存储单一结果的一维数组
	var path []byte
	// 定义回溯方法，参数为当前递归的深度，也相当于数组的下标
	var backtracking func(int)
	// 递归得到结果
	backtracking = func(depth int) {
		// 递归终止条件，当path一维数组的大小等于s的大小时候，将path加入result，并返回result
		if len(path) == l {
			temp := string(path)
			result = append(result, temp)
			return
		}
		// 将字符放进path数组
		path = append(path, chs[depth])
		// 递归将后面的字符组合起来放进结果数组
		backtracking(depth + 1)
		// 每一层递归完后，进行回溯，也就是恢复到上一层的状态
		path = path[:len(path)-1]
		// 判断当前字符是大写还是小写
		if chs[depth] >= 'a' && chs[depth] <= 'z' {
			chs[depth] -= 32
			path = append(path, chs[depth])
			backtracking(depth + 1)
			// 每一层递归完后，进行回溯，也就是恢复到上一层的状态
			path = path[:len(path)-1]
			chs[depth] += 32
		}
		// 判断当前字符是大写还是小写
		if chs[depth] >= 'A' && chs[depth] <= 'Z' {
			chs[depth] += 32
			path = append(path, chs[depth])
			backtracking(depth + 1)
			// 每一层递归完后，进行回溯，也就是恢复到上一层的状态
			path = path[:len(path)-1]
			chs[depth] -= 32
		}

	}
	backtracking(0)
	return result
}
func LetterCasePermutation(s string) []string {
	// 存储结果的二维数组
	var result []string
	// 将字符串转换成字节
	chs := []byte(s)
	dfs([]byte(s), 0, len(chs), &result)
	return result
}

/**
* @Param chs 字符串转换成的byte数组
* @Param index 当前递归来到了第几个字符
* @Param n 字符串数组长度
* @Param result 结果字符串数组
**/
func dfs(chs []byte, index int, n int, result *[]string) {
	// 当当前递归到的字符已经是最后一个字符的时候
	if index == n {
		*result = append(*result, string(chs))
		return
	}
	// 如果当前字符不是数字
	if !(chs[index] >= '0' && chs[index] <= '9') {
		// 将字符转换成大写或者小写，然后index + 1，指针继续递归向后移动，直到字符串末尾，然后存入结果在一层层返回
		// 例如字母'a'，Ascii码是97，二进制是0011 00001，32的二进制就是0010 0000，那么和'a'的异或结果为1001 1101，再转换成10进制是65，对应的Ascii码为'A'
		chs[index] ^= 32
		dfs(chs, index+1, n, result)
		// 将字符转换成大写或者小写
		chs[index] ^= 32
		dfs(chs, index+1, n, result)
	} else {
		dfs(chs, index+1, n, result)
	}
}
