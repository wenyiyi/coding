package review

/*
给定两棵二叉树 p 和 q，判断它们是否完全相同。
两棵树完全相同需要：
- 结构相同
- 对应节点值相同

p:       1          q:       1
        / \                 / \
       2   3               2   3
输出：true

p:       1          q:       1
        /                     \
       2                       2
输出：false
*/
/*
三问：
1 dfs 返回什么

2 递归结束条件
3 左右子树都处理好了，当前节点怎么做
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)
	return left && right
}
