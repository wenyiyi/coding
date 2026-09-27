package linked_list

/*
https://leetcode.com/problems/merge-two-sorted-lists/

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list.
The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
*/

/*
一开始
list1:
1 → 2 → 4
↑
list1

list2:
1 → 3 → 4
↑
list2

结果链表：
dummy

	 ↓
	[0]
	 ↑
	curr

比较完 1
结果链表：
dummy → 1

	 ↑
	curr

list1

	↓

2 → 4

list2

	↓

1 → 3 → 4

dummy = 永远不动，记住结果链表的起点
curr  = 结果链表的尾巴

	“下一个节点接到我后面”

list1 = 链表1当前还没处理的节点
list2 = 链表2当前还没处理的节点

list1 当前的人 ─┐

	├─ 比一下谁小 → 接到 curr 后面

list2 当前的人 ─┘
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for list1 != nil && list2 != nil {
		// next := cur.Next cur.Next 是nil，LC21 并不是在反转 cur.Next，而是在给结果链表接一个新节点，接完以后直接 cur = cur.Next
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		// todo ← 每接一个，结果链表的尾巴往前走。接节点之后，检查“施工指针”有没有跟着移动
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	// return dummy   dummy 是我们故意创建的假节点，真正的结果从下一个节点开始
	return dummy.Next
}

/*
dummy → 1 → 1 → 2 → 3 → 4 → 4
 ↑
假的

dummy → 1 → 1 → 2 → 3 → 4 → 4
        ↑
     dummy.Next
*/
