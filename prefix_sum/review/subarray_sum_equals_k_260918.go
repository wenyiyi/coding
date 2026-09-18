package review

/*

Input:
nums = [1,1,1]
k = 2

Output:
2

[1,1]  // index 0~1
[1,1]  // index 1~2


Input:
nums = [1,2,3]
k = 3

Output:
2
[1,2]
[3]

nums 中可以有正数、负数和 0。


1
当前和 1
前缀和 1

2
当前和 1+2=3
3-3=0 前缀和0存在
+数量

*/

func subarraySum918(nums []int, k int) int {
	var result int
	prefixSumMap := map[int]int{0: 1}
	var currentSum int

	for i := range nums {
		currentSum += nums[i]
		result += prefixSumMap[currentSum-k]
		prefixSumMap[currentSum]++

	}
	return result
}
