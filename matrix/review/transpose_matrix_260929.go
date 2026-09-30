package review

/*
Input:
[

	[1,2,3],
	[4,5,6]

]
Output:
[

	[1,4],
	[2,5],
	[3,6]

]
原矩阵：rows × cols
转置后：cols × rows
matrix[i][j]

	↓

result[j][i]

make([][]int, cols)
[nil, nil, nil]

result[j][i] = ...
*/
func transpose260929(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	result := make([][]int, cols)
	// todo 每一列也要初始化
	for i := 0; i < cols; i++ {
		result[i] = make([]int, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}
