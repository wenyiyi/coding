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
/*
root 入队

while queue 不为空：
    固定当前层 size
    创建 level

    重复 size 次：
        出队一个节点
        → 加入 level
        → 左孩子入队
        → 右孩子入队

    level 加入 result

todo Tree BFS = Queue；按层 BFS = 每层开始固定 size := len(queue)

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// todo 输出结果要按层按组，所以需要二维
	result := [][]int{}
	if root == nil {
		return result
	}
	// 创建队列，把root放入队列
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// todo 每层开始先固定 size := len(queue)，这一轮只处理 size 个节点。
		currQueueSize := len(queue)
		// level := make([]int, size) 默认都是0，不要用 todo 需要level去装每一层的元素
		level := []int{}
		// todo 处理完进入本层时记录的 currQueueSize 个节点，就表示这一层结束
		for i := 0; i < currQueueSize; i++ {
			// 取头节点
			node := queue[0]
			// todo 出队，1: 表示从1取到最后
			queue = queue[1:]
			// 加入层
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
