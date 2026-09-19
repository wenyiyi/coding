package review

/*
matrix =
[

	[1, 2, 3],
	[4, 5, 6]

]

[

	[1, 4],
	[2, 5],
	[3, 6]

]
*/
func transpose919(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])

	var result = make([][]int, cols)
	for k := range cols {
		result[k] = make([]int, rows)
	}

	// todo  i := range rows  等价于 for i := 0; i < rows; i++
	for i := range rows {
		for j := range cols {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}
