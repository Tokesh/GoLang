package main

import "fmt"

func fibbo(a int) int {
	if a == 1 || a == 0 {
		return 1
	}
	return (fibbo(a-1) + fibbo(a-2))
}

func fillArray(arr []int) {
	for i := 0; i < 10; i++ {
		arr[i] = 5
	}
}

func updateVal(emp *int) {
	PI := 24
	*emp = PI
	//fmt.Println(*emp, PI)
}
func updateVal_v2(emp *int) int {
	PI := 32
	emp = &PI
	fmt.Println(*emp)
	return *emp
}

func main() {
	//fmt.Println(fibbo(3))

	// array := [50]int{}
	// array[5] = 2
	// fmt.Println(array)

	// var array [15]int
	// fillArray(array[:])
	// fmt.Println(array)

	emp := 10
	updateVal(&emp)
	//emp = updateVal_v2(&emp)
	fmt.Println(emp)
	fmt.Println("Hello World")
}
