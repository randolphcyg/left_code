package main

import "fmt"

// TODO
func main() {
	arr := make([]int, 10) // 32bit * 10 -> 320 bits
	// arr[0] int 0 ~ 31
	// arr[1] int 32~ 63
	i := 178 // 想取得178个bit的状态

	numIndex := i / 32
	bitIndex := i % 32
	// 拿到178位的状态
	s := (arr[numIndex] >> bitIndex) & 1
	// 把178位的状态改成1
	arr[numIndex] = arr[numIndex] | (1 << bitIndex)
	fmt.Println(s)
	fmt.Println(arr[numIndex])
	// 把178位的状态改成0
	arr[numIndex] = arr[numIndex] & ^(1 << bitIndex)
	fmt.Println(arr[numIndex])

	// 把178位的状态拿出来
	bit := arr[i/32] >> (i % 32) & 1
	fmt.Println(bit)
}
