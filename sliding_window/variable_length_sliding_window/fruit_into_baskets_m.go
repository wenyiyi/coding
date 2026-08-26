package variable_length_sliding_window

/*
https://www.hellointerview.com/learn/code/sliding-window/variable-length

Fruit Into Baskets
DESCRIPTION (inspired by Leetcode.com)
Write a function to calculate the maximum number of fruits you can collect from an integer array fruits,
where each element represents a type of fruit.
You can start collecting fruits from any position in the array,
but you must stop once you encounter a third distinct type of fruit.
The goal is to find the longest subarray where at most two different types of fruits are collected.

Example:
Input: fruits = [3, 3, 2, 1, 2, 1, 0]
Output: 4
Explanation: We can pick up 4 fruit from the subarray [2, 1, 2, 1]
*/

func FruitIntoBaskets(fruits []int) int {
	start := 0
	basketMap := map[int]int{} // Use a map to track duplicates.
	maxLen := 0

	for end := range fruits {
		basketMap[fruits[end]]++

		// Keep adding elements from the right.
		// If the window becomes invalid, remove elements from the left
		// until it becomes valid again.
		for len(basketMap) > 2 {
			basketMap[fruits[start]]--
			if basketMap[fruits[start]] == 0 {
				delete(basketMap, fruits[start])
			}
			start++
			continue
		}
		// window length = end - start + 1
		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}
