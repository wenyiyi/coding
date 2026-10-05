package review

/*
       4
      / \
     2   7
    / \ / \
   1  3 6  9

        4
       / \
      7   2
     / \ / \
    9  6 3  1

三问
① invertTree(root) 返回什么？
→ 翻转完成后的、以 root 为根的树

② root == nil 返回什么？
→ nil

③ 左右子树都处理好了，当前节点做什么？
→ 交换左右子树，然后返回 root

*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}
