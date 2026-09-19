package review

/*
Input:

1 2 3
4 5 6
7 8 9

Output:

7 4 1
8 5 2
9 6 3



*/

func rotate(matrix [][]int) {
	top := 0
	left := 0
	right := len(matrix[0]) - 1
	bottom := len(matrix) - 1

	for left < right && top < bottom {
		temp := matrix[top][left]

		// top row不变，col++
		for col := left; col < right; col++ {
			matrix[top][col+1], temp = temp, matrix[top][col+1]
		}

		// right col不变，row++
		for row := top; row < bottom; row++ {
			matrix[row+1][right], temp = temp, matrix[row+1][right]
		}

		// bottom row不变，col--
		for col := right; col > left; col-- {
			matrix[bottom][col-1], temp = temp, matrix[bottom][col-1]
		}

		// left col不变，row--
		for row := bottom; row > top; row-- {
			matrix[row-1][left], temp = temp, matrix[row-1][left]
		}
		left++
		right--
		bottom--
		top++
	}

}
