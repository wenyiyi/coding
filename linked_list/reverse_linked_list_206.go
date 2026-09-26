package linked_list

/*
https://leetcode.com/problems/reverse-linked-list/
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]
*/

/*

prev        curr       next
 ↓           ↓          ↓
nil          1    →     2 → 3


① 保存后面
next := curr.Next
nil ← 1       2 → 3
      ↑       ↑
     curr    next

② 当前指针反转
curr.Next = prev

③ prev 前进
prev = curr

④ curr 前进
curr = next

              prev       curr
               ↓           ↓
nil ← 1 ← 2 ← 3           nil

结束循环，表头是prev

todo 改 .Next 之前，先保存原来的 .Next
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// prev = 已经反转好的前半段
// curr = 当前正在处理的节点
// next = 暂时保存后面的节点
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		// 1 保存后面
		next := curr.Next
		// 2 当前指针反转
		curr.Next = prev
		// 3 prev 往前
		prev = curr
		// 4 curr 往前
		curr = next
	}

	return prev
}
