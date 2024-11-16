package main

import "fmt"

func main() {
	// println("dede")
	// fmt.Println("deepak")

	// array learning
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)
	fmt.Println(arr1[1])
	// fmt.Println(arr1[4]) throws error - index out of bound

	arr2 := [3]int{5, 6, 7}
	fmt.Println(arr2)

	//  slices learning
	arr3 := [4]int{1, 2, 3, 4}
	slice1 := arr3[:]
	fmt.Println(slice1)
	arr3[1] = 34
	slice1[2] = 45
	fmt.Println(arr3, slice1)

	// slice continued without array
	slice2 := []int{7, 8, 9, 0}
	fmt.Println(slice2)

	slice2 = append(slice2, 10, 11, 12, 13)
	fmt.Println(slice2)

	// slice of slices
	s2 := slice2[1:]
	s3 := slice2[:4]
	s4 := slice2[1:3]
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	// maps
	fmt.Println("===== maps =====")
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 2
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	m1 := map[string]int{"foo": 42, "bar": 34}
	fmt.Println(m1)

	// struct datatypes
	fmt.Println("===== structs =====")
}