package two_pointers

import "sort"

func findPairs(nums []int, k int) int {
	// 先排序
	sort.Ints(nums)

	// 定义左右指针
	left := 0
	right := 0
	result := 0

	// map 存计算过的pair
	pairMap := make(map[[2]int]bool)

	for right < len(nums) {
		diff := nums[right] - nums[left]
		if diff < k {
			right++
		} else if diff > k {
			left++
		} else {
			pair := [2]int{nums[left], nums[right]}

			// 没出现过才计数
			if !pairMap[pair] {
				pairMap[pair] = true
				result++
			}
			right++
		}

		if left == right {
			right++
			continue
		}

	}
	return result
}
