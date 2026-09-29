package review

/*
list1: 1 → 2 → 4
list2: 1 → 3 → 4

result:
1 → 1 → 2 → 3 → 4 → 4


dummy -> 1 -> 1 -> 2
curr

分别遍历 list1 和 list2
对比节点

todo
dummy = 永远不动，只负责记住结果链表的起点
curr.Next = list1      // 当前结果链表接节点
list1 = list1.Next     // 原链表前进
curr = curr.Next       // 结果链表尾巴前进
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			// todo curr = list1 ❌ 是接上节点，不是指向
			curr.Next = list1
			list1 = list1.Next // 原链表前进
			curr = curr.Next   // 结果链表前进
		} else {
			curr.Next = list2
			list2 = list2.Next
			curr = curr.Next
		}
	}
	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}
	return dummy.Next
}
