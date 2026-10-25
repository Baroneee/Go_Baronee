package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	s2 := slow.Next
	slow.Next = nil
	s2 = reverseList(s2)

	for s2 != nil {
		n1 := head.Next
		n2 := s2.Next

		head.Next = s2
		s2.Next = n1

		head = n1
		s2 = n2
	}
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next

		cur.Next = prev
		prev = cur
		cur = next
	}
	head = prev
	return head
}

func main() {
	// Example usage
}