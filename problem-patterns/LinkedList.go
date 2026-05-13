package problempatterns

import "fmt"

type ListNode struct {
	val      int
	nextNode *ListNode
}

func printList(head *ListNode) {

	cur := head

	for cur != nil {

		fmt.Println(cur.val)
		cur = cur.nextNode
	}
}

func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	cur := head

	for cur != nil {

		next := cur.nextNode

		cur.nextNode = prev

		prev = cur
		cur = next
	}

	return prev
}
