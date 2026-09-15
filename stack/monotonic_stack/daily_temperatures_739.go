package monotonic_stack

/*
https://www.hellointerview.com/learn/code/stack/daily-temperatures
https://leetcode.com/problems/daily-temperatures/description/

Example 1:
Input: temperatures = [73,74,75,71,69,72,76,73]，华氏度 Fahrenheit (°F) 73°F ≈ 23°C
Output: [1,1,4,2,1,1,0,0]

Example 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:
Input: temperatures = [30,60,90]
Output: [1,1,0]

temps的元素表示每一天的温度
今天：73
明天：74，只需等1天

今天74
明天75，只需等1天

今天75
明天71，69,72,76  需要等4天

解题思路：
我需要知道「下一个更大的元素」

	↓

单调栈

	↓

栈里存 index，不存 temperature

	↓

当前 > 栈顶对应温度

	↓

peek → 算距离 i-index → pop

	↓

继续比较新的栈顶

	↓

最后当前 index 入栈
*/
func dailyTemperatures(temps []int) []int {
	var indexStack []int
	// 不能用 var result []int，一定要初始化并创建好长度
	result := make([]int, len(temps))

	for i := range temps {
		// 记得是for，当前温度比栈顶那天高，就找到了栈顶那天的下一个更高温度，然后出栈，继续比较栈顶元素
		for len(indexStack) > 0 && temps[i] > temps[indexStack[len(indexStack)-1]] {
			index := indexStack[len(indexStack)-1] // peek 获取栈顶元素
			// 计算要等几天
			result[index] = i - index // 不是append，是修改对应位置
			// 注意要重新赋值
			indexStack = indexStack[:len(indexStack)-1] // Go 的切片 [:2]表示取 index 0到 index 2 之前
		}
		// 当前这一天还没找到答案，入栈等待
		indexStack = append(indexStack, i)
	}
	return result
}
