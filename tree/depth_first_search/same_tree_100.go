package depth_first_search

/*
https://leetcode.com/problems/same-tree/

Given the roots of two binary trees p and q, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical完全相同的,
and the nodes have the same value.

p:       1          q:       1

	 / \                 / \
	2   3               2   3

true

p:       1          q:   1

	 /                     \
	2                       2

false
*/
/*
三问：
① dfs(node) 到底返回什么？
② node == nil 返回什么？（Base case 就是递归的终止条件）
③ 左右子树已经帮我算好了，我当前节点要做什么？

1 这个递归函数返回什么
返回 p、q 两棵子树是否完全相同

2 递归什么时候停止？停止时返回什么？
q&p == nil     q=nil || p=nil

3 子树已经处理好了，当前节点做什么
不要继续钻进递归内部。假设左右子树的答案已经拿到了，当前这一层怎么利用它们得到自己的答案
左右子树是否相同已经知道
→ leftSame && rightSame
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 2 递归什么时候停
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	// 3 当前层做什么：递归比较左右子树
	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	// 1 递归函数返回什么
	return left && right
}
