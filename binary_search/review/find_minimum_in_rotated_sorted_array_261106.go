package review

/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
For example, the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time
results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.
You must write an algorithm that runs in O(log n) time.

Example 1:

Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.

Example 2:
Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
*/

/*
二分查找
mid = left + (right-left)/2

if mid > right
4,5,6,7,0,1,2   4/5/6/7 > 2
最小值 在右边

if mid < right
最小值是mid或者，mid 的左边

*/

func findMin261106(nums []int) int {
	left := 0
	right := len(nums) - 1
	// todo 不是 <=, 目标是收缩到只剩一个候选元素，= 的时候就表示这个候选找出来了，不需要再处理了
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			// 最小值 在右边
			left = mid + 1
		} else {
			// 最小值是mid或者，mid 的左边
			right = mid
		}

	}
	return nums[left]
}
