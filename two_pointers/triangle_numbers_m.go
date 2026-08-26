package two_pointers

import "sort"

/*
	https://www.hellointerview.com/learn/code/two-pointers/valid-triangle-number

	Triangle Numbers
	DESCRIPTION (inspired by Leetcode.com)
	Write a function to count the number of triplets in an integer array nums that could form the sides of a triangle.

	For three sides to form a valid triangle, all three of these conditions must hold:
	(a + b > c), (a + c > b), and (b + c > a), where (a), (b), and (c) are the side lengths.
	In other words, the sum of every possible pair must exceed the third side.

    Input: nums = [11,4,9,6,15,18]
    Output: 10

    4, 15, 18
	6, 15, 18
	9, 15, 18
	11, 15, 18
	9, 11, 18
	6, 11, 15
	9, 11, 15
	4, 9, 11
	6, 9, 11
	4, 6, 9
*/

// Solution: The triangle condition is:  a + b > c, and c must be the longest side.  c>a,c>b,c>a+b
func TriangleNumber(nums []int) int {
	// 4 6 9 11 15 18
	sort.Ints(nums)
	count := 0
	// fix the longest side
	for i := len(nums) - 1; i > 0; i-- {
		right, left := i-1, 0
		for left < right {
			sum := nums[left] + nums[right] - nums[i]
			if sum > 0 {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}
	return count
}
