package dynamic_programming

/*
https://leetcode.com/problems/minimum-number-of-work-sessions-to-finish-the-tasks/description/

There are n tasks assigned to you. The task times are represented as an integer array tasks of length n,
where the ith task takes tasks[i] hours to finish.
A work session is when you work for at most sessionTime consecutive hours and then take a break.

You should finish the given tasks in a way that satisfies the following conditions:

If you start a task in a work session, you must complete it in the same work session.
You can start a new task immediately after finishing the previous one.
You may complete the tasks in any order.

Given tasks and sessionTime, return the minimum number of work sessions needed to finish all the tasks following the conditions above.
The tests are generated such that sessionTime is greater than or equal to the maximum element in tasks[i].



Example 1:
Input: tasks = [1,2,3], sessionTime = 3
Output: 2
Explanation: You can finish the tasks in two work sessions.
- First work session: finish the first and the second tasks in 1 + 2 = 3 hours.
- Second work session: finish the third task in 3 hours.

Example 2:
Input: tasks = [3,1,3,1,1], sessionTime = 8
Output: 2
Explanation: You can finish the tasks in two work sessions.
- First work session: finish all the tasks except the last one in 3 + 1 + 3 + 1 = 8 hours.
- Second work session: finish the last task in 1 hour.

Example 3:
Input: tasks = [1,2,3,4,5], sessionTime = 15
Output: 1
Explanation: You can finish all the tasks in one work session.


Constraints:

n == tasks.length
1 <= n <= 14
1 <= tasks[i] <= 10
max(tasks[i]) <= sessionTime <= 15
*/

/*
mask = 已经完成了哪些任务


*/

func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)
	// 把数字1向左移动n位，左移1位相当于*2，1 << n = 2^n
	size := 1 << n

	// dp[mask][0] = 完成 mask 里的任务最少需要多少个 session
	// dp[mask][1] = 最后一个 session 已经使用了多少时间
	//
	// 例如：
	// dp[011] = [2]int{1, 3}
	// 表示 task0 和 task1 已经完成
	// 一共用了 1 个 session
	// 当前这个 session 已经用了 3 小时
	//              第0列       第1列
	//              ↓           ↓
	//dp[0]      sessions      used
	//dp[1]      sessions      used
	//dp[2]      sessions      used
	//...
	//dp[size-1] sessions      used
	// 创建 DP 数组，size 行、2 列
	dp := make([][2]int, size)

	// 初始化成很大的值
	for mask := 0; mask < size; mask++ {
		dp[mask] = [2]int{n + 1, sessionTime + 1}
	}

	// 一个空 session，当前使用时间为 0
	dp[0] = [2]int{1, 0}

	for mask := 0; mask < size; mask++ {
		for i := 0; i < n; i++ {
			// task i 已经完成
			if mask&(1<<i) != 0 {
				continue
			}
			// 把 task i 加进去
			nextMask := mask | (1 << i)
			sessions := dp[mask][0]
			used := dp[mask][1]

			var nextSessions int
			var nextUsed int

			// 当前 session 放得下
			if used+tasks[i] <= sessionTime {
				nextSessions = sessions
				nextUsed = used + tasks[i]
			} else {
				// 新开 session
				nextSessions = sessions + 1
				nextUsed = tasks[i]
			}

			// sessions 是答案，used 是为了保证这个答案能正确算出来的辅助状态
			// 找 session 数量更少；如果 session 数量相同，就找当前 session 用时更少的方案
			if nextSessions < dp[nextMask][0] || (nextSessions == dp[nextMask][0] && nextUsed < dp[nextMask][1]) {
				dp[nextMask] = [2]int{nextSessions, nextUsed}
			}
		}
	}

	return dp[size-1][0]
}
