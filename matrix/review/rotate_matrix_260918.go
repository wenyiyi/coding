package review

/*
Input:
1 2 3
4 5 6
7 8 9

Output:
4 1 2
7 5 3
8 9 6

	0 1 2

0 1 2 3
1 4 5 6
2 7 8 9

top
1 (0,0)->(0,1)  (0,1) 暂存matrix[0][0]
*/
func rotateMatrix918(matrix [][]int) [][]int {
	top := 0
	left := 0
	// todo  right（列数） 和 bottom（行数） 写反了
	// right := len(matrix) - 1
	// bottom := len(matrix[0]) - 1
	right := len(matrix[0]) - 1
	bottom := len(matrix) - 1

	for left < right && top < bottom {
		// todo matrix[top][left] = 1  是要用一个变量来临时存，每一圈初始化一次
		temp := matrix[top][left]

		// 外圈
		// top  row固定=left  col++
		for col := left; col < right; col++ {
			// todo 不是交换，是临时存到 temp matrix[top][col+1], matrix[top][col] = matrix[top][col], matrix[top][col+1]
			matrix[top][col+1], temp = temp, matrix[top][col+1]
		}

		// right col固定=right  row++
		for row := top; row < bottom; row++ {
			matrix[row+1][right], temp = temp, matrix[row+1][right]
		}

		// bottom row固定=bottom col-- todo 缩圈后不需要走到0了
		for col := right; col > left; col-- {
			matrix[bottom][col-1], temp = temp, matrix[bottom][col-1]
		}

		// left col固定=left  row-- todo 缩圈后不需要走到0了
		for row := bottom; row > top; row-- {
			matrix[row-1][left], temp = temp, matrix[row-1][left]
		}

		//缩圈
		left++
		top++
		right--
		bottom--

	}
	return matrix
}
