package main

import "fmt"

type Node struct {
	field int
	next  *Node
}

func main() {
	head := &Node{0, nil}
	curr := head
	for i := 1; i < 5; i++ {
		curr.next = &Node{i, nil}
		curr = curr.next
	}
	prt(head)
	head = revLL(head)
	prt(head)
}

func prt(head *Node) {
	var p = head
	for {
		fmt.Println(p.field)
		p = p.next
		if p == nil {
			break
		}
	}
}

func revLL2(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	p := revLL2(head.next)
	head.next.next = head
	head.next = nil
	return p
}

func revLL(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	var revHead, curNode, preNode *Node = nil, head, nil
	for curNode != nil {
		var nextNode = curNode.next
		if nextNode == nil {
			revHead = curNode
		}
		curNode.next = preNode
		preNode = curNode
		curNode = nextNode
	}
	return revHead
}
