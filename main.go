package main

import (
	"fmt"
	"go_adt/bit_operation"
)

func main() {

	arr := []uint32{
		2, 2, 111, 111, 333, 333, 555, 555, 555, 777,
	}
	fmt.Println(bit_operation.XorArrOddNumberTwo(arr))
}
