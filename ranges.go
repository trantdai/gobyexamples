package main

import "fmt"

func ranges() {
	//range iterates over elements in a variety of data structures
	nums := []int{2, 3, 4}
	sum := 0
	//Here we use range to sum the numbers in a slice
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			if num == 3 {
				fmt.Println("index:", i)
			}
		}
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	//range on map iterates over key/value pairs
	for k, v := range kvs {
		fmt.Println("%s -> %s\n", k, v)
	}
	//range can also iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}
	//range on strings iterates over Unicode code points. The first value is the starting byte index of the range and the second the range itself.
	for i, c := range "AB" {
		fmt.Println(i, c)
	}
}
