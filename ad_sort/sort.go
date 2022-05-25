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
