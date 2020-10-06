package main

import "fmt"

func slices() {
	//An array type definition specifies a length and an element type. An array's size is fixed; its length is part of its type ([4]int and [5]int are distinct, incompatible types.
	//The type specification for a slice is []T, where T is the type of the elements of the slice. Unlike an array type, a slice type has no specified length.
	//Create an empty string of 3 elements
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "aaaa"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	//Add "d" as a new element [aaaa b c d]
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Appended:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	//Get a slice of elements s[2], s[3], and s[4] excluding s[5]
	//s = [a b c d e f] => l = [c d e]
	l := s[2:5]
	fmt.Println("sl1:", l)

	//This slices up to but exclding s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	//This slices up from and including s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	//Delare and initialize a variable for slice in a single line as well
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
