package ad_sort

import (
	"fmt"
	"testing"
)

func TestPopSort(t *testing.T) {
	arr := []int{
		4, 5, 89, 1, 2, 34, 56, 98, 7,
	}
	res := PopSort(arr)
	fmt.Println(res)
}

func TestSelectionSort(t *testing.T) {
	arr := []int{
		4, 5, 89, 1, 2, 34, 56, 98, 7,
	}
	res := SelectSort(arr)
	fmt.Println(res)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{
		4, 5, 89, 1, 2, 34, 56, 98, 7,
	}
	res := InsertionSort(arr)
	fmt.Println(res)
}

//               0 - 4
//       0 - 2          3 - 4
//      [4,5,89]        [1,2]
//   0 - 1   2 - 2     3 - 3   4- 4
//  [4,5]
//0 - 0  1 - 1
//
func TestMergeSort(t *testing.T) {
	arr := []int{
		4, 5, 89, 1, 2, 34, 56, 98, 7,
	}
	lenArr := len(arr)
	res := MergeSort(arr, 0, lenArr-1)
	fmt.Println(res)
}
