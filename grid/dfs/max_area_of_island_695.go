package dfs

/*
https://leetcode.com/problems/max-area-of-island/

You are given an m x n binary matrix grid.
An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.)
You may assume all four edges of the grid are surrounded by water.
The area of an island is the number of cells with a value 1 in the island.
Return the maximum area of an island in grid. If there is no island, return 0.

Example 1:
Input: grid =
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.
*/
/*
dfs(row,col)干什么的
从(row,col)开始走，返回能够访问到的整座岛屿的面积

什么时候停止
1 越界
2 为0

当前节点做什么
num++，设置为0，然后继续移动

re


*/
func maxAreaOfIsland(grid [][]int) int {
	// todo 1 DFS 返回从当前格子出发，能够访问到的整座岛屿的面积
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		nums := 0
		// todo 2 终止条件：当前格子越界，面积为 0
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			return nums
		}
		if grid[row][col] == 0 {
			return nums
		}
		// 找到1个格子，计数
		nums++

		// 将当前格子标记为已访问，防止重复统计
		grid[row][col] = 0

		// 继续移动 todo 需要累加上下左右四个方向递归返回的面积
		nums += dfs(row-1, col)
		nums += dfs(row+1, col)
		nums += dfs(row, col-1)
		nums += dfs(row, col+1)

		return nums
	}

	maxNums := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 发现未访问的陆地，说明找到一座新岛屿
			if grid[i][j] == 1 {
				// 通过 DFS 计算整座岛屿的面积
				nums := dfs(i, j)
				maxNums = max(maxNums, nums)
			}

		}
	}
	return maxNums

}
