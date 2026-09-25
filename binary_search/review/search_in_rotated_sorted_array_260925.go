package review

/*
给定经过旋转的严格递增数组 nums 和 target，找到返回下标，否则返回 -1
nums = [4,5,6,7,0,1,2], target = 0
→ 4

nums = [4,5,6,7,0,1,2], target = 3
→ -1

nums = [1], target = 0
→ -1
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 1 看哪边有序
		if nums[mid] >= nums[left] { // 左边有序 todo =
			// todo 2 判断target在不在有序的这一边，要跟 left 和 mid 对比
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1 // 在
			} else { // 不在
				left = mid + 1
			}
		} else { // 右边有序 todo target 要跟 mid 和 right 对比，才能确认在不在该区间
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
