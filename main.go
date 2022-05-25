package main

import (
	"fmt"
	"go_adt/ad_sort"
)

func main() {
	arr := []int{2, 3, 43, 4, 801, 8, 101, 9, 1, 12, 1, 1, 2, 121, 23}
	fmt.Println(ad_sort.SelectSort(arr))
}
