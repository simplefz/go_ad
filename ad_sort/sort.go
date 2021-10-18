package ad_sort

import "fmt"

// PopSort 冒泡
func PopSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > 0; j-- {
			if arr[j] > arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
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
