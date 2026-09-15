package matrix

/*
https://www.naukri.com/code360/problems/rotate-matrix_981260?utm_source=chatgpt.com
(traveloka一面面试题)

Given a 2-dimensional matrix of size ‘N’ x ‘M’, rotate the elements of the matrix clockwise.

For example1:
Input:
1 2 3
4 5 6
7 8 9
Output:

4 1 2
7 5 3
8 9 6

第1步：先观察每个数字怎么移动
1 → 2 → 3
↑       ↓
4       6
↑       ↓
7 ← 8 ← 9

1 → 2 → 3 → 6 → 9 → 8 → 7 → 4
变成
4 → 1 → 2 → 3 → 6 → 9 → 8 → 7

第2步：覆盖之前需要保留之前的值
temp := matrix[0][1]
matrix[0][1] = matrix[0][0]

第3步：能不能一直只用同一个变量来保存之前的值
先保存 2
1 → 2的位置
再保存 3
2 → 3的位置
再保存 6
3 → 6的位置

我想到
temp := matrix[0][1]     // 保存 2
matrix[0][1] = matrix[0][0] // 1 移过去

temp2 := matrix[0][2]    // 保存 3
matrix[0][2] = temp      // 2 移过去

但是没想到可以用！！！！！
temp, matrix[0][2] = matrix[0][2], temp
temp = 2
matrix[0][2] = 3
会变成
temp = 3
matrix[0][2] = 2

第4步：3 原本在右上角，顺时针移动一格后，它应该往下移动 matrix[1][2]
(0,2) -> (1,2)
temp,matrix[1][2] = matrix[1][2], temp

第5步: 6
(2,2) -> (2,1)
temp, matrix[2][1] = matrix[2][1], temp

第6步: 9
(2,2) -> (2,1)
temp, matrix[2][1] = matrix[2][1], temp

第7步: 8
(2,1) -> (2,0)
temp, matrix[2][0] = matrix[2][0], temp

第8步: 7
(2,0) -> (1,0)
temp, matrix[1][0] = matrix[1][0], temp

第9步: 4
(1,0) -> (0,0)
temp, matrix[0][0] = matrix[0][0], temp


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

里面还有一圈
6 → 7
↑   ↓
10 ← 11
外圈处理完以后，怎么让同样的四个循环继续处理内圈
我面试想到了4个边界
top
bottom
left
right

外圈：
top    = 0
bottom = 3
left   = 0
right  = 3
最外圈处理完之后，要进入里面这一圈
        ↓ 处理完

内圈：
top    = 1
bottom = 2
left   = 1
right  = 2

先存左上角，沿上右下左接力走一圈，走完缩四个边界，再处理下一圈
(先处理外圈，再处理内圈)
1 先自己画一个 4×4，标出 → ↓ ← ↑。
2 自己说出 top / bottom / left / right 分别是什么。
3 想出 temp 接力，解决覆盖问题。
4 自己写四条边。
5 自己想什么时候缩圈、什么时候停止。
6 写完用 3×3 和 4×4 dry run。
*/

// 外圈写死版本
func rotateMatrixFirst(matrix [][]int) [][]int {
	bottom := len(matrix) - 1
	right := len(matrix[0]) - 1
	var temp = matrix[0][0]

	// 上边 col++，row=0，col不能=right，不然 col+1 就越界了
	for col := 0; col < right; col++ {
		temp, matrix[0][col+1] = matrix[0][col+1], temp
	}

	// 右边 row++,col=cols
	// 3 原本在右上角，顺时针移动一格后，它应该往下移动 matrix[1][2]
	//	(0,2) -> (1,2)
	//	temp,matrix[1][2] = matrix[1][2], temp
	for row := 0; row < bottom; row++ {
		temp, matrix[row+1][right] = matrix[row+1][right], temp
	}

	// 下边：col--,row=rows
	// 9
	// (2,2) -> (2,1) -> (2,0)
	// temp, matrix[2][1] = matrix[2][1], temp
	for col := right; col > 0; col-- {
		temp, matrix[bottom][col-1] = matrix[bottom][col-1], temp
	}

	// 左边：row--,col=0
	//  7
	// (2,0) -> (1,0)
	// temp, matrix[1][0] = matrix[1][0], temp
	for row := bottom; row > 0; row-- {
		temp, matrix[row-1][0] = matrix[row-1][0], temp
	}
	return matrix
}

// 处理内圈
// 外圈：
//
//		top    = 0
//		bottom = 3
//		left   = 0
//		right  = 3
//		最外圈处理完之后，要进入里面这一圈
//	       ↓ 处理完
//
//		内圈：
//		top    = 1   top++
//		bottom = 2   bottom--
//		left   = 1   left++
//		right  = 2   right--
func rotateMatrix(matrix [][]int) [][]int {
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	// temp = matrix[0][0] 这个只初始化了一次。进入第二圈的时候，我们应该重新拿matrix[1][1]，当前圈的左上角
	for top < bottom && left < right { // 一圈一圈处理
		var temp = matrix[top][left]

		// 上边 col++，row=0，col不能=right，不然 col+1 就越界了
		for col := left; col < right; col++ {
			temp, matrix[top][col+1] = matrix[top][col+1], temp
		}

		// 右边 row++,col=cols
		// 3 原本在右上角，顺时针移动一格后，它应该往下移动 matrix[1][2]
		//	(0,2) -> (1,2)
		//	temp,matrix[1][2] = matrix[1][2], temp
		for row := top; row < bottom; row++ {
			temp, matrix[row+1][right] = matrix[row+1][right], temp
		}

		// 下边：col--,row=rows
		// 9
		// (2,2) -> (2,1) -> (2,0)
		// temp, matrix[2][1] = matrix[2][1], temp
		for col := right; col > left; col-- {
			temp, matrix[bottom][col-1] = matrix[bottom][col-1], temp
		}

		// 左边：row--,col=0
		//  7
		// (2,0) -> (1,0)
		// temp, matrix[1][0] = matrix[1][0], temp
		for row := bottom; row > top; row-- {
			temp, matrix[row-1][left] = matrix[row-1][left], temp
		}

		// 一圈结束之后，缩圈
		top++
		bottom--
		left++
		right--
	}

	return matrix
}
