package two_pointers

import "sort"

/*
	https://www.hellointerview.com/learn/code/two-pointers/3-sum

	https://leetcode.com/problems/3sum/description/

	3-Sum
	DESCRIPTION (inspired by Leetcode.com)
	Given an input integer array nums, write a function to find all unique triplets [nums[i], nums[j], nums[k]]
	such that i, j, and k are distinct indices,
	and the sum of nums[i], nums[j], and nums[k] equals zero.
	Ensure that the resulting list does not contain any duplicate triplets.
    input: nums = [-1,0,1,2,-1,-1]
    output: [[-1,-1,2],[-1,0,1]]

	题目特征：
	找三个不同位置的数
	a + b + c = 0
	答案不能重复
	最暴力是三层循环：O(n³)

	优化的关键是：先固定一个数，把 3Sum 降维成 Two Sum
	a + b + c = 0
	固定 a
	→ b + c = -a
	→ Two Sum
	1 为了能够使用双指针，先排序
	2 然后枚举固定的第一个数
	3 剩下两个数用 left 和 right
	sum = nums[i] + nums[left] + nums[right]
	sum < 0
	→ 太小
	→ left++        // 换大一点的数

	sum > 0
	→ 太大
	→ right--       // 换小一点的数

	sum == 0
	→ 找到答案
	4 两层去重
	- 固定数 i 去重
	- 找到答案以后去重

	i 重复
	→ 整个这个 i 不处理
	→ continue

	left/right 重复
	→ 要跨过所有连续重复值
	→ for + left++ / right--

	总结：3Sum = Sort + 固定一个数 + Two Pointers + 去重。
	以及去重的本质：不是得到结果以后再查重，而是已经使用过的相同数字，不让它再次进入搜索
*/

func threeSum918(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// todo num[i]需要去重
		// nums[i] == nums[i+1]  ❌ 看“下一个是不是重复”
		// nums[i] == nums[i-1]  ✅ 看“这个是不是已经处理过”
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		// 先固定一个数，然后找另外两个
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				// 找到了收答案
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// todo 去重，不是“结果出来以后查重”，而是“已经用过的数字，不再让它进入下一轮搜索” for left < right
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}

	}
	return result
}
