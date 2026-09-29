package review

/*
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

1 head是什么
head -> 1 → 2 → 3 → nil ❌   todo head 不是一个节点，是存的1的地址

2 删除节点需要找到前一个节点
因为单链表只能往后走，不能往前走。走到要删除的节点，就不知道前一个节点是谁。
所以链表删除题经常不是找到要删除的节点，而是找到要删除节点的前一个节点

3 todo 为什么需要dummy
普通节点都有前驱，但如果删除的是 head 呢？
head
 ↓
 1 → 2 → 3 → 4 → 5
n = 5，要删除 1，节点 1 前面没有节点

不是“必须”，而是为了消除 head 的特殊情况，让所有删除操作统一
if 要删除的是head {
    return head.Next
}

// 否则找到待删除节点的前驱
prev.Next = prev.Next.Next
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// todo 让第一个节点也有前驱
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	// todo 不是要返回删除的节点 return result ❌
	return dummy.Next
}
