package monotonic_stack

/*
https://www.hellointerview.com/learn/code/stack/daily-temperatures

Daily Temperatures
DESCRIPTION (inspired by Leetcode.com)
Given an integer array temps representing daily temperatures,
write a function to calculate the number of days one has to wait for a warmer temperature after each given day.
The function should return an array answer where answer[i] represents the wait time for a warmer day after the ith day.
If no warmer day is expected in the future, set answer[i] to 0.
找后面第几天会第一次出现更高温度

Inputs:

temps =

	    [65, 70, 68, 60, 55, 75, 80, 74]
		0    1   2   3   4   5   6   7

Output:

[1,4,3,2,1,1,0,0]

65 push
70
65 <70 output 1-0=1
65 pop
70 push
68 push
60 push
55 push
75
70 < 75 output 5-1=4 delete len(stack)-1 栈底
*/
func dailyTemperatures(temps []int) []int {
	var indexStack []int
	result := make([]int, len(temps))

	for i := range temps {
		// for not if
		for len(indexStack) > 0 && temps[indexStack[len(indexStack)-1]] < temps[i] {
			index := indexStack[len(indexStack)-1]            // pop
			result[index] = i - indexStack[len(indexStack)-1] // 不是append，是修改对应位置
			indexStack = indexStack[:len(indexStack)-1]       // : delete
		}
		indexStack = append(indexStack, i)
	}
	return result
}
