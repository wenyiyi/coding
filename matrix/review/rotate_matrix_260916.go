package review

/*
For example1:
Input:

	0 1 2

0 1 2 3
1 4 5 6
2 7 8 9
Output:

	0 1 2

0 4 1 2
1 7 5 3
2 8 9 6

1-2-3-6-9-8-7-4
变成
4-1-2-3--6-9-8-7

top
1 (0,0)->(0,1)
2 (0,1)->(0,2)
right
3 (0,2)->(1,2)
6 (1,2)->(2,2)
bottom
9 (2,2)->(2,1)
8 (2,1)->(2,0)
left
7 (2,0)->(1,0)
4 (1,0)->(0,0)

example2
input：

	1   2   3   4
	5   6   7   8
	9  10  11  12

13  14  15  16

外圈

	1 →  2 →  3 →     4
	↑                 ↓
	5     6    7      8
	↑                 ↓
	9    10   11     12
	↑                 ↓

13 ← 14 ← 15 ←    16
*/
func rotateMatrix1(matrix [][]int) [][]int {
	right := len(matrix[0]) - 1
	bottom := len(matrix) - 1
	left := 0
	top := 0

	// 再处理内圈
	for left < right && top < bottom {

		// 顺序别写反了
		var temp = matrix[top][left] //todo temp 存左上角的数

		// 先输出外圈

		// 上边
		// row=0,往右移动  1 (0,0)->(0,1)   2 (0,1)->(0,2)  temp=1，把1放入(0,2)，把原来的(0,2)放入temp
		for col := left; col < right; col++ {
			matrix[top][col+1], temp = temp, matrix[top][col+1]
		}

		// 右边
		// 3 (0,2)->(1,2)
		// 6 (1,2)->(2,2)
		for row := top; row < bottom; row++ {
			matrix[row+1][right], temp = temp, matrix[row+1][right]
		}

		// 下边
		// 9 (2,2)->(2,1)
		// 8 (2,1)->(2,0)
		for col := right; col > left; col-- {
			matrix[bottom][col-1], temp = temp, matrix[bottom][col-1]
		}

		// 左边
		// 7 (2,0)->(1,0)
		// 4 (1,0)->(0,0)
		for row := bottom; row > top; row-- {
			matrix[row-1][left], temp = temp, matrix[row-1][left]
		}

		// 往里走
		left++
		right--
		bottom--
		top++
	}
	return matrix
}

/*
复盘写错的地方：
1 temp 想复杂了，直接用 matirx[][] 左上角来存就行
2 top/right 含义混淆，一开始 top := len(matrix[0])-1， top/bottom 管 row；left/right 管 col
3 <= 导致越界，col <= right、row <= botto，因为访问 col+1/row+1，所以必须 < right / < bottom
4 外圈坐标写死，matrix[0][...]、从 col := 0 开始，内圈要使用当前 top/left/right/bottom
5 下边忘记 left col > 0 内圈不能回到 0，要 col > left
6 左边用了 right for row := right... row 对应纵向边界，所以从 bottom 开始
7 cursor 写成 boundary 循环 row--，里面却一直 matrix[bottom-1] row 在移动，就应该使用 matrix[row-1][left]
8 row/col 顺序写反 matrix[left][top] matrix[row][col] → matrix[top][left]

核心问题：
① row / col 和四个 boundary 容易混
row ↔ top / bottom
col ↔ left / right
② boundary 和 cursor 容易混
top/bottom/left/right = 当前这一圈的边界
row/col                = 真正移动的游标

重写时，用这四句话在脑子里走
上：top 固定，col 从 left → right
右：right 固定，row 从 top → bottom
下：bottom 固定，col 从 right → left
左：left 固定，row 从 bottom → top
*/
