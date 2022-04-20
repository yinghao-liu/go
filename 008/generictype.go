package main

import "fmt"

type signed interface {
	int | int16
}

type calculate[T signed] struct {
	a T
	b T
}

func (c calculate[T]) add() T {
	return c.a + c.b
}

func GenericType() {
	var cal calculate[int]
	cal.add()
}

/****************************** genericc interface **********************************/
type signer[T signed] interface {
	Notify(T)
}

type signee[T signed] struct {
}

func (s signee[T]) Notify(a T) {
	fmt.Printf("[%T]: %+v\n", a, a)
	return
}
func GenericInterface() {
	sr := new(signee[int])
	sr.Notify(1)
}
