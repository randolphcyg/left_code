package main

import "fmt"

func printAllFolds(N int) {
	printProcess(1, N, true)
}

// 递归过程 来到了某个节点
// i是节点的层数，N一共的层数，down == true 凹
func printProcess(i, N int, down bool) {
	if i > N {
		return
	}
	printProcess(i+1, N, true)
	if down {
		fmt.Println(i, "凹")
	} else {
		fmt.Println(i, "凸")
	}
	printProcess(i+1, N, false)
}

func main() {
	// 折纸游戏 中序遍历
	printAllFolds(4)

}
