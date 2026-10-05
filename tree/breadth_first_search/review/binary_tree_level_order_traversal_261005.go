package review

/*
     3
    / \
   9   20
      /  \
     15   7

[
  [3],
  [9, 20],
  [15, 7]
]
*/

/*
queue

level

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	// todo root 需要判空，否则 queue=1，node=nil
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := []int{}
		// 出队
		currSize := len(queue)
		for i := 0; i < currSize; i++ {
			node := queue[0]
			level = append(level, node.Val)
			queue = queue[1:]
			// 左右子树入队 todo 需要判nil
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
