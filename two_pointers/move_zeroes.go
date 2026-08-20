package two_pointers

/*
	Given an integer array nums, write a function to rearrange the array by moving all zeros to the end
	while keeping the order of non-zero elements unchanged.
	Perform this operation in-place without creating a copy of the array.

	input: [2,0,4,1,9]
	output: [2,4,9,1,0]
*/

func moveZeroes(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		// No zero: left and right move together.
		// Zero found: left stops, right finds the next non-zero.
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
