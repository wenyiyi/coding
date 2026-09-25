package review

/*
Input:
nums = [-1,0,3,5,9,12]
target = 9

找到 target：
→ 返回下标

找不到：
→ 返回 -1

Output:
4
*/
func search704(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}
