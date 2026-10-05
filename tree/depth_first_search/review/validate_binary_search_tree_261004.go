package review

import "math"

/*
给你一棵二叉树 root，判断它是否是一棵有效的二叉搜索树（BST）。
有效 BST 满足：
- 节点的左子树只包含严格小于当前节点的值。
- 节点的右子树只包含严格大于当前节点的值。
- 左右子树本身也必须是 BST。
    2
   / \
  1   3

输出：true

      5
     / \
    1   4
       / \
      3   6

输出：false
*/

/*
三问
1 dfs 函数返回值
当前节点是否是合法的bst
2 递归结束条件
nil：true
3 左右子树都满足条件，当前节点做什么
比较值

todo
只有 nil 可以直接 true。40-0
BST 是开区间 (lower, upper)，等于边界也不行
*/
func isValidBST261004(root *TreeNode) bool {
	var dfs func(node *TreeNode, lower, upper int64) bool
	dfs = func(node *TreeNode, lower, upper int64) bool {
		if node == nil {
			return true
		}
		// if node.Left == nil && node.Right == nil {
		//			return true
		//		} ❌ 叶子节点左右为nil，但有可能val不符合条件

		// todo bst也不能=
		if int64(node.Val) <= lower || int64(node.Val) >= upper {
			return false
		}
		left := dfs(node.Left, lower, int64(node.Val))
		right := dfs(node.Right, int64(node.Val), upper)
		return left && right
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}
