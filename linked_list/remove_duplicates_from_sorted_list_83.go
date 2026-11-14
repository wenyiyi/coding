package linked_list

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
Input: head = [1,1,2,3,3]
Output: [1,2,3]0\
*/
/*
时间复杂度：O(n)
空间复杂度：O(1)
todo 核心：利用链表已排序的特点，只比较相邻节点。
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

	// todo 少了 cur.Next != nil，因为我们在循环中需要访问 cur.Next.Val，所以必须先确保 cur.Next != nil
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
