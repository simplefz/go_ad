package fib

import "fmt"

/**
 斐波那契数
 F_{0}=0
 F_{1}=1}F_{1}=1
 F_{n}=F_{{n-1}}+F_{{n-2}}（n≧2)
**/

func PrintFib0To10() {
	for i := 0; i <= 10; i++ {
		fmt.Println(fib(i))
	}
}

func fib(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return fib(num-1) + fib(num-2)
}
