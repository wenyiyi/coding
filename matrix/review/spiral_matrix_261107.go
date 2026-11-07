package review

/*
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

1 先处理外圈，再处理内圈
2 边角不要重复处理
3 left 和 bottom 不要重复处理，只有一行的情况下
*/
func spiralOrder261107(matrix [][]int) []int {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	left := 0
	right := len(matrix[0]) - 1

	top := 0
	bottom := len(matrix) - 1

	var result []int
	for left <= right && top <= bottom {
		// top col++
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		// right row++  处理边角
		for row := top + 1; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		// bottom col-- 处理边角
		if bottom != top {
			for col := right - 1; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
		}
		// left row-- 处理边角
		if left != right {
			for row := bottom - 1; row > top; row-- {
				result = append(result, matrix[row][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}
