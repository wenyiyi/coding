package review

/*
3 → 2 → 0 → -4
    ↑         ↓
    └─────────┘
Output: true


1 → 2 → 3 → nil
Output: false

只要 fast 还能继续走，就让 slow 走 1 步、fast 走 2 步；走完以后再检查 slow == fast

重点：快慢指针
slow：一次走 1 步
fast：一次走 2 步 todo
有环：slow 和 fast 最终会相遇，也就是指向同一个节点
没有环：fast 走得更快，所以它会先到链表末尾

*/

func hasCycle(head *ListNode) bool {
	// todo slow, fast := head, head.Nxet ❌ head可能为nil
	slow, fast := head, head
	// todo  for slow != fast ❌  slow, fast := head, head 压根进不到循环

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
