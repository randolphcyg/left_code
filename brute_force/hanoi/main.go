package main

import (
	"fmt"
	"strconv"
)

/*
汉诺塔问题 打印n层汉诺塔从最左边移动到最右边的全部过程
*/

// 1..N的圆盘 from to other
// 1）i~i-1 from -> other
// 2）i from -> to
// 3）1~i-1 other -> to

func process(i int, start, end, other string) {
	if i == 1 {
		fmt.Println("Move 1 from" + start + " to " + end)
	} else {
		process(i-1, start, other, end)
		fmt.Println("Move " + strconv.Itoa(i) + " from " + start + " to " + end)
		process(i-1, other, end, other)
	}
}

func hanoi(n int) {
	if n > 0 {
		process(n, "左", "右", "中")
	}
}

func main() {
	hanoi(3)
}
