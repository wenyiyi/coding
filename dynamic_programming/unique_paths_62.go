package dynamic_programming

/*
https://leetcode.com/problems/unique-paths/

There is a robot on an m x n grid.
The robot is initially located at the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]).
The robot can only move either down or right at any point in time.

Given the two integers m and n,
return the number of possible unique paths that the robot can take to reach the bottom-right corner.
The test cases are generated so that the answer will be less than or equal to 2 * 109.

Example 1:
Input: m = 3, n = 7
Output: 28

Example 2:
Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down

Constraints:
1 <= m, n <= 100
*/

/*
todo
DP四问
第一步：怎么定义 dp
看到 DP，先问自己一句：题目最后让我求什么？把“大问题”缩小到某个位置/某个前缀，就是 dp 的含义。
题目最终问的是：从 (0,0) 到 (m-1,n-1) 有多少条路径？
这时候它自然就变成：
dp[i][j] = 从左上角(0,0)走到格子(i,j)的路径数量
(不要让：m,n = 整个问题的固定大小，而是引入 i,j = 任意一个中间状态)

第二步：我要得到 dp[i][j]，最后一步可能从哪里来
可能从左边过来
dp[i][j-1]

可能从上边过来
dp[i-1][j]

所以 dp[i][j] = dp[i-1][j] + dp[i][j-1]

第三步：初始化状态
dp[0][j]=1
dp[i][0]=1
因为第 1 行只能一直向右，
第 1 列只能一直向下

第四步：return 什么
return dp[m-1][n-1]

重点
DP 状态定义技巧：
题目最终问：
从 (0,0) 到 (m-1,n-1) 有多少种路径？

把“最终状态”泛化成“任意中间状态”：
从 (0,0) 到 (i,j) 有多少种路径？
→ 得到 dp[i][j] 的定义
*/
func uniquePaths(m int, n int) int {

	// 定义
	var dp = make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 1; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 最后一步可能从哪里来
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
