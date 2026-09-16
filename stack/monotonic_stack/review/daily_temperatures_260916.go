package review

/*
Input: temperatures = [73,74,75,71,69,72,76,73]，华氏度 Fahrenheit (°F) 73°F ≈ 23°C
Output: [1,1,4,2,1,1,0,0]
*/

/*
0  1  2  3  4  5  6  7
73,74,75,71,69,72,76,73

73 入栈
74 > 73 输出 1-0=1
73出栈
74进栈
75>74 输出2-1=1
74出栈
75进栈
71<75 71进栈
69<75进栈
72<75进栈
76>75
75 出栈 输出 6-2=4
*/
func dailyTemperatures(temps []int) []int {
	result := make([]int, len(temps))
	indexStack := []int{}

	for i := range temps {
		// 栈顶
		for len(indexStack) > 0 && temps[indexStack[len(indexStack)-1]] < temps[i] {
			peekIndex := indexStack[len(indexStack)-1]
			result[peekIndex] = i - peekIndex
			indexStack = indexStack[:len(indexStack)-1] // 出栈
		}
		indexStack = append(indexStack, i)
	}

	return result
}

/*
复盘：
1 make([]int, len(temps)) 初始化 stack Stack 一开始应该是空的
2 空栈时提前 peek，应该先判断 len(stack) > 0，再取栈顶
3 比较方向写反，当前更热：temps[i] > temps[stackTop]
4 peekIndex 一开始放在循环外 每次 pop 后栈顶会变，所以必须重新读取
*/
