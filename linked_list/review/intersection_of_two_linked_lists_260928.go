package review

/*
A:
4 → 1  ─┐

		    ↓
	        8 → 4 → 5
		    ↑

5 → 6 → 1
B:

返回相交的节点8

A 走到尾，然后指向B
B 走到尾，然后指向A

PA = PB 就是相交节点
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB

	for pA != pB {
		// todo pA.Next == nil ❌
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}
