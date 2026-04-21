/*
 * dataTypes.go
 */

package main

import (
	"fmt"
)

func main() {
	// Basic data types
	var v1 bool
	v1 = true

	var v2 int = 5
	var v3 float32 = 3.14
	var v4 string = "Hi!"

	v5 := 15.5
	v6 := 10e-3

	fmt.Printf("%T\n", v1)
	fmt.Printf("%T\n", v2)
	fmt.Printf("%T\n", v3)
	fmt.Printf("%T\n", v4)
	fmt.Printf("%T\n", v5)
	fmt.Printf("%T\n", v6)

	fmt.Printf("%v\n", v6)
	fmt.Printf("%#v\n", v6)
	fmt.Printf("%v%%\n", v6)

	// Arrays
	var arr1 = []int{10, 20, 30, 35, 45, 65}
	arr2 := [...]float64{10.5, 25.5, 30.1}

	fmt.Printf("%T\n%T\n", arr1, arr2)

	arr3 := [4]string{}
	arr3[0] = "One"
	arr3[2] = "Two"

	fmt.Printf("%v\n", arr3)
	fmt.Printf("%#v\n", arr3)
	fmt.Printf("%v\n", arr3[0])

	fmt.Printf("%T\n", arr3)

	fmt.Printf("length  : %v\n", len(arr3))
	fmt.Printf("Capacity: %v\n", cap(arr3))

	slice1 := arr1[2:4]
	fmt.Printf("%#v\n", slice1)

	// slice_name := make([]type, length, capacity)
	slice2 := make([]bool, 5, 10)
	fmt.Printf("-> length  : %v\nCapacity: %v\n", len(slice2), cap(slice2))
	slice2 = append(slice2, true, true, true, true, true, true)
	fmt.Printf("-> length  : %v\nCapacity: %v\n", len(slice2), cap(slice2))
	fmt.Printf("%#v\n", slice2)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	neededNumbers := numbers[:len(numbers)-5]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", numbersCopy)
}
