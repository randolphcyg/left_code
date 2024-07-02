package main

import (
	"fmt"
	"slices"
)

// Program 会议
type Program struct {
	Index int
	Start int
	End   int
}

// compare 比较结束时间
func compare(o1, o2 *Program) int {
	return o1.End - o2.End
}

// bestArrange 计算可以安排的最多会议数
func bestArrange(programs []*Program, timePoint int) (int, []*Program) {
	// 按照结束时间早排序
	slices.SortFunc(programs, compare)
	result := 0
	var arrangedMettings []*Program
	// 从左往右依次遍历会议
	for _, program := range programs {
		if timePoint <= program.Start {
			result++
			timePoint = program.End
			arrangedMettings = append(arrangedMettings, program)
		}
	}
	return result, arrangedMettings
}

func main() {
	// 测试会议
	programs := []*Program{
		{Index: 1, Start: 8, End: 9},
		{Index: 2, Start: 9, End: 10},
		{Index: 3, Start: 9, End: 11},
		{Index: 4, Start: 10, End: 11},
		{Index: 5, Start: 10, End: 12},
		{Index: 6, Start: 11, End: 12},
		{Index: 7, Start: 13, End: 14},
		{Index: 8, Start: 13, End: 15},
		{Index: 9, Start: 12, End: 15},
		{Index: 10, Start: 13, End: 17},
		{Index: 11, Start: 14, End: 18},
		{Index: 12, Start: 15, End: 18},
	}

	// 早上9点
	timePoint := 9
	maxMeetings, arrangedMettings := bestArrange(programs, timePoint)

	fmt.Printf("可安排最多会议数: %d\n", maxMeetings)
	for _, program := range arrangedMettings {
		fmt.Println(program.Index)
	}
}
