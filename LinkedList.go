package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func New(q int) *LinkedList {
	head := &Node{}
	list := LinkedList{Head: head}

	cur := head
	for nodes := 1; nodes < q; nodes++ {
		node := &Node{}
		cur.Next = node
		cur = cur.Next
	}

	return &list
}

func (l *LinkedList) Add(val int) {
	node := Node{Val: val}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &node
}

func (l *LinkedList) Pop() {
	if l.Size() == 1 {
		return
	}

	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur = nil
}

func (l *LinkedList) At(pos int) int {
	if pos < 0 || pos > l.Size() {
		fmt.Println("list index out of range")
		return 0
	}

	cur := l.Head
	c := 0
	for c != pos {
		c += 1
		cur = cur.Next
	}
	return cur.Val
}

func (l *LinkedList) Size() int {
	c := 0
	cur := l.Head
	for cur != nil {
		c += 1
		cur = cur.Next
	}
	return c
}

func (l *LinkedList) DeleteFrom(pos int) {
	if pos < 0 || pos > l.Size() {
		fmt.Println("list index out of range")
		return
	}

	cur := l.Head
	c := 0
	for c != pos {
		c += 1
		cur = cur.Next
	}
	*cur = *cur.Next
}

func (l *LinkedList) UpdateAt(pos int, val int) {
	if pos < 0 || pos > l.Size() {
		fmt.Println("list index out of range")
		return
	}

	cur := l.Head
	c := 0
	for c != pos {
		c += 1
		cur = cur.Next
	}
	cur.Val = val
}

func NewFromSlice(s []int) *LinkedList {
	l := New(len(s))
	cur := l.Head
	for i := 0; i < len(s); i++ {
		cur.Val = s[i]
		cur = cur.Next
	}
	return l
}

func (l *LinkedList) Print() {
	cur := l.Head
	if cur == nil {
		fmt.Println("None")
		return
	}
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("")
}
