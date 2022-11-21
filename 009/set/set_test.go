package set_test

import (
	"fmt"
	"set/set"
	"testing"
)

func TestSet(t *testing.T) {
	s := set.New("a", "b", 1)
	fmt.Printf("s is contains:%v\n", s.Contains("s"))
	s.ShowAll()

}
