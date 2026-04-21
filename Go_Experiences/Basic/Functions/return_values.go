// return_values.go

package main

import "fmt"

func response(v1, v2, v3 int) (bool, string) {
	status := v1 >= 0 && v2 >= 0 && v3 >= 0
	if status {
		return status, "Ok"
	}
	return status, "One or more values ​​are invalid"
}

func calc(nums ...float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(nums, " : ", total)
}

func intSeq() func() int {

	i := 0 // closure

	// return anonymous function
	return func() int {
		i++
		return i
	}
}

func main() {
	// Multiple return values
	fmt.Println(response(0, 0, 0))

	_, res := response(-1, 0, 1)
	fmt.Println(res)

	// Variadic functions
	fmt.Println(
		"Sum: ",
		calc(1.1, 2.2, 3.3, 4.4, 5.5))

	sum([]int{1, 3, 7, 9}...)

	// Anonymous functions
	func() {
		fmt.Println("Anonymous Self-Execution Function")
	}()

	mult := func(n1, n2 int) { fmt.Println(n1, " * ", n2, " = ", n1*n2) }

	mult(10, -10)
	mult(25, 25)

	// With Closure
	nextInt := intSeq()
	fmt.Printf("%v %v %v\n", nextInt(), nextInt(), nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
