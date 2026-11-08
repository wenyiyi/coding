package review

/*
1 初始化，找出腐烂的橘子，统计新鲜橘子的数量

2 腐烂的橘子放入队列

3 出队，4个方向移动，校验坐标边界，新鲜的改成腐烂，数量变更

4 新腐烂的橘子放入队列

5 最终还有新鲜的橘子，返回-1


*/

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	fresh := 0
	// [[0,0],[1,1]]
	queue := make([][]int, 0)

	// 初始化
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				fresh++
			}
			if grid[i][j] == 2 {
				// 腐烂的橘子入队
				queue = append(queue, []int{i, j})
			}
		}
	}

	minute := 0
	// todo 四个方向，通过+1，-1来识别，用 {} 而不是 []
	directions := [][2]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}
	for len(queue) > 0 && fresh > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			// 腐烂的橘子出队
			rotten := queue[0]
			queue = queue[1:]

			// 向4个方向扩展
			for _, dir := range directions {
				newRow := rotten[0] + dir[0]
				newCol := rotten[1] + dir[1]

				// todo 检查坐标是否越界
				if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
					continue
				}

				// 新鲜的改成腐烂
				if grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2
					fresh--
					// todo 加入队列
					queue = append(queue, []int{newRow, newCol})
				}
			}
		}
		minute++
	}

	// todo 如果还有新鲜的，表示无法腐烂到
	if fresh > 0 {
		return -1
	}

	return minute
}
