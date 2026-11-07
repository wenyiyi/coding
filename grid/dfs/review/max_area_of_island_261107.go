package review

/*
1 dfs(row,col) 表示从 (row,col)开始走，计算目前的岛面试 calculates the total area

2 什么时候停止 base cases
- 越界
- 面积为0

3 dfs return 什么
当前岛面试
*/

func maxAreaOfIsland(grid [][]int) int {
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		var area = 0
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			return area
		}
		if grid[row][col] == 0 {
			return area
		}
		area++
		grid[row][col] = 0

		// 上下左右移动
		area += dfs(row-1, col)
		area += dfs(row+1, col)
		area += dfs(row, col-1)
		area += dfs(row, col+1)

		return area
	}

	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := dfs(i, j)
				maxArea = max(maxArea, area)
			}
		}
	}
	return maxArea
}
