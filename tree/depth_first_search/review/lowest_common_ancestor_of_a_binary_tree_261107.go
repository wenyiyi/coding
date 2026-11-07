package review

/*

1 dfs(node) 返回有效的节点

2 base cases
node == nil   root=q  root=q
返回node

3 当前节点做什么
看左右子树是否有p或q

*/

func lowestCommonAncestor261107(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// 看左子树有没有p或q
	left := lowestCommonAncestor261107(root.Left, p, q)
	// 看右子树有没有p或q
	right := lowestCommonAncestor261107(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
