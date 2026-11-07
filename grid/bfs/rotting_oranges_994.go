package bfs

/*
https://leetcode.com/problems/rotting-oranges/description/
You are given an m x n grid where each cell can have one of three values:
0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally4个方向 adjacent to a rotten orange becomes rotten烂.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

Example 1:
Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4

Example 2:
Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.

Example 3:
Input: grid = [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
*/

/*
如果网格里一开始有多个腐烂橘子，为什么需要把它们全部作为 BFS 的起点，而不是从其中一个开始？
✅ 因为所有腐烂橘子都在同时传播腐烂

All rotten oranges spread the rot simultaneously, so we add all of them to the queue at the beginning.
所有腐烂橘子同时传播，因此一开始就把它们全部加入队列。

step 1 初始化
- 遍历整个 Grid
- 找到所有值为 2 的格子，加入队列
- 同时统计新鲜橘子（值为 1）的数量

	queue := [][2]int{
	    {0, 0},
	    {2, 2},
	}

step 2 bfs 按分钟传播
假设初始队列里有 2 个腐烂橘子,处理第 0 分钟时，这两个橘子会让相邻的新鲜橘子腐烂，并把它们加入队列。
但新加入的橘子不能在同一分钟继续传播，所以，我们需要在每一轮 BFS 开始时，记录当前队列长度

向上下左右四个方向传播，新鲜的改成烂，然后加入队列，同时改变数量
*/
func orangesRotting(grid [][]int) int {
	// Queue stores the positions of rotten oranges.
	// 队列保存腐烂橘子的坐标。
	queue := [][2]int{}

	// Count fresh oranges.
	// 统计新鲜橘子的数量。
	fresh := 0

	// 遍历网格，将所有腐烂橘子加入队列，
	// 同时统计新鲜橘子的数量。
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				fresh++
			}
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	minutes := 0
	for len(queue) > 0 && fresh > 0 {
		size := len(queue)
		// 只处理当前这一分钟的橘子。
		for size > 0 {
			rotten := queue[0]
			row := rotten[0]
			col := rotten[1]

			// 向上下左右四个方向传播

			// 上 row-1  col 不变
			if row > 0 {
				top := row - 1
				// 是新鲜橘子，改成烂
				if grid[top][col] == 1 {
					grid[top][col] = 2
					fresh--
					// 加入队列
					queue = append(queue, [2]int{top, col})
				}
			}

			// left  row 不变 col-1
			if col > 0 {
				left := col - 1
				if grid[row][left] == 1 {
					grid[row][left] = 2
					fresh--
					queue = append(queue, [2]int{row, left})
				}
			}
			// 下 row+1 col不变
			bottom := row + 1
			if bottom < len(grid) {
				if grid[bottom][col] == 1 {
					grid[bottom][col] = 2
					fresh--
					queue = append(queue, [2]int{bottom, col})
				}
			}

			// 右 col+1 row不变
			right := col + 1
			if right < len(grid[0]) {
				if grid[row][right] == 1 {
					grid[row][right] = 2
					fresh--
					queue = append(queue, [2]int{row, right})
				}
			}
			queue = queue[1:]
			size--
		}
		minutes++
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}

func orangesRotting2(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	queue := [][2]int{}
	fresh := 0

	// Initialize all rotten oranges and count fresh ones.
	// 初始化所有腐烂橘子，并统计新鲜橘子数量。
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			switch grid[r][c] {
			case 1:
				fresh++
			case 2:
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	// Four directions: up, down, left, right.
	// 四个方向：上、下、左、右。
	directions := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	minutes := 0

	for len(queue) > 0 && fresh > 0 {
		size := len(queue)

		// Process all oranges at the current BFS level.
		// 当前层的所有橘子同时传播。
		for i := 0; i < size; i++ {
			rotten := queue[0]
			queue = queue[1:]

			row, col := rotten[0], rotten[1]

			// 分别处理4个方向
			for _, dir := range directions {
				nr := row + dir[0]
				nc := col + dir[1]

				// Skip invalid positions and non-fresh oranges.
				// 跳过越界位置和非新鲜橘子。
				if nr < 0 || nr >= rows ||
					nc < 0 || nc >= cols ||
					grid[nr][nc] != 1 {
					continue
				}

				// Rot the fresh orange and enqueue it.
				// 腐烂新鲜橘子，并加入下一层队列。
				grid[nr][nc] = 2
				fresh--
				queue = append(queue, [2]int{nr, nc})
			}
		}

		minutes++
	}

	// Some fresh oranges cannot be reached.
	// 仍有无法被腐烂的新鲜橘子。
	if fresh > 0 {
		return -1
	}

	return minutes
}
