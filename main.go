package main

import "fmt"

func main() {
	l := LinkedList{}
	l.Head = &Node{Val: 1}
	fmt.Println(l.Head.Val)
	l.UpdateAt(0, 15)
	fmt.Println(l.Head.Val)
	l.Add(8)
	fmt.Println(l.At(1))
	l.Add(3)
	fmt.Println(l.At(2))
	l.DeleteFrom(1)
	fmt.Println(l.At(1))
	l.Pop()
	fmt.Println(l.At(1))
	l = *NewFromSlice([]int{1, 2, 3, 4, 5})
	l.Print()
	l.DeleteFrom(2)
	l.Print()
}
