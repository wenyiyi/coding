package linked_list

/*
You are given the head of a sorted linked list.

Delete all nodes that have duplicate numbers,
leaving only distinct numbers from the original list.

Return the linked list sorted as well.

Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]

*/

/*
*
head = cur：改变链表的头节点。
prev.Next = cur：改变链表中节点之间的连接。
prev = cur：只是移动指针变量。

head：最终要返回的链表头节点。
prev：记住前面最后一个要保留的节点，以便把后面的重复节点全部跳过
cur：当前正在检查的节点。
*/
func deleteDuplicates82(head *ListNode) *ListNode {
	var prev *ListNode // 因为开始时，我们还没有确定任何节点可以保留
	cur := head

	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			val := cur.Val

			// 跳过所有重复节点
			// 3 → 3 → 3 → 4
			// 每轮执行 cur = cur.Next，直到 cur 指向 4。
			// 此时 cur 是重复组后面的第一个节点，而不是最后一个重复节点。
			for cur != nil && cur.Val == val {
				cur = cur.Next
			}
			/*
				1 → 2 → 3 → 3 → 4 → 4 → 5
				    ↑   ↑
				   prev cur
				删除两个 3
				1 → 2 → 4 → 4 → 5
					↑   ↑
				   prev cur
			*/
			if prev == nil {
				head = cur
			} else {
				// 发现重复：跳过整组，prev不动，继续比较cur
				prev.Next = cur
			}
		} else {
			// 没有重复，后移一位
			prev = cur
			cur = cur.Next
		}
	}
	return head
}
