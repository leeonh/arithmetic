package day14

/**
颠倒给定的 32 位无符号整数的二进制位。

提示：

请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在 示例 2中，输入表示有符号整数 -3，输出表示有符号整数 -1073741825。


示例 1：

输入：n = 00000010100101000001111010011100
输出：964176192 (00111001011110000010100101000000)
解释：输入的二进制串 00000010100101000001111010011100 表示无符号整数 43261596，
     因此返回 964176192，其二进制表示形式为 00111001011110000010100101000000。
示例 2：

输入：n = 11111111111111111111111111111101
输出：3221225471 (10111111111111111111111111111111)
解释：输入的二进制串 11111111111111111111111111111101 表示无符号整数 4294967293，
    因此返回 3221225471 其二进制表示形式为 10111111111111111111111111111111 。

*/
func reverseBits1(num uint32) uint32 {
	// 将num最末尾的编码拿出来和res拼接
	// 拿出来的方法用到和1做与运算(&)，例如1101和0001做与运算，就能拿到最后一位是1
	// 然后拼接到res，因为res左移后都是0，所以和上面num拿出来的最后一位做或运算(|)，都为0就是零，其中一个为1就是1，就可以实现拼接了
	//
	res := uint32(0)
	for i := 0; i < 32; i++ {
		res = res<<1 | (num & 1)
		num >>= 1
	}
	return res
}

func reverseBits(num uint32) uint32 {
	// 使用分而治之的方法，将32位切割成两个16然后交换位置，然后再分为4个8分别交换位置，再分为8个4，再分为16个2，再分为32个1

	// 0000000000000000 1111111111111111  		1111111111111101 0000000000000000
	// 		左边部分就是右移16位  							右边部分就是左移16位
	// 或运算(|)，两个都为0才为0，其他都是1，正好相当于将32位二进制，分割成两份做交换
	// 运算后为：1111111111111101 1111111111111111
	num = (num >> 16) | (num << 16)
	// 接下来就是分割成4个8来交换了
	// 	   11111111 11111101 11111111 11111111
	// 分别和
	//	11111111 00000000 11111111 00000000
	//  00000000 11111111 00000000 11111111
	// 做与(&)运算 两个都为1的时候才为1，其他都为0
	// 结果分别为：
	//	11111111 00000000 11111111 00000000
	//	00000000 11111101 00000000 11111111
	//  两个结果上面的右移8位，下面的左移8位然后再做或运算
	//  00000000 11111111 00000000 11111111 | 11111101 00000000 11111111 00000000
	// 	11111101 11111111 11111111 11111111
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)

	return num
}
