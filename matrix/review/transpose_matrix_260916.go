package review

/*
2    4   -1
-10  5   11
18  -7   6

变成

2    -10   18
4     5    -7
-1    11    6

	0 1 2

0 1 2 3
1 4 5 6
变成

	0 1

0 1 4
1 2 5
2 3 6

1 (0,0)->(0,0)
2 (0,1)->(1,0)
3 (0,2)->(2,0)
*/
func transpose(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	// 行数 = 原来的列数
	var result = make([][]int, cols)
	for i := range result {
		// 列数 = 原来的行数
		result[i] = make([]int, rows)
	}

	// 行
	for row := 0; row < rows; row++ {
		// 列
		for col := 0; col < cols; col++ {
			// todo 直接填值就行，不需要交换
			result[col][row] = matrix[row][col]
		}
	}
	return result
}

/*
复盘：
1 “只遍历一半”只适用于方阵原地交换，不适用于创建新矩阵
2
新建 result？
    ↓ YES
全部遍历
result[col][row] = matrix[row][col]

方阵 + 要求原地？
    ↓ YES
只遍历一半
matrix[row][col] ↔ matrix[col][row]
*/
