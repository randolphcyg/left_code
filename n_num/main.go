package main

import "fmt"

// 在一个无序不重复数组中，求两个元素相加值为100的有多少对
//思路：hash表法——给定一个map，遍历数组，如果该元素另一半不在map中则将元素塞入，
//如果另一半在map中，count加1。

func GetSumNum(arr []int, sum int) int {
	tmp := make(map[int]int) //map存数组里的元素，注意，这里map初始为空
	var count int            //计个数

	for i := 0; i < len(arr); i++ { //遍历数组
		//判断这个元素的“另一半”是否已经在map里
		if _, ok := tmp[sum-arr[i]]; ok { //如果在，就说明数组中有这个元素的“另一半”
			count++
		} else { //如果不在，就把这个元素加入map中
			tmp[arr[i]] = i
			fmt.Println(tmp)
		}

	}
	return count
}

func main() {
	array := []int{1, 12, 14, 40, 56, 60, 88, 78, 99}
	sum := 100
	fmt.Println(GetSumNum(array, sum))
}
