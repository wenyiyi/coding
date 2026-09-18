package review

/*
请返回数组中 和等于 k 的连续子数组（subarray）的数量

Input:
nums = [1, 1, 1]
k = 2

Output:
2

1
1+1=2
1+1+1=3

k=2

Input:
nums = [1, 2, 3]
k = 3

Output:
2

1
1+2=3
1+2+3=6

当前和 - k = 需要的前缀和

1
当前和 1
1-k=1-3=-2
1放入前缀和数组

2
当前和 3
3-k=0 前缀和=0存在， 数量+前缀和=0的数量
3放入前缀和数组，数量1

3
当前和6
6-3=3 前缀和=3存在 数量+前缀和=3的数量

有可能有负数和0
*/
func subarraySum(nums []int, k int) int {
	var prefixMap = map[int]int{0: 1}
	var currSum = 0
	var result int

	for i := range nums {
		currSum += nums[i]
		needPrefixSum := currSum - k
		//todo if prefixMap[needPrefixSum] > 0 { 可以直接不要，默认有0值
		result += prefixMap[needPrefixSum]
		//}
		prefixMap[currSum]++
	}

	return result
}
