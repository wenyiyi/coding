package review

/*
Input:
s1 = "ab"
s2 = "eidbaooo"

Output:
true


eidbaooo
ab

*/

func checkInclusion918(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map := make(map[byte]int)
	windowMap := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]]++
	}

	for i := 0; i < len(s1); i++ {
		windowMap[s2[i]]++
	}

	// 判断是否相等
	if same918(s1Map, windowMap) {
		return true
	}

	var start = 0
	for end := len(s1); end < len(s2); end++ {
		windowMap[s2[end]]++
		windowMap[s2[start]]--
		start++

		if same918(s1Map, windowMap) {
			return true
		}

	}

	return false
}

func same918(a, b map[byte]int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
