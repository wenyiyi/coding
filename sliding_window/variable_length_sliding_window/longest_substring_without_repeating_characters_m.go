package variable_length_sliding_window

/*
https://www.hellointerview.com/learn/code/sliding-window/longest-substring-without-repeating-characters

Longest Substring Without Repeating Characters
DESCRIPTION (inspired by Leetcode.com)
Write a function to return the length of the longest substring in a provided string s where all characters in the substring are distinct.

Example 1:
Input: s = "eghghhgg"
Output: 3
The longest substring without repeating characters is "egh" with length of 3.

Example 2:
Input:s = "substring"
Output:8
The answer is "ubstring" with length of 8.
*/

func longestSubstringWithoutRepeat(s string) int {
	start := 0
	maxLen := 0
	strMap := make(map[byte]int)

	for end := range s {
		// add elements from the right
		strMap[s[end]]++
		// If the window becomes invalid, remove elements from the left
		// until it becomes valid again.
		for strMap[s[end]] > 1 {
			strMap[s[start]]--
			start++
		}
		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}
