package int_sum

/**
 整数求和  不使用+-
 逻辑 ： ^ 求和  & 计算进位
 例如 ： ^ 求和
    1 + 0 = 1
    0 + 1 = 1
    1 + 1 = 0  // 需要补位  10
    0 + 0 = 0
**/

func Sum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
