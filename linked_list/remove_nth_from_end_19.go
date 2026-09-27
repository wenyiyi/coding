package linked_list

/*
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]
*/
/*

todo 为什么不能走到要删除的节点？
因为单链表只能往后走，不能往前走。走到要删除的节点，就不知道前一个节点是谁。
所以链表删除题经常不是找到要删除的节点，而是找到要删除节点的前一个节点

1 一开始
fast
 ↓
dummy → 1 → 2 → 3 → 4 → 5 → nil
 ↑
slow

2 fast先走n步
          fast
            ↓
dummy → 1 → 2 → 3 → 4 → 5
 ↑
slow


3 slow和fast一起走
                    fast
                    ↓
dummy → 1 → 2 → 3 → 4 → 5
            ↑
          slow

4 todo 什么时候知道走到待删节点前面了？
不是“识别出 slow 到了待删节点前面”，而是通过 fast 的位置保证它一定在那里
fast 和 slow 之间永远保持 n=2 步的距离，fast 走到尾，slow 就在要删节点的前一个节点
不直接寻找 slow 的目标位置，而是人为制造两个指针的固定距离

dummy → 1 → 2 → 3 →  4 → 5
                ↑    ↑
              slow  删除
slow.Next = slow.Next.Next 完成删除

*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// todo 初始化要先指向 head
	var dummy = &ListNode{Next: head}
	slow := dummy
	fast := dummy

	// 1 让 fast 先走n步
	for n > 0 {
		fast = fast.Next
		n--
	}

	// 2 保持 n 步距离，一起移动
	//   fast 到最后一个节点时，slow 正好在待删除节点前面
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 3 删除 slow 后面的节点
	slow.Next = slow.Next.Next
	return dummy.Next
}
