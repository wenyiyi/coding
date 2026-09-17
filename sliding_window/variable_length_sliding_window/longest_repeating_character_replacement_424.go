package variable_length_sliding_window

/*
https://www.hellointerview.com/learn/code/sliding-window/longest-repeating-character-replacement
https://leetcode.com/problems/longest-repeating-character-replacement/description/

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
# A A B B A
A → 3  ← 最多
B → 2
window length = 5
max frequency = 3
需要替换 = 5 - 3 = 2

windowSize - maxFrequenty <= k

窗口用start和end来表示，窗口长度=end-start+1

移动窗口
如果合法：
windowSize - maxFrequency <= k
→ 继续扩大窗口

如果非法：
windowSize - maxFrequency > k
→ frequency[s[left]]--
→ left++

核心逻辑：
 1. right 不断向右
 2. 统计当前字符 frequency
 3. 更新 maxFrequency
 4. windowSize - maxFrequency > k
    → shrink left
 5. 记录最大 windowSize
*/
func characterReplacement(s string, k int) int {
	var result int
	// var window = make([]byte, k) 不需要创建窗口，窗口本身由 start 和 end 表示了，窗口的长度 = end-start+1
	var windowMap = make(map[byte]int)
	var maxFrequency int

	start := 0
	for end := 0; end < len(s); end++ {
		windowMap[s[end]]++
		if windowMap[s[end]] > maxFrequency {
			maxFrequency = windowMap[s[end]]
		}

		// 非法      windowSize - maxFrequency <= k     windowSize=end-start
		if end-start+1-maxFrequency > k {
			windowMap[s[start]]--
			start++
		}

		// result 应该表示历史上出现过的最大合法窗口，而不是“当前窗口长度”
		result = max(result, end-start+1)
	}

	return result
}
