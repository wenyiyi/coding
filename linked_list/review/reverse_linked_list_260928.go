package review

/*
1 → 2 → 3 → 4 → 5 → nil

↓

5 → 4 → 3 → 2 → 1 → nil

prev 应该指向哪里？
原链表：
nil     1 → 2 → 3
↑      ↑

prev   curr

2 先保存下来
curr -> prev

curr 移动
prev 移动

最后

	  prev
		↓
		3   →  2 → 1 → nil

curr

	↓

nil
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// todo prev := head ❌
	var prev *ListNode
	curr := head

	// todo curr.Next != nil ❌
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	// todo curr ❌
	return prev
}
