package bit_operation

import (
	"testing"
)

func TestSwapValue(t *testing.T) {
	if a, b := SwapValue(10, 18); a != 18 || b != 10 {
		t.Errorf("a,b want to swap but it not succee !, old a %d,old b %d,new a %d,new b %d", 10, 18, a, b)
	}
}

func BenchmarkSwapValue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SwapValue(10, 18)
	}
}

func BenchmarkSwapValueOther(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := 10
		b := 18
		/*c := a
		a = b
		b = c*/
		a, b = b, a
	}
}

func TestXorArrOddNumberOne(t *testing.T) {
	if oddNum := XorArrOddNumberOne([]int{111, 111, 333, 333, 555, 555, 555}); oddNum != 555 {
		t.Errorf("oddNum want give in 555 ! %d", oddNum)
	}
}

func TestXorArrOddNumberOneMatch(t *testing.T) {
	cases := []struct {
		Name     string
		Param    []int
		Expected int
	}{
		{"single", []int{1, 1, 2, 3, 3, 4, 4}, 2},
		{"more", []int{1, 1, 2, 2, 2, 3, 3, 4, 4}, 2},
	}
	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			if oddNum := XorArrOddNumberOne(v.Param); oddNum != v.Expected {
				t.Errorf("oddNum want give in %d ! %d", v.Expected, oddNum)
			}
		})
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
