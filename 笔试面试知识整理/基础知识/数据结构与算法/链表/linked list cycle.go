package main

/*
// Hash Table
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	nodes := map[*ListNode]bool{}
	curr := head
	for curr != nil {
		if nodes[curr] {
			return true
		}
		nodes[curr] = true
		curr = curr.Next
	}
	return false
}
*/

// Two Pointers
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
