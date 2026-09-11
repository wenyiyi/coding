package stack

/*
Basic Calculator II

Given a string s which represents an expression, evaluate this expression and return its value.
The integer division should truncate toward zero.
You may assume that the given expression is always valid.
All intermediate(中间) results will be in the range of [-231, 231 - 1].
Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().

Example 1:
Input: s = "3+2*2"
Output: 7

Example 2:
Input: s = " 3/2 "
Output: 1


Example 3:
Input: s = " 3+5 / 2 "
Output: 5

All the integers in the expression are non-negative integers(非负数) in the range [0, 231 - 1].

*/

/*
+  → push(num)
-  → push(-num)
*  → top *= num
/  → top /= num

num   当前数字
op    当前数字num前面的操作符
stack 已经处理过的数字

遇到 +- 压栈
遇到 * / 立刻计算

注意减号：
12 - 3 * 4
[12, -3]
*/
func calculate(s string) int {
	// 当前数字
	currNum := 0
	// 当前数字num前面的操作符，第一个数字前面是+
	currNumOp := byte('+')
	// 存已经处理过的数字
	var numStack []int

	for i := 0; i < len(s); i++ {
		ch := s[i]

		// 12 + 3
		if ch >= '0' && ch <= '9' {
			currNum = currNum*10 + int(ch-'0')
		}

		// 当遇到操作符，才表示前面的数字读完了，比如 12 + 3 ，读到 + 才表示 12 读完了
		if ch != ' ' && (ch < '0' || ch > '9') || i == len(s)-1 { // i == len(s)-1 最后一个数字后面没有操作符了，直接处理
			switch currNumOp {
			case '+':
				numStack = append(numStack, currNum)
			case '-':
				numStack = append(numStack, -currNum) // 12 - 3 * 4  存 [12,-3]
			case '*': // 遇到 * / 立刻计算
				numStack[len(numStack)-1] *= currNum
			case '/':
				numStack[len(numStack)-1] /= currNum
			}
			// 表示下一个数字前面的操作符
			currNumOp = ch
			currNum = 0
		}
	}
	result := 0
	// 栈里的直接相加，不用再单独处理减号
	for _, num := range numStack {
		result += num
	}

	return result
}
