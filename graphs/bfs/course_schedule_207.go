package bfs

/*
https://leetcode.com/problems/course-schedule/description/

There are a total of numCourses courses you have to take,
labeled from 0 to numCourses - 1.
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates
that you must take course bi first if you want to take course ai.
For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.



Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.

Example 2:
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
*/

/*
[0, 1]  1->0  必须先完成课程 1，才能学习课程 0

入度：有多少个箭头指向这个节点，0入度=0，1入度=1.
只有入度为 0 的课程，才表示没有尚未完成的先修要求，可以立即学习

indegree ：= []int{0,1}

prerequisites := [][]int{
    {1, 0},
    {2, 1},
}
[1,0]：0 → 1，所以课程 1 的入度加 1。
[2,1]：1 → 2，所以课程 2 的入度加 1。


0 → 2 ← 1
graph[0] = []int{2} 0->2
graph[1] = []int{2} 1->2
graph[2] = []int{}  2 没有指向任何课程

indegree = [0, 0, 2]
queue = [0, 1]


5 个步骤
1. 计算入度（Indegree）：每门课程有多少个先修条件。
2. 建立邻接表（Graph）：每门课程指向哪些后续课程。
3. 初始化队列（Queue）：把所有入度为 0 的课程入队。
4. BFS 遍历：每次出队一门课程，完成课程数 completed++，将它指向的节点入度减 1；如果减到 0，就加入队列。
5. 判断结果：completed == numCourses，返回 true，否则返回 false（存在环）。

*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1. Build graph and indegree array.
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	// 计算入度   [0, 1]  1 ——> 0
	for _, p := range prerequisites {
		course, pre := p[0], p[1]
		// 建立邻接表
		graph[pre] = append(graph[pre], course)
		indegree[course]++
	}

	// 2. Enqueue all courses with zero indegree.
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 3. BFS topological sort.
	completed := 0

	for len(queue) > 0 {
		// 表示当前课程已经完成
		cur := queue[0]
		queue = queue[1:]
		// 完成的课程数++
		completed++

		// Reduce the indegree of dependent courses.
		// todo  graph[0] = []int{2} 0->2 所以 graph[cur] = 指向的下一个节点
		for _, next := range graph[cur] {
			indegree[next]--

			// 入度为0加入队列
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// 4. Check whether all courses are completed.
	return completed == numCourses
}
