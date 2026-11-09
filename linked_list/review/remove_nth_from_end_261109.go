package review

/*
1 快慢指针
2 删除慢指针的下一个节点
3 n = 链表长度，要删除的是head.next

1 2 3 4 5
n=2

slow    fast
1    2   3      n=3，要删除的是head.next
*/
func removeNthFromEnd261109(head *ListNode, n int) *ListNode {

	slow, fast := head, head

	for ; n > 0; n-- {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	// todo 移动指针
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}
