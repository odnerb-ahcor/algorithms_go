package linkedList

import "fmt"

type linkedList struct {
	size int
	node nodeList
}

type nodeList struct {
	value any
	next  *nodeList
}

func NewLinkedList() *linkedList {
	return &linkedList{size: 0}
}

func newNode(value any) *nodeList {
	return &nodeList{
		value: value,
		next:  nil,
	}
}

func (l *linkedList) Insert(value any) {
	if l.size == 0 {
		l.node = *newNode(value)
	} else {
		node := &l.node
		for node.next != nil {
			node = node.next
		}
		node.next = newNode(value)
	}

	l.size++
}

func (l *linkedList) Remove(value any) {
	if l.size > 0 {
		if l.node.value == value {
			l.node = *l.node.next
			return
		}

		node := &l.node
		for node.next != nil && node.next.value != value {
			node = node.next
		}

		if node.next != nil {
			node.next = node.next.next
		}
	}
}

func (l *linkedList) PrintList() {
	if l.size == 0 {
		return
	}
	node := &l.node
	fmt.Print("[")
	for node != nil {
		fmt.Printf("%v", node.value)
		if node.next != nil {
			fmt.Print(", ")
		}
		node = node.next
	}
	fmt.Print("]\n")
}

func (l *linkedList) Size() int {
	return l.size
}
