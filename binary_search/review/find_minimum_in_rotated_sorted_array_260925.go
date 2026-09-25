package review

/*
[4,5,6,7,0,1,2]
找出其中的最小元素

nums = [3,4,5,1,2]
→ 1
3,4,5,1,2

left=0
right=4
mid=0+2=2 nums[2]=5>right
*/

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	// todo 目标是把候选区间缩到一个元素
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			// 最小值在右边
			left = mid + 1
		} else {
			// todo 说明往右是递增的，最小值不在右边
			right = mid
		}
		//if nums[mid] < nums[left] {
		//	// 最小值可能是mid，或者mid前面
		//	right = mid
		//}
	}
	return nums[left]
}
