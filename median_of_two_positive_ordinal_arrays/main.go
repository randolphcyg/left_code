package main

import (
	"fmt"
)

// 寻找两个正序数组的中位数

/*
双指针法 O(m+n)

双指针法合并数组：
使用两个指针 i 和 j 分别指向 nums1 和 nums2 的起始位置。
比较 nums1[i] 和 nums2[j] 的大小，将较小的值放入 merged 数组中，并移动相应的指针。
如果一个数组遍历完毕，则将另一个数组的剩余部分直接复制到 merged 数组中。
计算中位数：
计算合并后的数组总长度 total。
如果 total 是偶数，中位数是 merged[total/2-1] 和 merged[total/2] 的平均值。
如果 total 是奇数，中位数是 merged[total/2]。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	total := m + n

	// 使用双指针方法合并两个有序数组
	i, j, k := 0, 0, 0
	merged := make([]int, total)

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			merged[k] = nums1[i]
			i++
		} else {
			merged[k] = nums2[j]
			j++
		}
		k++
	}

	// 如果nums1还有剩余元素
	for i < m {
		merged[k] = nums1[i]
		i++
		k++
	}

	// 如果nums2还有剩余元素
	for j < n {
		merged[k] = nums2[j]
		j++
		k++
	}

	fmt.Println(merged)

	// 计算中位数
	if total%2 == 0 {
		return float64(merged[total/2-1]+merged[total/2]) / 2
	}
	return float64(merged[total/2])
}

/*
二分查找法 O(log(m+n))
确保数组 nums1 较短：首先确保 nums1 总是较短的数组，这样可以减少二分查找的范围。
二分查找：使用二分查找来确定合适的分割点 i 和 j，使得 nums1[0:i] 和 nums2[0:j] 组合后的两个部分满足中位数的定义。
合适的分割点：
如果 nums2[j-1] > nums1[i]，则 i 太小，需要增加 i。
如果 nums1[i-1] > nums2[j]，则 i 太大，需要减少 i。
计算中位数：
找到合适的 i 和 j 后，根据总长度的奇偶性计算中位数。
*/
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	// 确保 nums1 是较短的数组
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	imin, imax, halfLen := 0, m, (m+n+1)/2
	var maxOfLeft, minOfRight float64

	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i
		if i < m && nums2[j-1] > nums1[i] {
			// i 需要增大
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// i 需要减小
			imax = i - 1
		} else {
			// 找到了合适的 i
			if i == 0 {
				maxOfLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxOfLeft = float64(nums1[i-1])
			} else {
				maxOfLeft = float64(max(nums1[i-1], nums2[j-1]))
			}

			if (m+n)%2 == 1 {
				return maxOfLeft
			}

			if i == m {
				minOfRight = float64(nums2[j])
			} else if j == n {
				minOfRight = float64(nums1[i])
			} else {
				minOfRight = float64(min(nums1[i], nums2[j]))
			}

			return (maxOfLeft + minOfRight) / 2
		}
	}

	return 0.0
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	median := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("双指针法 The median is: %.2f\n", median)

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}

	median = findMedianSortedArrays(nums1, nums2)
	fmt.Printf("双指针法 The median is: %.2f\n", median)

	nums1 = []int{1, 3}
	nums2 = []int{2}
	median2 := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("二分法 The median is: %.2f\n", median2)

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}

	median2 = findMedianSortedArrays(nums1, nums2)
	fmt.Printf("二分法 The median is: %.2f\n", median2)
}
