package two_pointers

/*
	https://www.hellointerview.com/learn/code/two-pointers/two-sum

	Two Sum (Sorted Array)
	DESCRIPTION
	Given a sorted array of integers nums, determine if there exists a pair of numbers that sum to a given target.
	Example:
	Input: nums = [1,3,4,6,8,10,13], target = 13
	Output: True (3 + 10 = 13)
	Input: nums = [1,3,4,6,8,10,13], target = 6
	Output: False
*/

func TowSum(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return true
		}

		if sum > target {
			right--
		} else if sum < target {
			left++
		}
	}
	return false
}
