package review

/*
nums = [5,7,7,8,8,10]
target = 8
→ [3,4]

nums = [5,7,7,8,8,10]
target = 6
→ [-1,-1]
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
			// todo 找到了继续往左，不能直接return
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

func findEnd(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			result = mid
			// todo 找到了继续往右，不能直接return
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
