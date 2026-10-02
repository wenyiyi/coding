package depth_first_search

/*
https://leetcode.com/problems/diameter-of-binary-tree/description/
Given the root of a binary tree, return the length of the diameter直径 of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.
The length of a path between two nodes is represented by the number of edges between them.

        1
       / \
      2   3
     / \
    4   5
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3]
*/

/*
1 当前树的最大深度 = 左子树和右子树最大深度的较大值 + 当前节点这一层

	    root
	   /    \
	left    right

左边深度 = dfs(root.Left)
右边深度 = dfs(root.Right)

当前深度 = max(左边深度, 右边深度) + 1
dfs(node) = 以 node 为根的这棵树的最大深度

2 引入新的概念：直径
        1
       / \
      2   3
     /
    4
站在节点 1
leftDepth  = 2
rightDepth = 1
从左边最深的节点 4，经过节点 1，走到右边的节点 3
4 → 2 → 1 → 3
这里一共有3条边
= 2+1

经过当前节点的直径
= leftDepth + rightDepth

深度：max(left, right) + 1
直径：left + right

3.整棵树的直径一定要经过最上面的节点 1 吗
          1
         /
        2
       / \
      3   4
     /     \
    5       6
不一定
所以我们不能只在根节点算一次 leftDepth + rightDepth
而是 DFS 走到每一个节点的时候，都算
经过当前节点的直径 = leftDepth + rightDepth
然后保留目前见过的最大值
maxDiameter = max(maxDiameter, leftDepth+rightDepth)

左右都要
→ 在当前节点“横着连起来”
→ 算直径
→ left + right

只能选一边
→ 往父节点“向上汇报”
→ 算深度
→ max(left, right) + 1

如果 X 要告诉爸爸 P：
“从我 X 开始，往下只走一条路，最多有多少个节点？”

总结：算直径：左右都要；往上返回：只能选一边
直径 = left + right
深度 = max(left, right) + 1

todo 一边正常算树的高度，一边顺手更新目前见过的最大直径
三问
① dfs(node) 到底返回什么？
② node == nil 返回什么？
③ 左右子树已经帮我算好了，我当前节点要做什么？


① dfs 返回：当前树最大深度
② nil：0
③ 顺便更新直径 left + right
   然后 return max(left, right) + 1

*/

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		// 顺手记录：经过当前节点的直径
		maxDiameter = max(maxDiameter, left+right)
		// 正常返回高度
		return max(left, right) + 1
	}
	dfs(root)
	return maxDiameter
}
