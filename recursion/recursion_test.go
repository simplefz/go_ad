package recursion

import (
	"fmt"
	"testing"
)

func TestProcessGetArrMaxValue(t *testing.T) {
	arr := []uint64{3, 4, 5, 89, 821, 2, 122, 891, 921}

	value := ProcessGetArrMaxValue(arr, 0, len(arr)-1)

	fmt.Println(value)
}
