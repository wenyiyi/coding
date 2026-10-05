package dfs

/*
https://leetcode.com/problems/number-of-islands/

Given an m x n 2D binary grid  which represents a map of '1's (land) and '0's (water),
return the number of islands.

An island is surrounded by water and is formed by connecting adjacent邻近的 lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.

*/
/*
DP：
“这个格子的答案，怎么从之前的状态推出来？”
→ 状态转移

Grid DFS：
“从这个格子出发，我能走到哪些相连的格子？”
→ 搜索 / 遍历

Tree DFS 是沿着 Left / Right 走。
Grid DFS 是沿着 上 / 下 / 左 / 右走。

Tree DFS                Grid DFS

	    node                   (i,j)
	   /    \                ↑
	left   right          ← (i,j) →
	                          ↓
*/
/*
【1】1 0 0
  1 1 0 0
  0 0 1 0
  0 0 0 1
从 (row, col) 出发，四个下一步分别是什么坐标
上：(row-1, col)
下：(row+1, col)
左：(row, col-1)
右：(row, col+1)

Tree：递归 2 个方向
      Left / Right

Grid：递归 4 个方向
      上 / 下 / 左 / 右

终止条件
row < 0
row >= len(grid)
col < 0
col >= len(grid[0])

走到水 0 也停止
grid[row][col] == '0'

todo 因此 Grid DFS 很经典的 Base Case 是：
if row < 0 || row >= len(grid) ||
   col < 0 || col >= len(grid[0]) ||
   grid[row][col] == '0' {
    return
}

todo 注意这里不是 return 0，因为我们这个 DFS 暂时不需要计算数字，它只是负责：
从一个陆地出发，把整个相连的岛走一遍。

现在有一个关键问题
1 1
1 1

A → B
↑   ↓
D ← C

DFS 可以一直绕回来。
怎么避免同一个 1 被反复访问，造成死循环？
- 访问到一个 '1' 后，立刻把它改成 '0'


Tree DFS 三问                 Grid DFS 三问

① 返回什么？                 ① dfs负责什么？
② 什么时候停？               ② 什么时候停？
③ 当前节点做什么？            ③ 当前格做什么 + 往哪走？


第一问：dfs(row, col) 是干什么的？
从 (row, col) 出发，把与它相连的整座岛全部访问一遍，所以它其实不需要返回值
dfs(row, col) 只负责一件事：
把从这个位置连接出去的整座岛淹掉
里面不负责 num++

第二问：什么时候停止？
① 越界了
row < 0 || row >= len(grid)
col < 0 || col >= len(grid[0])
② 这个位置不是需要继续访问的陆地
grid[row][col] == '0'

第三问：当前这一格做什么？
当前 1 → 标记成 0
然后 DFS 上下左右
grid[row][col] = '0'

        上
        ↑
左 ← 当前格 → 右
        ↓
        下


todo 外层扫描负责“数岛”，DFS 负责“淹岛”
*/
func numIslands(grid [][]byte) int {

	var dfs func(row, col int)
	// 从当前格子出发，把与它相连的整个岛访问掉
	dfs = func(row, col int) {
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			return
		}
		if grid[row][col] == '0' {
			return
		}
		grid[row][col] = '0'
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}
	num := 0

	// 外层双循环：负责发现“新岛”
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			// 发现一座新的、还没访问过的岛
			if grid[row][col] == '1' {
				// num++
				num++
				// dfs 把整座岛淹掉
				dfs(row, col)
			}
		}
	}

	// todo 返回岛的数量，不是陆地格子的数量
	return num
}
