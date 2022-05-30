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
