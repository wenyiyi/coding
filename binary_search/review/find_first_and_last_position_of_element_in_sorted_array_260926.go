package review

func searchRange260926(nums []int, target int) []int {
	start, end := findStart260926(nums, target), findEnd260926(nums, target)
	return []int{start, end}
}

func findStart260926(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			result = mid
			// 找到了，继续往左找
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	return result
}

func findEnd260926(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			result = mid
			// 找到了，继续往右找
			left = mid + 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	return result
}
