package depth_first_search

/*
https://leetcode.com/problems/path-sum/description/
Given the root of a binary tree and an integer targetSum,
return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

A leaf is a node with no children.

        5
       / \
      4   8
     /
    11
   /  \
  7    2

Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
5 → 4 → 11 → 2 = 22

Output: true
Explanation: The root-to-leaf path with the target sum is shown.

*/

/*
三问
① 返回什么？
   当前 root 出发，是否存在和为 targetSum 的 root→leaf 路径
② 什么时候停？
   root == nil → false
   到叶子 → 判断 root.Val == targetSum todo Path Sum 的终点是 leaf，不是“中途刚好凑够 targetSum”
③ 当前层做什么？
   targetSum 减掉当前 root.Val
   左右任意一边成功即可
   → left || right


题目给我的函数，能不能直接表达“子问题”？
能 → 直接递归原函数。
不能，需要额外参数/额外返回含义/共享状态 → 单独写 dfs
LC104  maxDepth(root)          → 直接递归 ✅
LC100  isSameTree(p,q)         → 直接递归 ✅
LC236  LCA(root,p,q)           → 直接递归 ✅
LC112  hasPathSum(root,target) → 直接递归 ✅

LC98   还需要 lower/upper      → dfs
LC543  dfs要返回高度，
       外层要返回直径           → dfs
LC101  需要比较 left/right，
       原函数只有 root          → dfs

*/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 叶子结点 todo 因为题目要求路径必须从 root 一直走到叶子节点
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	targetSum -= root.Val
	left := hasPathSum(root.Left, targetSum)
	right := hasPathSum(root.Right, targetSum)

	return left || right
}
