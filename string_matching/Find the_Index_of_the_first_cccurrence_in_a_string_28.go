package string_matching

/*
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

Given two strings needle and haystack,
return the index of the first occurrence of needle in haystack,
or -1 if needle is not part of haystack.

Example 1:
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Example 2:
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.

Constraints:
1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.
*/
func strStr(haystack string, needle string) int {
	// sadbutsad 长度9  sad 长度3   9-3=6 -> s
	// len(haystack)-len(needle)，i 最多只能走到还能完整放下 needle 的位置
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// 表示从 i 开始，截取一个和 needle 一样长的字符串
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func strStrKMP(haystack string, needle string) int {
	// 构建 LPS，needle 的「失败跳转表」 ，失败以后 j 应该跳到哪里的表
	lps := make([]int, len(needle))
	for i, j := 1, 0; i < len(needle); {
		if needle[i] == needle[j] {
			j++
			lps[i] = j
			i++
		} else if j > 0 {
			// 当前长度的方案失败了，直接跳到上一个合法的、更短的前后缀长度
			j = lps[j-1]
		} else {
			// 已经没有候选了
			i++
		}
	}

	// KMP 匹配
	for i, j := 0, 0; i < len(haystack); {
		if haystack[i] == needle[j] {
			i++
			j++
			// haystack = "abababc"
			//             ↑
			//             index 2
			//
			// needle   =   "ababc"
			if j == len(needle) {
				return i - j
			}
		} else if j > 0 {
			j = lps[j-1]
		} else {
			i++
		}
	}

	return -1
}

// Rabin-Karp = Sliding Window + Hash
func strStrRabinKarp(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)
	if n > h {
		return -1
	}

	// needle 的 hash
	needleHash := 0
	// 当前窗口的 hash，窗口跟needle一样长
	windowHash := 0

	for i := 0; i < n; i++ {
		needleHash += int(needle[i])
		windowHash += int(haystack[i])
	}
	// len(haystack)-len(needle)，i 最多只能走到还能完整放下 needle 的位置
	for i := 0; i <= h-n; i++ {
		// hash 一样，再比较字符串
		if windowHash == needleHash && haystack[i:i+n] == needle {
			return i
		}
		// 滑动窗口
		if i < h-n {
			windowHash -= int(haystack[i])   // 删除左边
			windowHash += int(haystack[i+n]) // 加入右边
		}
	}

	return -1
}
