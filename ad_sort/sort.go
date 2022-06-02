package ad_sort

import "fmt"

// PopSort 冒泡
func PopSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

// SelectSort 选择排序
func SelectSort(arr []int) []int {
	lenArr := len(arr)
	for i := lenArr - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

// InsertionSort 插入排序
func InsertionSort(arr []int) []int {
	arrLen := len(arr)
	// 控制0-i 有序
	for i := 0; i < arrLen; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}

type Uint64Slice []uint64

func QuickSort(numbers Uint64Slice, start, end int) {
	var middle int
	tempStart := start
	tempEnd := end

	if tempStart >= tempEnd {
		return
	}
	pivot := numbers[start]
	for start < end {
		for start < end && numbers[end] > pivot {
			end--
		}
		fmt.Println(end)
		if start < end {
			numbers.swap(start, end)
			start++
		}
		fmt.Println("kkk:=", numbers, start)
		for start < end && numbers[start] < pivot {
			start++
		}
		if start < end {
			numbers.swap(start, end)
			end--
		}
		fmt.Println("eeee:=", numbers, start)
	}
	numbers[start] = pivot
	middle = start

	fmt.Println(middle)

	QuickSort(numbers, tempStart, middle-1)
	QuickSort(numbers, middle+1, tempEnd)

}

func (numbers Uint64Slice) swap(i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}

//MergeSort  归并排序
func MergeSort(arr []int, L int, R int) []int {
	// 0 - 0 位置有序直接返回
	if L == R {
		return arr
	}
	mid := L + ((R - L) >> 1)
	MergeSort(arr, L, mid)
	MergeSort(arr, mid+1, R)
	mergeList(arr, L, mid, R)
	return arr
}

func mergeList(arr []int, L int, mid int, R int) {
	var (
		tempArr []int
		tempL   int
		tempR   int
	)
	tempL = L
	tempR = mid + 1
	for tempL <= mid && tempR <= R {
		if arr[tempL] <= arr[tempR] {
			tempArr = append(tempArr, arr[tempL])
			tempL++
		} else {
			tempArr = append(tempArr, arr[tempR])
			tempR++
		}
	}

	for tempL <= mid {
		tempArr = append(tempArr, arr[tempL])
		tempL++
	}

	for tempR <= R {
		tempArr = append(tempArr, arr[tempR])
		tempR++
	}
	for _, v := range tempArr {
		arr[L] = v
		L++
	}
}

// 求小和 问题
// 例如：[3,5,1,29,2]
// 3 左边 小和 0
// 5 左边 小和 3
// 1 左边 小和 0
// 29 左边 小和 3 + 5 + 1
// 2 左边 小和  1
// 小和 整体 0 + 3 + 0 + 3 + 5 + 1 + 1 =  13
// 思路转换 ：比x大的右边边有几位
// 3 右边比他大的有 5,29 所以  3 *  2
// 5 右边比他大的有  29 所以  5 * 1
// 1 右边比他大的有   29,2 所以 1 * 2
// 29 右边比他的有 0  所以 29 × 0
// 场景分析 ： 这种场景我们可以使用归并排序，
//           1.最左侧优优先出栈 ， 右边其次 就可以在合并的过程中进行对比
//           2.为什么需要排序 ，因为我们期望右侧是有序的 这样我们就可以用很低时间复杂度来求出有几位大于左侧值
//           3.合并过程中右侧值相等的情况下优先进临时数组,如果左侧先进数组则容易丢失数据举例 ：
//              before  left [1,1,1,2,2,3,4,5,]  right [1,2,8,9]
//              错误 : after  temp [1]    left [1,1,2,2,3,4,5,]  [1,2,8,9]  左侧1 还需与 右侧 2,8 ,9 对比
//              正确 ： after  temp [1]    left [1,1,1,2,2,3,4,5,]  [2,8,9]
//
// coding:

func MinSumMerge(arr []int, L int, R int) (sum int) {
	if L == R {
		return
	}
	mid := L + ((R - L) >> 1)
	return MinSumMerge(arr, L, mid) + MinSumMerge(arr, mid+1, R) + MergeSum(arr, L, mid, R)
}

func MergeSum(arr []int, L int, mid, R int) (sum int) {
	var (
		tmpArr []int
		tmpL   int
		tmpR   int
	)
	tmpL = L
	tmpR = mid + 1
	for tmpL <= mid && tmpR <= R {
		if arr[tmpL] < arr[tmpR] {
			sum += (R - tmpR + 1) * arr[tmpL]
			tmpArr = append(tmpArr, arr[tmpL])
			tmpL++
		} else {
			tmpArr = append(tmpArr, arr[tmpR])
			tmpR++
		}
	}
	for tmpL <= mid {
		tmpArr = append(tmpArr, arr[tmpL])
		tmpL++
	}
	for tmpR <= R {
		tmpArr = append(tmpArr, arr[tmpR])
		tmpR++
	}
	for _, v := range tmpArr {
		arr[L] = v
		L++
	}
	return
}
