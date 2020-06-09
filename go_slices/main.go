package main

import "fmt"

func main() {
	// variable s is an array of strings length of 3
	// look up make function
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// set values to array
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	// check length
	fmt.Println("len:", len(s))

	// append (might add to array? will it break the array)?

	/*
	 append = returns a slice containing one or more new values

	 slice = operator with the syntax slice(low:high). Gets a slice of the elements 
	 s[2], s[3] and s[4]
	*/
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	// look up copy function
	copy(c, s)
	fmt.Println("apd:", s)

	l := s[2:5]
	fmt.Println("sl2:", l)

	l = s[:5]
	fmt.Println("sl3:", l)

	l = s[2:]
	fmt.Println("dcl:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
