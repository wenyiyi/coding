package review

/*
Input:
s1 = "ab"
s2 = "eidbaooo"

Output:
true
*/

/*
s1Map
a-1
b-1

windowMap 长度=len(1)
e-1
i-1

数量相同，长度相同，true

ei d b a o o o
——

Time:  O(n)
Space: O(1)
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map := map[rune]int{}
	windowMap := map[rune]int{}
	windowSize := len(s1)

	// 初始化
	for _, ch := range s1 {
		s1Map[ch]++
	}

	for i := 0; i < len(s1); i++ {
		windowMap[rune(s2[i])]++
	}

	if same(s1Map, windowMap) {
		return true
	}

	start := 0
	// 0 1 2 3 4 5 6 7
	// e i d b a o o o 长度8
	// ——       长度2
	// todo 一开始写成了 end + windowSize < len(s2)
	for end := start + windowSize; end < len(s2); end++ {
		windowMap[rune(s2[end])]++
		windowMap[rune(s2[start])]--
		// todo start++忘记移动了
		start++

		if same(s1Map, windowMap) {
			return true
		}

	}

	// todo 注意是返回 false
	return false
}

// todo 忘记抽出一个专门比较是否相等的方法
func same(a, b map[rune]int) bool {
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}
