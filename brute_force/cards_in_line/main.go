package main

import (
	"fmt"
)

/*
给定一个整型数组arr，代表数值不同的纸牌排成一条线。玩家A和玩家B依次拿走每张纸牌，规定玩家A先拿，玩家B后拿，
但是每个玩家每次只能拿走最左或最右的纸牌，玩家A和玩家B都绝顶聪明。请返回最后获胜者的分数。
【举例】
arr=[1,2,100,4]。 开始时，玩家A只能拿走1或4。如果开始时玩家A拿走1，则排列变为[2,100,4]，
接下来玩家B可以拿走2或4，然后继续轮到玩家A...如果开始时玩家A拿走4，则排列变为[1,2,100]，
接下来玩家B可以拿走1或100，然后继 续轮到玩家A...玩家A作为绝顶聪明的人不会先拿4，因为拿4之后，
玩家B将拿走100。所以玩家A会先拿1， 让排列变为[2,100,4]，接下来玩家B不管怎么选，100都会被玩家 A拿走。玩家A会获胜，分数为101。所以返回101。

arr=[1,100,2]。开始时，玩家A不管拿1还是2，玩家B作为绝顶聪明的人，都会把100拿走。玩家B会获胜，分数为100。所以返回100。
*/

// 暴力解法 递归 O(2^n)
func win1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	return max(first(arr, 0, len(arr)-1), second(arr, 0, len(arr)-1))
}

// 先手函数 先手自己会拿最大的
func first(arr []int, i int, j int) int {
	if i == j {
		return arr[i]
	}
	return max(arr[i]+second(arr, i+1, j), arr[j]+second(arr, i, j-1))
}

// 后手函数 别人会把最小的给我
func second(arr []int, i int, j int) int {
	if i == j {
		return 0
	}

	return min(first(arr, i+1, j), first(arr, i, j-1))
}

// 动态规划解法 O(n^2)
func win2(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	f := make([][]int, len(arr))
	s := make([][]int, len(arr))
	for i := range f {
		f[i] = make([]int, len(arr))
		s[i] = make([]int, len(arr))
	}
	for j := 0; j < len(arr); j++ {
		f[j][j] = arr[j]
		for i := j - 1; i >= 0; i-- {
			f[i][j] = max(arr[i]+s[i+1][j], arr[j]+s[i][j-1])
			s[i][j] = min(f[i+1][j], f[i][j-1])
		}
	}
	return max(f[0][len(arr)-1], s[0][len(arr)-1])
}

func main() {
	arr := []int{1, 2, 100, 4}
	fmt.Println(win1(arr))
	fmt.Println(win2(arr))
}
