package review

/*
	https://leetcode.com/problems/spiral-matrix/
   （卓驭一面面试题）

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

1 2 3
4 5 6
7 8 9

输出 1 ->2-> 3-> 6-> 9-> 8-> 7-> 4-> 5


Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

1 2  3  4
5 6  7  8
9 10 11 12

Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

/*
一圈一圈向内输出矩阵：
先处理外圈
再处理内圈

todo 四条边的角不能重复
todo 只有1行，bottom不能重复输出；只有1列，right不能重复输出
*/
func spiralOrder261106(matrix [][]int) []int {
	// col
	left := 0
	// col
	right := len(matrix[0]) - 1
	// row
	top := 0
	bottom := len(matrix) - 1

	result := []int{}

	// todo 是for 不是if
	for left <= right && top <= bottom {
		// top，col++
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		// right，row++
		for row := top + 1; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		// todo 只有1行，bottom不能重复输出
		if bottom != top {
			// bottom，col--
			for col := right - 1; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
		}
		// todo 只有1列，left不能重复输出
		if left != right {
			// left, row--
			for row := bottom - 1; row > top; row-- {
				result = append(result, matrix[row][left])
			}
		}

		left++
		right--
		bottom--
		top++
	}
	return result

}
