package bit_operation

import "testing"

func TestSwapValue(t *testing.T) {
	if a, b := SwapValue(10, 18); a != 18 || b != 10 {
		t.Errorf("a,b want to swap but it not succee !, old a %d,old b %d,new a %d,new b %d", 10, 18, a, b)
	}
}

func TestXorArrOddNumberOne(t *testing.T) {
	if oddNum := XorArrOddNumberOne([]int{111, 111, 333, 333, 555, 555, 555}); oddNum != 555 {
		t.Errorf("oddNum want give in 555 ! %d", oddNum)
	}
}

func TestXorArrOddNumberTwo(t *testing.T) {
	value, value2 := XorArrOddNumberTwo([]uint32{111, 111, 333, 333, 555, 555, 555, 777})
	if value != 555 && value != 777 {
		t.Errorf("value want give in 555 or 777  ! %d", value)
	}
	if value2 != 555 && value2 != 777 {
		t.Errorf("value2 want give in 555 or 777  ! %d", value)
	}
}
