package depth_first_search

/*
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

Given a binary tree, find the lowest common ancestor祖先 (LCA) of two given nodes in the tree.
According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes p and q
as the lowest node in T that has both p and q as descendants后代
(where we allow a node to be a descendant of itself).”

          3
        /   \
       5     1
      / \   / \
     6   2 0   8
        / \
       7   4

p = 5
q = 4

答案：5，因为节点也可以是自己的祖先
*/

/*
       3
      / \
     5   1
    / \
   6   2

我们要找：6 和 2 最近的公共祖先是谁
5，因为从 5 往下，既能找到 6，也能找到 2，而且 5 是离它们最近的共同祖

假设我们站在节点 5
我们让左子树帮我们找 p=6 或 q=2
左边搜索：
5.Left → 6
发现：
6 == p
所以左边可以把 节点 6 返回给 5：
       5
      / \
   返回6   ?

todo 第一个关键点
这个递归函数不是返回 true/false
它返回的是：
“我这棵子树里找到的有效节点。”

同理右子树返回 2

所以现在站在节点 5，它收到了两个结果：
       5
      / \
     6   2
     ↑   ↑
   返回6 返回2
左边：找到了一个目标
右边：也找到了一个目标

这时候 5 就知道：
一个目标在我左边，一个目标在我右边。

所以它们第一次在我这里“汇合”。
因此：
5 就是最近公共祖先

于是节点 5 这一层应该：
return root // 也就是返回 5

往上走来到3（todo 左边的5也要往上传）
        3
       / \
    返回5  1
     / \
    6   2
右边的 1 下面既没有 6，也没有 2，所以右边最终返回 nil


        3
       /
      5  ← p
     /
    6    ← q
但 5 本身就是 p，而且 6 是 5 的后代，所以 5 本身就是最近公共祖先


① 返回什么？
   → 子树里找到的有效节点。因为有时候这棵子树里只找到了 p，它也会返回 p

② 什么时候停？
   → nil / p / q 直接返回

③ 当前层干什么？
   → 两边都有：返回自己
   → 只有一边：返回那一边
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	// 因为题目保证 p 和 q 都存在于树中, 那么不管 q 是在 p 的下面，还是在其他地方，这一支都可以直接把 p 返回上去
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 两边都有发现，表示 p，q 分布在两边，root是汇合点
	if left != nil && right != nil {
		return root
	}

	// 左子树有“有效发现”，右子树什么都没找到，
	// 所以把左边的结果继续往上传
	if left != nil {
		return left
	}

	return right
}
