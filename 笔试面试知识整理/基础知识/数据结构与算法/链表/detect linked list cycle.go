package main

/*
// Hash Table
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	nodes := map[*ListNode]bool{}
	curr := head
	for curr != nil {
		if nodes[curr] {
			return curr
		}
		nodes[curr] = true
		curr = curr.Next
	}
	return nil
}
*/

// Two Pointers
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	entry := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != entry {
				slow = slow.Next
				entry = entry.Next
			}
			return entry
		}
	}
	return nil
}
