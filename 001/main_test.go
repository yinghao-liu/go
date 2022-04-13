package main

import (
	"fmt"
	"testing"
)

func TestArray2slice(t *testing.T) {
	array2slice()
}
func TestSlice(t *testing.T) {
	var arra []string
	fmt.Printf("111 %p\n", arra)
	arra = append(arra, "123")
	fmt.Printf("222 %+v\n", arra)
	arra = append(arra, "456")
	fmt.Printf("333 %+v\n", arra)

}
func TestSlices(t *testing.T) {
	Slices()

}
func TestForRange(t *testing.T) {
	ForRange()

}
