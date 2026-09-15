package matrix

/*
https://leetcode.com/problems/transpose-matrix/
Given a 2D integer array matrix, return the transpose of matrix.
The transpose of a matrix is the matrix flipped over its main diagonal, switching the matrix's row and column indices.
2    4   -1
-10  5   11
18  -7   6

变成

2    -10   18
4     5    -7
-1    11    6


1 2 3
4 5 6
变成
1 4
2 5
3 6

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[1,4,7],[2,5,8],[3,6,9]]

Example 2:
Input: matrix = [[1,2,3],[4,5,6]]
Output: [[1,4],[2,5],[3,6]]

Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 1000
1 <= m * n <= 105
-109 <= matrix[i][j] <= 109
*/

/*
左上角不变 matrix[0][0] matrix[1][1]

我们实际上只需要遍历主对角线的一侧：
      0  1  2
0     -  ✓  ✓
1        -  ✓
2           -

2    4   -1
-10  5   11
18  -7   6

变成

2    -10   18
4     5    -7
-1    11    6

4 (0,1)->(1,0)
-1 (0,2)->(2,0)

如果外层是
for row := 0; row < len(matrix); row++
那么内层的 col 应该从哪里开始？
for col := row+1; col < len(matrix[0]); col++


坑：不保证是正方形矩阵
1 2 3
4 5 6
变成
1 4
2 5
3 6

原矩阵：matrix[row][col]
转置后：result[col][row]
(row, col) → (col, row)
不是所有 Matrix 题都要缩圈
*/

func transpose(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	result := make([][]int, cols)
	for i := 0; i < cols; i++ {
		result[i] = make([]int, rows)
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			result[col][row] = matrix[row][col]
		}
	}
	return result
}
