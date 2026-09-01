package monotonic_stack

/*
https://www.hellointerview.com/learn/code/stack/monotonic-stack
Monotonic Stack 单调栈

Next Greater Element
DESCRIPTION
Given an array of integers, find the next greater element for each element in the array.
The next greater element of an element x is the first element to the right of x that is greater than x.
If there is no such element, then the next greater element is -1.
Example
Input: [2, 1, 3, 2, 4, 3]
Output: [3, 3, 4, 4, -1, -1]
*/

// 初始化为 -1 [-1 -1 -1 -1 -1 -1]

// 2 1  <-3   1 的 next greater 是 3
// output 3   index=1
// 1 pop 出栈  [-1 3 -1 -1 -1 -1]

// 2  <-3   2 的 next greater 是 3
// output 3
// 2 pop      index=0

// 3 push 进栈
// 3 <-2
// 3 2

// 3 2 <-4
// 2 < 4 output 4 index=3
// 3 < 4 output 4 index=2
func nextGreaterElement(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}

	var indexStack []int
	for i := range n {
		for len(indexStack) > 0 && nums[indexStack[len(indexStack)-1]] < nums[i] {
			index := indexStack[len(indexStack)-1]      // pop
			indexStack = indexStack[:len(indexStack)-1] // delete
			result[index] = nums[i]                     // 不是append，是修改对应位置
		}
		indexStack = append(indexStack, i)
	}
	return result
}
