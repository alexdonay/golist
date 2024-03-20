package main

type nodo struct {
	valor int
	next  *nodo
}

func newNodo(valor int) *nodo {
	return &nodo{valor: valor}
}