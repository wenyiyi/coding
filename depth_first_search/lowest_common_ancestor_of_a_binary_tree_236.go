package depth_first_search

/*
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
p 和 q 的公共祖先在哪里？

它不是在问：
p、q 在第几层？

而更像是在问：
p 在不在我的左/右子树？q 在不在我的左/右子树？
所以这题更自然的是 DFS

todo 层 → BFS；子树 / 路径 / 祖先 → 先想 DFS

三问
1 函数返回值是什么
lowestCommonAncestor：在以 root 为根的这棵子树中，找到的 p / q / 它们的最近公共祖先

2 递归终止条件

	if root == nil || root == p || root == q {
	    return root
	}

3 左右节点都处理完，当前节点怎么处理
left := lowestCommonAncestor(root.Left, p, q)
right := lowestCommonAncestor(root.Right, p, q)
左子树找到了东西，右子树也找到了东西
左右两边都有发现 → 当前节点就是公共祖先

	   3   ← 当前 root
	  / \
	 5   1
	↑     ↑
	p     q

现在你站在 3
你让左子树帮你找：left = 5，表示 左边找到了 p=5
再让右子树帮你找：right = 1，表示 右边找到了 q=1

所以现在局面是：

	     3
	    / \
	   ↓   ↓
	找到p  找到q
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 两边都有发现，表示 p，q 分布在两边，root是汇合点
	if left != nil && right != nil {
		return root
	}

	// 左子树有“有效发现”，右子树什么都没找到，所以把左边的结果继续往上传
	//         root
	//       /
	//      p
	//  q 不在这棵子树里，此时 left = p      right = nil
	if left != nil {
		return left
	}

	return right
}
