package ad_sort

import "fmt"

// PopSort 冒泡
func PopSort(arr []int) []int {
	n := len(arr)
	fmt.Println("未排序", arr, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			fmt.Println("j,j+", j, j+1)
			if arr[j] > arr[j+1] {
				fmt.Println("对比值：", arr[j], arr[j+1], arr)
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
		fmt.Println("第", i, "次", arr)
	}
	return arr
}

// SelectSort 选择排序
func SelectSort(arr []int) []int {
	lenArr := len(arr)
	fmt.Println("未排序", arr, lenArr)
	for i := lenArr - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Println("i,j", i, j)
			fmt.Println("对比值：", arr[i], arr[j], arr)
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		fmt.Println("第", i, "次", arr)
	}
	return arr
}

// InsertionSort 插入排序
func InsertionSort(arr []int) []int {
	arrLen := len(arr)
	// 控制0-i 有序
	for i := 0; i < arrLen; i++ {
		fmt.Println(arr[i])
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
		fmt.Println(arr)
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
	fmt.Println(L, mid)
	MergeSort(arr, L, mid)
	fmt.Println(mid+1, R)
	MergeSort(arr, mid+1, R)
	fmt.Println(arr, L, mid, R)
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
	fmt.Println(tempArr)
	for _, v := range tempArr {
		arr[L] = v
		L++
	}
}
