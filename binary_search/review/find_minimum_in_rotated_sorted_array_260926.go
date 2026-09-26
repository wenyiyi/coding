package review

/*
nums = [3,4,5,1,2]
→ 1

nums = [4,5,6,7,0,1,2]
→ 0

nums = [11,13,15,17]
→ 11


不断缩小最小值的范围
mid能不能排除
往哪边
*/

func findMin260926(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// mid 可能是最小值，或者往左找
		if nums[mid] < nums[right] {
			// todo 不能把mid排除，mid也可能是最小值 right=mid+1
			right = mid
		}
		if nums[mid] > nums[right] { // 往右找
			left = mid + 1
		}
	}
	return nums[left]
}
