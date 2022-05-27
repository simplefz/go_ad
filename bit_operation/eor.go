package bit_operation

// ^(xor)异或运算定理：
//         相同为假,不同为真
//  二进制不进位加法
//        理解：
//              10101 ^ 10001 = 100
//        公式转换：
//           1.任何值^0 都是他自己：  a^0 = a    任何值^他本身都等于0  a ^ a = 0
//           2.交换顺序^不影响最终的^结果 例如 ：   a^b^a = a^(b^a)

// SwapValue 题目： 不开辟任何内存空间交换两个变量的值 例如： a = 10 b = 18
func SwapValue(a, b int) (swapA, swapB int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

// XorArrOddNumberOne 题目给定一个数组，数组中有一种不同的数出现次数是奇数，其他都是偶数，求这个奇数
//   例如：[111,111,333,333,555,555,555]
func XorArrOddNumberOne(arr []int) (oddNumber int) {
	eor := 0
	for _, v := range arr {
		eor ^= v
	}
	oddNumber = eor
	return
}

// XorArrOddNumberTwo 题目给定一个数组，数组中有两种不同的数出现次数是奇数，其他都是偶数，求这两个奇数
//   例如：[111,111,333,333,555,555,555，777]
func XorArrOddNumberTwo(arr []uint32) (one uint32, two uint32) {
	var eor uint32
	var eorOther uint32
	eor = 0
	for _, v := range arr {
		eor ^= v
	}
	// eor  = 奇数a ^ 奇数 b
	// eor != 0
	// 设 eor 其中的某一位为 1 这里取最右一位
	rightBit := eor & (^eor + 1)
	eorOther = 0
	for _, v := range arr {
		if v&rightBit == 0 {
			eorOther ^= v
		}
	}
	one = eorOther
	two = eor ^ eorOther
	return
}
