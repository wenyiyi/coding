package review

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
*/

/*
todo 解题 4 步（不用 dummy）
1. 初始化：slow = head，fast = head。
2. fast 先走 n 步：如果 fast == nil，说明要删除头节点，返回 head.Next。
3. 双指针一起移动：直到 fast.Next == nil，此时 slow 位于待删除节点的前一个位置。
4. 删除节点：slow.Next = slow.Next.Next，返回 head。

todo
删除节点需要找到前一个节点
因为单链表只能往后走，不能往前走。走到要删除的节点，就不知道前一个节点是谁。
所以链表删除题经常不是找到要删除的节点，而是找到要删除节点的前一个节点
*/
func removeNthFromEnd261108(head *ListNode, n int) *ListNode {
	slow := head
	fast := head
	for n > 0 {
		fast = fast.Next
		n--
	}

	// 表示要删的是头节点
	// head = [1, 2, 3, 4, 5] n = 5 倒数第5就是头节点
	if fast == nil {
		return head.Next
	}

	// 直到fast走到最后的节点
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 要删除的节点
	//node := slow.Next
	//node.Next = nil
	//slow.Next = fast todo 万一n>2就不对了

	// head = [1,2,3,4,5], n = 2
	// fast = 5，slow=3，需要删的是slow的下一个节点
	// 3 todo 删除 slow 后面的节点
	slow.Next = slow.Next.Next

	return head
}
