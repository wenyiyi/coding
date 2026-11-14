package review

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
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			// cur = cur.Next ❌，不需要，因为cur还要跟新的cur.Next对比
		} else {
			cur = cur.Next
		}
	}
	return head
}
