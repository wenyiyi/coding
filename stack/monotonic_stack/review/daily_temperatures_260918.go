package review

/*
Input:
temperatures = [73,74,75,71,69,72,76,73]

Output:
[1,1,4,2,1,1,0,0]

0  1  2  3  4  5  6  7
73,74,75,71,69,72,76,73

73 进栈
74>73 输出 1-0=1
73出栈
74进栈

*/

func dailyTemperatures918(temps []int) []int {

	result := make([]int, len(temps))
	indexStack := []int{}

	for i, temperature := range temps {
		for len(indexStack) > 0 && temps[indexStack[len(indexStack)-1]] < temperature {
			// 出栈
			result[indexStack[len(indexStack)-1]] = i - indexStack[len(indexStack)-1]
			indexStack = indexStack[:len(indexStack)-1]
		}
		// 不初始化就用 append
		indexStack = append(indexStack, i)
	}

	return result
}
