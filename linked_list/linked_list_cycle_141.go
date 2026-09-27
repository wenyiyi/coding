package linked_list

/*
https://leetcode.com/problems/linked-list-cycle/

Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote表示 the index of the node that tail's next pointer is connected to.
Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.


Example 1:
1 → 2 → 3 → 4
    ↑       ↓
    ← ← ← ← ←
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

Example 2:
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
*/
/*
重点：快慢指针

slow：一次走 1 步
fast：一次走 2 步
有环：slow 和 fast 最终会相遇，也就是指向同一个节点
没有环：fast 走得更快，所以它会先到链表末尾


fast != nil && fast.Next != nil
          ↓
slow 走 1 步
fast 走 2 步
          ↓
slow == fast ?
是 → 有环
          ↓
循环正常结束 → 无环

*/
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
