package main

import (
	"fmt"
	"go_adt/code_string"
)

func main() {
	/*arr := []int{9,10,210,21,1,33,12,5521,221029291,11}
	fmt.Println(ad_sort.PopSort(arr))
	fmt.Println(ad_sort.SelectSort(arr))

	s := ad_sort.Uint64Slice{
		9,10,210,21,1,8,33,12,5521,221029291,11,
	}
	ad_sort.QuickSort(s,0,len(s)-1)
	fmt.Println(s)*/

	t := "abcababed"
	//s := "cabd"
	fmt.Println(code_string.GetNext(t))

	fmt.Println(code_string.KMPSearch("be", t))
}
