package depth_first_search

/*
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

        1
      /   \
     2     2
    / \   / \
   3   4 4   3
true

        1
      /   \
     2     2
      \     \
       3     3

false
*/

/*
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
	// todo isSymmetric 只有一棵树作为参数，它会判断“这棵树自己是否对称”。
	// 但我们真正要判断的是两棵不同的子树是否互为镜像，所以这里需要一个辅助函数 dfs
	var dfs func(left, right *TreeNode) bool
	dfs = func(left, right *TreeNode) bool {
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
		// 相信子树已经解决，然后只处理当前这一层
		return out && in
	}
	return dfs(root.Left, root.Right)
}
