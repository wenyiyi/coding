package linked_list

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
Input: head = [1,1,2,3,3]
Output: [1,2,3]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	//todo 不需要两个指针
	cur := head

	// todo 少了 cur.Next != nil
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			// Skip the duplicate node.
			cur.Next = cur.Next.Next
		} else {
			// Move forward only when values differ.
			cur = cur.Next
		}
	}

	return head
}
