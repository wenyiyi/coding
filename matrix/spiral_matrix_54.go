package matrix

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
*/
func spiralOrder(matrix [][]int) []int {
	result := []int{}

	left := 0
	top := 0
	right := len(matrix[0]) - 1
	bottom := len(matrix) - 1

	// 单行单列  1 2 3     left=0 right=2 top=0 bottom=0
	for left <= right && top <= bottom {
		// todo top(包括left和right两个角) row 不变 col++
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}

		// todo right（不包含第一行的右上角，包含最后一行的右下角） col不变 row++
		for row := top + 1; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}

		// todo 单行的话，bottom=top，不能重复输出
		if top != bottom {
			// todo bottom（不包含最右的角，包含最左的角） row不变，col--
			for col := right - 1; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
		}

		// todo 单列的话，left=right
		if left != right {
			// todo left（不包含上下两个角） col不变，row++
			for row := bottom - 1; row > top; row-- {
				result = append(result, matrix[row][left])
			}
		}

		// todo 等外圈处理完再统一缩圈
		right--
		top++
		bottom--
		left++
	}
	return result
}
