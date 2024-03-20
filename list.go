package main

import "fmt"

type list struct {
	head *nodo
}

func newList() *list {
	return &list{}
}

func (l *list) appendNodo(N int) {
	nodo := newNodo(N)
	if l.head == nil {
		l.head = nodo
		return
	} else {
		last := l.head
		for last.next != nil {
			last = last.next
		}
		last.next = nodo
	}
}
func (l *list) printList()string {
	str := ""
	for n := l.head; n != nil; n = n.next {
		divisor := ""
		if n.next !=nil{
			 divisor = ","
			}
		str += fmt.Sprint(n.valor) + divisor
	}
	return "[" + str + "]"
}
func (l *list) deleteNodo(N int) {
	if l.head == nil {
		return
	} else {
		for n := l.head; n != nil; n = n.next {
			if n.next.valor == N {
				n.next = n.next.next
				return
			}
		}
	}
}