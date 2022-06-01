package recursion

import (
	"fmt"
)

// Master公式：分析递归函数的时间复杂度，要求子问题规模一致
// T(N) = a * T(N/b) + O(N^d)（其中a、b、d都是常数）的递归函数，可以直接通过Master公式来确定
// 时间复杂度
// 1）如果log(b)(a) < d，时间复杂度为O(N^d)
// 2）如果log(b)(a) > d，时间复杂度为O(N^log(b,a))
// 3）如果log(b)(a) == d，时间复杂度为O(N^d * logN)

// ProcessGetArrMaxValue 使用递归一个数组中最大值 例如：[3,4,5,89,821,2,122,891,921] 最大值 是 921
func ProcessGetArrMaxValue(arr []uint64, L int, R int) (max uint64) {
	if L == R {
		max = arr[L]
		return
	}
	mid := L + ((R - L) >> 1)
	fmt.Println("L,R :", L, R)
	fmt.Println("L,mid:", L, mid)
	leftMax := ProcessGetArrMaxValue(arr, L, mid)
	fmt.Println("mid+1,R:", mid+1, R)
	rightMax := ProcessGetArrMaxValue(arr, mid+1, R)
	max = compareMax(leftMax, rightMax)
	return
}

func compareMax(a, b uint64) (value uint64) {
	if a > b || a == b {
		value = a
		return
	}
	return b
}
