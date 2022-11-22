package main

import (
	"fmt"
)

// Assuming that "i-th bit" means that i=0 is rightmost bit
// We`ll use bitwise operators. Set bit to 1 is simple: we move one bit i bits to the left and use bitwise OR.
// To set bit to zero, we move one bit i bits to the left, reverse mask (this give smth like 11...11011..11) and use bitwise AND
func SetBit(num *int64, i int, bit int) {
	if i >= 64 {
		return
	}
	switch bit {
	case 1:
		*num = *num | 1<<i
	case 0:
		*num = *num & ^(1 << i)
	}
}

func main() {
	var num int64
	//let`s set value of num to 9 (1001 in binary)
	SetBit(&num, 0, 1)
	SetBit(&num, 3, 1)
	fmt.Println(num)
	num = 15
	//now let`s turn 15 into 10 (1111 -> 1010)
	SetBit(&num, 2, 0)
	SetBit(&num, 0, 0)
	fmt.Println(num)
}
