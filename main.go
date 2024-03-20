package main

import "fmt"

func main() {
	ls := newList()
	ls.appendNodo(1)
	ls.appendNodo(2)
	ls.appendNodo(3)
	ls.printList()
	ls.deleteNodo(2)
	fmt.Println(ls.printList())
}