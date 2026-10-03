package review

/*
       1
      / \
     2   3
    / \
   4   5

三问
① dfs(node) 到底返回什么？
② node == nil 返回什么？
③ 左右子树已经帮我算好了，我当前节点要做什么？

*/

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		maxDiameter = max(left+right, maxDiameter)
		return max(left, right) + 1
	}
	dfs(root)
	return maxDiameter
}
