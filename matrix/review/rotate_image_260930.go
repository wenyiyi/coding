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

// todo 这个很容易混乱，不考虑这个解法

	0 1 2

0  1 2 3
1  4 5 6
2  7 8 9

	0 1 2

0  7 4 1
1  8 5 2
2  9 6 3

1 (0,0)-(0,2)
2 (0,1)-(1,2)
3 (0,2)-(2,2)
6 (1,2)-(2,1)
9 (2,2)-(2,0)
8 (2,1)-(1,0)
7 (2,0)-(0,0)
4 (1,0)-(0,1)
----------------
1 (0,0)-(0,2)
3 (0,2)-(2,2)
9 (2,2)-(2,0)
7 (2,0)-(0,0)

2 (0,1)-(1,2)
6 (1,2)-(2,1)
8 (2,1)-(1,0)
4 (1,0)-(0,1)
----------------
Input:
1 2 3
4 5 6
7 8 9

Output:
7 4 1
8 5 2
9 6 3

// todo 沿主对角线原地转置
1 4 7
2 5 8
3 6 9

4 (0,1)-(1,0)

// 每一行交互对角
7 4 1
8 5 2
9 6 3
*/
func rotate260930(matrix [][]int) {
	// todo 原地转置
	// 当 i = 0 时，应该交换 (0,1)、(0,2)；
	// 当 i = 1 时，应该只交换 (1,2)
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// todo 每一行左右反转
	for i := 0; i < len(matrix); i++ {
		left := 0
		right := len(matrix[0]) - 1
		// todo 不止是两个边角元素需要反转
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			// todo left 和 right 需要移动
			left++
			right--
		}
	}
}
