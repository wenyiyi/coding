package review

/*
        1
      /   \
     2     2
    / \   / \
   3   4 4   3
输出：true

三问
1 dfs 返回什么
左右节点完全对称
2 递归结束条件
左右节点为nil
3 左右子树都符合条件，当前节点做什么
对比两个节点

① dfs(left, right) 返回什么？
→ 两棵子树是否互为镜像

② 什么时候停？
→ 都 nil：true
→ 一个 nil：false
→ 值不同：false

③ 当前层怎么处理？
→ 外侧和外侧比：left.Left  ↔ right.Right
→ 内侧和内侧比：left.Right ↔ right.Left
→ 两边都 true 才是镜像
*/

func isSymmetric(root *TreeNode) bool {

	var dfs func(left *TreeNode, right *TreeNode) bool
	dfs = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		out := dfs(left.Left, right.Right)
		in := dfs(left.Right, right.Left)
		return out && in
	}
	return dfs(root.Left, root.Right)
}
