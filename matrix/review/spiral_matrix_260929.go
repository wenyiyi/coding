package review

/*
给定一个 m × n 矩阵，按照顺时针螺旋顺序返回所有元素。

Input:
[

	[1,2,3],
	[4,5,6],
	[7,8,9]

]

Output:
[1,2,3,6,9,8,7,4,5]

先输出外圈，再输出内圈

todo 四条边的角不能重复
todo 只有1行，bottom不能重复输出；只有1列，right不能重复输出
*/
func spiralOrder(matrix [][]int) []int {
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1

	result := make([]int, 0)

	for left <= right && top <= bottom {
		// top
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		// right（上1不能输出）
		for row := top + 1; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		if top != bottom { //todo
			// bottom（右1不能输出）
			for col := right - 1; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
		}
		if left != right { //todo
			// left（上下角不能输出）
			for row := bottom - 1; row > top; row-- {
				result = append(result, matrix[row][left])
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return result
}
