package main

/*
// Hash Table
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodes := map[*ListNode]bool{}
	for headA != nil {
		nodes[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if nodes[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
*/

// 得到两个链表的长度，计算差值。
// 两个指针分别指向两个链表的head，让较长链表的指针先走差值步，然后同时走
// 当两个指针相遇时，即为链表的交点
// 倘若不相交，则两个指针最后都是null（相同），交点亦为null

// Two Pointers
// https://leetcode.com/problems/intersection-of-two-linked-lists/discuss/49785/Java-solution-without-knowing-the-difference-in-len!
// 两个指针分别指向两个链表的head，同时走
// 当某个指针达到tail时，指向另一链表的head。
// 假设两个链表的长度分别为la和lb，此时较长链表的指针距离tail有lb-la
// 当较长链表到达tail时，指向较短链表的head，此时较短链表的指针在较长链表已走了lb-la，即两个指针在距离相交点相同距离处，并同时走
// 当两个指针相遇（相同）时，该点即为相交点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
