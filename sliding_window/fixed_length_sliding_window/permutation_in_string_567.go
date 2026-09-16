package fixed_length_sliding_window

/*
https://leetcode.com/problems/permutation-in-string/description/
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
In other words, return true if one of s1's permutations is the substring of s2.

Example 1:
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").

Example 2:
Input: s1 = "ab", s2 = "eidboaoo"
Output: false


Constraints:
1 <= s1.length, s2.length <= 104
s1 and s2 consist of lowercase English letters.
*/

/*
0 1 2 3 4 5 6 7
e i d b a o o o 长度8

ab 长度2
8-2=6

window-> size=len(s1)

hashMap a->1   b->1  if count equal, len equal, true
*/

// 暴力解法
func checkInclusion(s1 string, s2 string) bool {
	s1Map := map[rune]int{}

	for _, ch := range s1 {
		if s1Map[ch] > 0 {
			s1Map[ch] = s1Map[ch] + 1
		} else {
			s1Map[ch] = 1
		}
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		windowMap := map[rune]int{}
		for j := i; j < i+len(s1); j++ {
			if windowMap[rune(s2[j])] > 0 {
				windowMap[rune(s2[j])] += 1
			} else {
				windowMap[rune(s2[j])] = 1
			}
		}

		result := true
		for ch, count := range s1Map {
			if windowMap[ch] != count {
				result = false
			}
		}
		if result {
			return result
		}
	}
	return false
}
