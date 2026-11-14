package review

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

todo 核心：2个指针，先保存next → 再反转 → 移动 prev → 移动 curr
*/
func reverseList261114(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	// todo curr != nil && next != nil ❌ 会导致最后一个节点漏掉
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
