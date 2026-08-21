package two_pointers

import "fmt"

/*
		DESCRIPTION (inspired by Leetcode.com)
		Write a function to sort a given integer array nums in-place (and without the built-in sort function),
		where the array contains n integers that are either 0, 1, and 2
		and represent the colors red, white, and blue. Arrange the objects so that same-colored ones are adjacent,
		in the order of red, white, and blue (0, 1, 2).

	 	Input: nums = [2,1,2,0,1,0,1,0,1]
		Output: [0,0,0,1,1,1,1,2,2]
*/
func SortColors1(nums []int) {
	left := 0
	// Process 0, 1, and 2 in order.
	for i := range 3 {
		for j := left; j < len(nums); j++ {
			if nums[j] == i {
				// Move the matched number to the left
				nums[left], nums[j] = nums[j], nums[left]
				left++
			}
		}
	}
	fmt.Println(nums)
}

// Move 0s left, 2s right, and 1s in the middle
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	i := 0
	for i <= right {
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			// 0: Left side is already processed → i++
			i++
		} else if nums[i] == 2 {
			nums[right], nums[i] = nums[i], nums[right]
			right--
			// 2: Right side is unprocessed → don't i++
		} else {
			i++
		}
	}
}
