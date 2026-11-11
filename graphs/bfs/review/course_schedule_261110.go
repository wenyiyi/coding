package review

/*
入度：有多少个箭头指向这个节点，0入度=0，1入度=1.
只有入度为 0 的课程，才表示没有尚未完成的先修要求，可以立即学习

	prerequisites := [][]int{
		{1, 0},
		{2, 1},
	}

[1,0]：0 → 1，所以课程 1 的入度加 1。
[2,1]：1 → 2，所以课程 2 的入度加 1。
0 → 2 ← 1

需要3个数据结构
todo graph[i] 存储的是：学完课程 i 后，可以继续学习哪些课程。
学完0学1：gragh[0]=[1]
学完1学2: gragh[1]=[2]
graph[0] = []int{1} // 学完 0，可以学 1
graph[1] = []int{2} // 学完 1，可以学 2
graph[2] = []int{}  // 学完 2，没有后续课程

todo indegree[course] 存的是每门课程的入度值
indegree[0] =0
indegree[1] =1
indegree[2] =1

todo queue 存的是入度值为0的课程

5 个步骤
1. 计算入度（Indegree）：每门课程有多少个先修条件。
2. 建立邻接表（Graph）：每门课程指向哪些后续课程。
3. 初始化队列（Queue）：把所有入度为 0 的课程入队。
4. BFS 遍历：每次出队一门课程，完成课程数 completed++，将它指向的节点入度减 1；如果减到 0，就加入队列。
5. 判断结果：completed == numCourses，返回 true，否则返回 false（存在环）。
*/

func canFinish261111(numCourses int, prerequisites [][]int) bool {
	// 1 indegree
	indegree := make([]int, numCourses)
	// 2 graph
	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		course := p[0]
		pre := p[1]

		indegree[course]++
		graph[pre] = append(graph[pre], course)
	}

	// 3 queue
	queue := make([]int, 0)
	// 4 indegree = 0 in queue
	for course, item := range indegree {
		if item == 0 {
			queue = append(queue, course)
		}
	}

	completed := 0
	for len(queue) > 0 {
		course := queue[0]
		completed++
		queue = queue[1:]

		for _, next := range graph[course] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	return completed == numCourses
}
