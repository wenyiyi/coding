package binary_search

/*
Given an array of integers nums sorted in non-decreasing order非递减,
find the starting and ending position of a given target value.
If target is not found in the array, return [-1, -1].
You must write an algorithm with O(log n) runtime complexity.


Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Example 3:
Input: nums = [], target = 0
Output: [-1,-1]


Constraints:
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
*/

/*
0 1 2 3 4 5
5,7,7,8,8,10

left=0
right=5
mid=2   nums[2]=7
8>7

left=mid+1=3
right=5

mid = 4
8=8
怎么区分是end还是start？

普通 Binary Search：
找到 target → return

这道题 todo 需要往左找，再往右找
找左边界：
找到 target → 记下来 → 继续往左

找右边界：
找到 target → 记下来 → 继续往右
*/
func searchRange(nums []int, target int) []int {
	start := findStart(nums, target)
	end := findEnd(nums, target)
	return []int{start, end}
}

func findStart(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			result = mid
			// todo 找到了继续往左
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}

func findEnd(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			result = mid
			// todo 找到了继续往右
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}
