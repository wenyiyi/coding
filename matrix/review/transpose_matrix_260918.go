package review

/*
Input:
matrix = [
    [1,2,3],
    [4,5,6],
    [7,8,9],
]

Output:
[
    [1,4,7],
    [2,5,8],
    [3,6,9],
]

Input:
matrix = [
    [1,2,3],
    [4,5,6],
]

Output:
[
    [1,4],
    [2,5],
    [3,6],
]
*/

func transpose918(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	// todo rows和cols是跟原来反过来了 不是rows
	result := make([][]int, cols)
	for i := 0; i < cols; i++ {
		result[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// todo 不是 result[i][j] = matrix[j][i]
			result[j][i] = matrix[i][j]
		}
	}

	return result

}
