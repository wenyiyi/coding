package prefix_sum

/*
https://leetcode.com/problems/subarray-sum-equals-k/

Example 1:
Input: nums = [1,1,1], k = 2
Output: 2

Example 2:
Input: nums = [1,2,3], k = 3
Output: 2


Constraints:
1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107

*/

/*
i=0
sum=1

i=2
sum=2
sum == k
result++

sum 从i=2开始重新加（怎么知道从哪里开始重新加？？？）
i=2
sum=1
i=3
sum=1+1=2

需要两个指针
start → 子数组从哪里开始
end   → 子数组延伸到哪里

start = 0
	end = 0 → [1]
	end = 1 → [1,1]
	end = 2 → [1,1,1]

start = 1
	end = 1 → [1]
	end = 2 → [1,1]

start = 2
	end = 2 → [1]
*/
// 暴力解法 O(n*n)
func subarraySum(nums []int, k int) int {
	result := 0
	for start := 0; start < len(nums); start++ {
		var sum int
		for end := start; end < len(nums); end++ {
			sum += nums[end]
			if sum == k {
				result++
				// 不能 break：后续加入负数或 0 后，sum 仍可能再次等于 k
			}
		}
	}
	return result
}

/*
nums = [1, 2, 3]
index 0: 1
index 1: 1 + 2 = 3
index 2: 1 + 2 + 3 = 6

prefixSum:
1, 3, 6

如何已知
nums[0..2] 的和 = 6
nums[0..0] 的和 = 1

那么
nums[1..2] = 2+3 = nums[0..2]-nums[0..0]

Prefix Sum（前缀和）：已经算过的东西不要重新加，两个前缀和相减就能得到中间一段的和

假设当前走到 end
nums = [1, 2, 3]
              ↑
             end
当前前缀和：currentSum = 6

我们想找一个连续子数组，它的：
subarraySum = k = 5

刚才你已经知道： currentSum - 前面某个prefixSum = k
需要保存之前的prefixSum，可以用 HashMap

nums = [1, 2, 3]   k=3
读 1
currentSum=1
保存1

读 2
currentSum=1+2=3 =k
保存3
result++

读 3
currentSum=1+2+3=6
保存 6

prefixMap := map[int]int{}  key=某个prefixSum     value=这个prefixSum之前出现了多少次
这里马上有一个 LC560 最经典的小坑
nums = [3]
k = 3

走到 3：
currentSum = 3
要找：
currentSum - k
= 3 - 3
= 0

这时候我们需要：
prefixMap[0]，一个数字都还没取的时候，前缀和是 0。在开始遍历数组之前，前缀和 0 已经存在 1 次。

*/
// Prefix Sum
func subarraySum2(nums []int, k int) int {
	result := 0
	currentSum := 0

	// map 初始化，记住格式 :
	prefixMap := map[int]int{0: 1}

	for _, num := range nums {
		currentSum += num
		// 先查历史
		result += prefixMap[currentSum-k]
		// 再把自己加入历史
		prefixMap[currentSum]++

		/*
			为什么先查后存？
			nums = [1]
			k = 0

			开始
			prefixMap = {0: 1}
			currentSum = 0

			读到1
			currentSum = 1

			先存
			prefixMap = {{0: 1},{1: 1}}
			在查
			currentSum-0=1
			prefixMap[1]=1 但实际上根本没有和为 0 的子数组

			因为刚刚先把自身放进去了，应该先查历史，再把自己加入历史
		*/
	}

	return result
}
