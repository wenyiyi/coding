package review

/*
给定一个升序且元素不重复的数组和 target。找到就返回下标；找不到就返回它应该插入的位置
nums = [1,3,5,6], target = 2
left=0
right=3
mid=1 nums[1]=3 >2

left=0
right=mid-1=0
mid=0 nums[0]=1 <2

left=mid+1=1 todo 然后才跳出循环，最后直接返回left
right=0

跳出循环
*/
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	return left
}
