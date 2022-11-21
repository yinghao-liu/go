package set

import "fmt"

var exist struct{}

type void struct{}

func New(items ...interface{}) Set {
	var s Set
	s.m = make(map[interface{}]void)
	s.Add(items...)
	return s
}

type Set struct {
	m map[interface{}]void
}

func (s *Set) Add(items ...interface{}) {
	for _, item := range items {
		s.m[item] = exist
	}
	return
}
func (s *Set) Del(item interface{}) {
	delete(s.m, item)
	return
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.m = make(map[interface{}]void)
}
func (s *Set) ShowAll() {
	fmt.Printf("%+v\n", s.m)
}
