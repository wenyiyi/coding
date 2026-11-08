package dfs

/*
https://leetcode.com/problems/clone-graph/description/
Given a reference of a node in a connected undirected graph.
Return a deep copy (clone) of the graph.
Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}


Test case format:

Example 1:
For simplicity, each node's value is the same as the node's index (1-indexed).
For example, the first node with val == 1, the second node with val == 2, and so on.
The graph is represented in the test case using an adjacency list.

An adjacency list(邻接表) is a collection of unordered lists used to represent a finite graph.
Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1.
You must return the copy of the given node as a reference to the cloned graph.


Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]

数组位置	   节点	邻居	    含义
adjList[0]	1	[2,4]	节点 1 连接 2、4
adjList[1]	2	[1,3]	节点 2 连接 1、3
adjList[2]	3	[2,4]	节点 3 连接 2、4
adjList[3]	4	[1,3]	节点 4 连接 1、3

Explanation: There are 4 nodes in the graph.
1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
4th node (val = 4)'s neighbors are 1st node (val = 1) and

输入输出一样，因为题目要求 deep cope，深拷贝
- 节点的值不变
- 节点之间的连接关系不变
- 所有的节点都是新创建的对象，内存地址不变
*/

/*
1 用 map 来存访问过的节点
visted := make(map[*Node]*Node)
位置	    类型	    作用
Key	    *Node	原图中的节点
Value	*Node	新创建的克隆节点

2 先克隆节点，放入 map，再克隆邻居
DFS 的执行过程：
1. 访问节点 1，创建 clone1，存入 Map。
2. 访问节点 2，创建 clone2，存入 Map。
3. 访问节点 3，创建 clone3，存入 Map。
4. 再次遇到节点 1，发现 Map 中已经有 clone1，直接返回它。

3 dfs 4步
- DFS 干什么：输入一个原图节点，返回这个节点对应的克隆节点，并确保它的邻居关系也被克隆
- 终止条件 base cases：node==nil  ｜｜ 节点已经被克隆过
- 单层递归逻辑？ 创建节点 → 存 Map → 递归克隆邻居并建立连接 → 返回新节点
- 返回值：返回当前节点对应的克隆节点的指针
*/

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	var dfs func(*Node) *Node
	dfs = func(cur *Node) *Node {
		// 1 Handle nil.
		if cur == nil {
			// 表示原节点不存在，返回nil
			return nil
		}
		// 2 Return existing clone if visited.
		if _, ok := visited[cur]; ok {
			// todo 不能返回nil
			// 因为 visited[cur] 保存的是已经创建好的克隆节点。
			//如果返回 nil，就相当于告诉上一层 DFS：“这个邻居不存在”，导致克隆图的连接关系丢失。
			return visited[cur]
		}
		// 3: Create a new node.
		newNode := &Node{Val: cur.Val, Neighbors: make([]*Node, 0)}
		// 4: Store original -> clone in visited.
		visited[cur] = newNode
		// 5: Recursively递归地 clone neighbors.
		for _, neighbor := range cur.Neighbors {
			// 返回当前节点对应的克隆节点的指针
			clonedNeighbor := dfs(neighbor)
			// todo 加入 map
			newNode.Neighbors = append(newNode.Neighbors, clonedNeighbor)
		}
		// 6: Return cloned node
		return newNode
	}
	return dfs(node)
}
