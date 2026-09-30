package review

/*
Input:
1 2 3
4 5 6
7 8 9

每一圈顺时针移动 1 格

4 1 2
7 5 3
8 9 6

*/

func rotateMatrix(matrix [][]int) {
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1

	for left <= right && top <= bottom {
		temp := matrix[top][left]
		// top row 不变=top，col++
		for col := left; col < right; col++ {
			matrix[top][col+1], temp = temp, matrix[top][col+1]
		}

		// right col 不变=right，row++
		for row := top; row < bottom; row++ {
			matrix[row+1][right], temp = temp, matrix[row+1][right]
		}

		// bottom row不变=bottom，col--
		for col := right; col > left; col-- {
			matrix[bottom][col-1], temp = temp, matrix[bottom][col-1]
		}

		// left col 不变=left，row--
		for row := bottom; row > top; row-- {
			matrix[row-1][left], temp = temp, matrix[row-1][left]
		}
		left++
		right--
		top++
		bottom--
	}
}
