package bit_operation

import (
	"fmt"
	"testing"
)

func TestRightBit(t *testing.T) {
	arr := []uint8{2, 3, 9, 11, 12, 14, 239, 255}

	for _, v := range arr {
		t.Run("test", func(t *testing.T) {
			res := rightBit(v)
			fmt.Printf("十进制%d,原二进制: %b ,^ 值: %b , ^加1 %b ，后二进制 : %b ", v, v, ^v, ^v+1, res)
		})
	}
}
