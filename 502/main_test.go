package main

import (
	"fmt"
	"testing"
)

/*****************************test***********************************/
func TestAbs(t *testing.T) {
	hello()
	t.Log("Log")
	t.Error("Error")
}

/*****************************benchmark***********************************/
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello()
	}
}

/*****************************example***********************************/
func Example() {
	fmt.Printf("hell")
	//Output: hello
}

/*****************************fuzz***********************************/
