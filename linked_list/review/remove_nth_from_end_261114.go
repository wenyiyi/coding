package review

/*
给定一个链表的头节点 head 和整数 n，删除链表的倒数第 n 个节点，并返回删除后的链表头节点

1 slow fast 初始化
2 移动fast
3 删除的节点是slow的下一个节点
4 todo  n = 链表长度，要删除的是head.next，此时fast=nil

	s   f

1 2 3 4 5

n=2，要删除4
*/
func removeNthFromEnd261114(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for ; n > 0; n-- {
		fast = fast.Next
	}

	// todo n = 链表长度，fast走到最后
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 删除节点
	slow.Next = slow.Next.Next

	return head
}
