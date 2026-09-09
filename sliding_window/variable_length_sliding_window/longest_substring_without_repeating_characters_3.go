package variable_length_sliding_window

/*
https://www.hellointerview.com/learn/code/sliding-window/longest-substring-without-repeating-characters
https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

Longest Substring Without Repeating Characters
DESCRIPTION (inspired by Leetcode.com)
Write a function to return the length of the longest substring in a provided string s where
all characters in the substring are distinct.

Example 1:
Input: s = "eghghhgg"
Output: 3
The longest substring without repeating characters is "egh" with length of 3.

Example 2:
Input:s = "substring"
Output:8
The answer is "ubstring" with length of 8.
*/

/*
 */
func longestSubstringWithoutRepeat(s string) int {
	start := 0
	maxLen := 0
	// 维护一个窗口 [start, end]，保证窗口里面没有重复字符
	strMap := make(map[byte]int)

	for end := 0; end < len(s); end++ {
		// 1 向右移动，加入元素
		strMap[s[end]]++
		// 2 不合法就缩右边
		for strMap[s[end]] > 1 {
			strMap[s[start]]--
			start++
		}
		// 3 此时一定合法，就看最大长度
		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}
