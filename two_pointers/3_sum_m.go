package two_pointers

import "sort"

/*
	https://www.hellointerview.com/learn/code/two-pointers/3-sum

	3-Sum
	DESCRIPTION (inspired by Leetcode.com)
	Given an input integer array nums, write a function to find all unique triplets [nums[i], nums[j], nums[k]]
	such that i, j, and k are distinct indices,
	and the sum of nums[i], nums[j], and nums[k] equals zero.
	Ensure that the resulting list does not contain any duplicate triplets.
    input: nums = [-1,0,1,2,-1,-1]
    output: [[-1,-1,2],[-1,0,1]]
*/

// solution: sort + fix one number, then use Two Sum to find the other two numbers + skip duplicates
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	// each number can become fix number, so need to foreach
	for i := 0; i < len(nums)-2; i++ {
		// num[i] is the fix one number, also need to avoid duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// left and  right need to skip duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return result
}
