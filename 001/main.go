package main

import "fmt"

func main() {
	letters := [5]string{"A", "B", "C", "D", "E"}
	// The cap built-in function returns the capacity of v, according to its type:
	// Array: the number of elements in v (same as len(v)).
	// Slice: the maximum length the slice can reach when resliced;
	fmt.Println("Before letters", letters, len(letters), cap(letters))

	slice1 := letters[0:3]
	fmt.Println("letters", &letters[0], &letters[1], &letters[2], &letters[3], &letters[4])
	fmt.Printf("slice1 len:%v cap:%v\n", len(slice1), cap(slice1))
	// len:3 cap:5
	fmt.Println("slice1", &slice1[0], &slice1[1], &slice1[2])
	// panic: runtime error: index out of range [3] with length 3
	// fmt.Println("slice1", &slice1[0], &slice1[1], &slice1[2], &slice1[3])

	// append之后分两种情况
	// 1.没超过cap, 返回原先依赖的数组地址
	// 2.超过cap, 返回的新的切片 实际已经是一个新的副本了，修改新切片不会修改原数组的值
	slice2 := append(slice1, "123")
	fmt.Println("slice1", &slice1[0], &slice1[1], &slice1[2])
	// panic: runtime error: index out of range [3] with length 3
	// fmt.Println("slice1", &slice1[0], &slice1[1], &slice1[2], &slice1[3])
	fmt.Printf("slice2 len:%v cap:%v\n", len(slice2), cap(slice2))
	fmt.Println("slice2", &slice2[0], &slice2[1], &slice2[2], &slice2[3])
	slice2[0] = "bb"
	slice1[0] = "cc"
	fmt.Println("after letters", letters, len(letters), cap(letters))

	// slice2 := letters[1:4]
	// slice3 := letters[4:5]

	// fmt.Println("before slice1", slice1)
	// slice1[1] = "Z"
	// fmt.Println("before slice3", slice3, len(slice3), cap(slice3), &slice3[0])
	// slice3 = append(slice3, "123")
	// fmt.Println("after slice3", slice3, len(slice3), cap(slice3), &slice3[0], &slice3[1])
	// fmt.Println("After letters", letters, len(letters), cap(letters))
	// fmt.Println("After slice1", slice1)
	// fmt.Println("after Slice2", slice2)
}
