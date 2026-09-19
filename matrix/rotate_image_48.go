package matrix

/*
https://leetcode.com/problems/rotate-image/description/
You are given an n x n 2D matrix representing an image,
rotate the image by 90 degrees (clockwise顺时针).

You have to rotate the image in-place原地旋转,
which means you have to modify the input 2D matrix directly.
DO NOT allocate分配 another 2D matrix and do the rotation.

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

  0 1 2
0 1 2 3
1 4 5 6
2 7 8 9

   0 1 2
 0 7 4 1
 1 8 5 2
 2 9 6 3

Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

  0  1  2  3
0 5  1  9  11
1 2  4  8  10
2 13 3  6  7
3 15 14 12 16

  0   1   2  3
0 15  13  2  5
1 14  3   4  1
2 12  6   8  9
3 16  7  10 11

*/

/*
	0  1  2  3

0 5  1  9  11
1 2  4  8  10
2 13 3  6  7
3 15 14 12 16

第一步：Transpose

	0  1  2  3

0 5  1  9  11
1 -  4  8  10
2 -  -  6  7
3 -  -  -  16

第二步：每一行左右反转

	  0 1 2
	0 7 4 1
	1 8 5 2
	2 9 6 3
*/
func rotate(matrix [][]int) {

	// 第1步，先转置
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}

	// 第2步，每一行左右反转
	for i := 0; i < len(matrix); i++ {
		left := 0
		right := len(matrix[i]) - 1

		// todo 不止是两个边角元素需要反转
		for left < right {
			matrix[i][right], matrix[i][left] = matrix[i][left], matrix[i][right]
			left++
			right--
		}
	}

}

/*
  0 1 2
0 1 2 3
1 4 5 6
2 7 8 9

   0 1 2
 0 7 4 1
 1 8 5 2
 2 9 6 3

top  col=right row=top++
1 (0,0)-(0,2)
2 (0,1)-(1,2)
3 (0,2)-(2,2)

right
6 (1,2)-(2,1)
9 (2,2)-(2,0)

bottom
8 (2,1)-(1,0)
7 (2,0)-(0,0)

left
4 (1,0)-(0,1)

先处理外圈，再处理内圈

这个版本坐标太绕了，先不记这个版本
*/

func rotate2(matrix [][]int) {
	top := 0
	left := 0
	right := len(matrix[0]) - 1
	bottom := len(matrix) - 1

	for left < right && top < bottom {
		for col := left; col < right; col++ {
			offset := col - left

			temp := matrix[top][left+offset]

			// top -> right
			matrix[top+offset][right], temp =
				temp, matrix[top+offset][right]

			// right -> bottom
			matrix[bottom][right-offset], temp =
				temp, matrix[bottom][right-offset]

			// bottom -> left
			matrix[bottom-offset][left], temp =
				temp, matrix[bottom-offset][left]

			// left -> top
			matrix[top][left+offset] = temp
		}

		top++
		bottom--
		left++
		right--
	}
}
