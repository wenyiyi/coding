package binary_search

/*
Suppose假设 an array of length n sorted in ascending order is rotated between 1 and n times.
For example, the array nums = [0,1,2,4,5,6,7] might become:
[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results
in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
Given the sorted rotated array nums of unique elements,
return the minimum element of this array.
You must write an algorithm that runs in O(log n) time.

Example 1:
Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.

Example 2:
Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

Example 3:
Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

Constraints:
n == nums.length
1 <= n <= 5000
-5000 <= nums[i] <= 5000
All the integers of nums are unique.
nums is sorted and rotated between 1 and n times.
*/

/*
input [3,4,5,1,2]

nums = [3,4,5,1,2]
index   0 1 2 3 4
        3 4 5 1 2
        L   M   R

nums[mid]   = 5
nums[right] = 2
5 > 2
正常升序数组不应该出现 5..2
所以从 mid 到 right 之间一定发生了旋转断点
5 ↓ 1 2
  断点
而最小值就是这个断点后面的第一个数

nums[mid] > nums[right]
→ 最小值一定在 mid 右边，但不一定就是mid+1
→ left = mid + 1


nums = [3,4,5,6,7,1,2]

left=0 right=6
mid=3
nums[mid]   = 6
nums[right] = 2

6 > 2  最小值!=nums[3+1]=nums[4]=7





nums = [4,5,1,2,3]
index   0 1 2 3 4
        4 5 1 2 3
        L   M   R

nums[mid]   = 1
nums[right] = 3

1 < 3
[mid ... right]
[1,2,3]
这一段已经正常升序
最小值有可能就是 mid


nums = [5,1,2,3,4]
left=0 right=4
mid=2
nums[mid]   = 2
nums[right] = 4
2 < 4
但不是真正的最小值


todo 重点：不断缩小最小值所在范围
*/

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	// todo 不是 <=, 目标是收缩到只剩一个候选元素，= 的时候就表示这个候选找出来了，不需要再处理了
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			// 最小值在右边
			left = mid + 1
		} else {
			// 最小值可能是mid
			right = mid
		}
	}
	return nums[left]
}
