package main

import "fmt"

func main() {
	var a[5]int
	// creates array of 5 int
	fmt.Println("emp:", a)

	a[4] = 100
	// assign 100 to array {a} index{4}
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))
	// length is a function

	b := [5]int{1, 2, 3, 4, 5}
	// create array of 5 and plug in the following values
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	//Two arrays one of 2 and one of 3

	for i := 0; i < 2; i++ {
		// parent array loop has 2 length
		for j :=0; j < 3; j++ {
			// we now assign the value to the index of the array and assign i + j value
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}