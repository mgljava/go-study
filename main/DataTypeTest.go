package main

import (
	"fmt"
	"math"
)

func main() {

	// bool类型
	var b1, b2 bool
	b1 = false
	b2 = true
	fmt.Println(b1)
	fmt.Println(b2)

	// 整数类型
	var sum int
	var a = 10
	b := 20
	sum = a + b
	fmt.Println(sum)

	var num int8 = 127
	fmt.Println(num + 1)

	fmt.Printf("Min(int8) = %v, Max(int8) = %v\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("Min(int16) = %v, Max(int16) = %v\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("Min(int32) = %v, Max(int32) = %v\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("Min(int64) = %v, Max(int64) = %v\n", math.MinInt64, math.MaxInt64)
}
