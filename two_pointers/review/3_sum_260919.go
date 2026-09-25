package review

import "sort"

/*
给定整数数组 nums，返回所有满足下面条件且不重复的三元组：

nums[i] + nums[j] + nums[k] == 0

要求：
i、j、k 两两不同
答案中不能出现重复的三元组

Example 1
Input:
nums = [-1, 0, 1, 2, -1, -4]
Output:
[[-1, -1, 2], [-1, 0, 1]]
*/

/*
排序
固定 一个数(也不能重复)
找两外两个数，left right

重复的数就跳过

*/

func threeSum(nums []int) [][]int {
	result := [][]int{}

	// todo 排序
	sort.Ints(nums)

	// -1, 0, 1  i=0  len=3 i<3-2=1
	// todo 应该是 i < len(nums)-2 而不是 i < len(nums)，固定 i，后面还需要 left + right 两个元素，i 后面至少留 2 个
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// todo left从i+1开始，不是0
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// todo 找到后，先移动，再跟刚才使用过的值对比
				// for left < right && nums[left] == nums[left+1] {
				//					left++
				//				}
				//				for left < right && nums[right] == nums[right-1] {
				//					right--
				//				}
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
			if sum > 0 {
				right--
			}
			if sum < 0 {
				left++
			}
		}
	}
	return result
}
