package variable_length_sliding_window

/*
https://www.hellointerview.com/learn/code/sliding-window/longest-repeating-character-replacement

Longest Repeating Character Replacement
DESCRIPTION (inspired by Leetcode.com)
Write a function to find the length of the longest substring containing the same letter in a given string s,
after performing at most k operations in which you can choose any character of the string and change it to any other uppercase English letter.

Input:s = "BBABCCDD" k = 2
Output: 5
Explanation: Replace the first 'A' and 'C' with 'B' to form "BBBBBCDD".
The longest substring with identical letters is "BBBBB", which has a length of 5.

*/

/*
思路：计算符合条件的最大窗口的长度，符合条件的最大窗口的长度 - 窗口内相同字符的个数 = k
Solution: Find the longest valid window
A valid window: window length - the count of the most frequent character <= k
*/
func characterReplacement(s string, k int) int {
	start := 0
	countMap := make(map[byte]int)
	maxSameCount := 0 // The maximum number of the same character in the current window
	result := 0

	for end := range s {
		countMap[s[end]]++
		maxSameCount = max(maxSameCount, countMap[s[end]])

		// replacements needed = window length - most same character count
		if end-start+1-maxSameCount > k {
			countMap[s[start]]--
			start++
		}

		result = max(result, end-start+1)
	}

	return result
}
