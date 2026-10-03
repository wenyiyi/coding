package review

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
三问
① dfs(node) 到底返回什么？
② node == nil 返回什么？
③ 左右子树已经帮我算好了，我当前节点要做什么？

① dfs 返回：当前树最大深度
② nil：0
③ max(left, right) + 1
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return 1 + max(left, right)
}
