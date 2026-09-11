package breadth_first_search

/*
https://leetcode.com/problems/binary-tree-level-order-traversal/description/?utm_source=chatgpt.com

# Binary Tree Level Order Traversal

Given the root of a binary tree,
return the level order traversal of its nodes' values.
(i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	// 把root放入队列
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// 记录当前 queue 的数量，固定下来
		currQueueSize := len(queue)
		// level := make([]int, size) 默认都是0，不要用
		level := []int{}

		// 一层一层处理，队列没有数据了就表示这一层处理完了
		for i := 0; i < currQueueSize; i++ {
			// 模拟出队
			node := queue[0]
			// 1: 表示从1取到最后
			queue = queue[1:]

			level = append(level, node.Val)

			// 把左右节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)

	}
	return result
}
