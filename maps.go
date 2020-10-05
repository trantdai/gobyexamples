package main

import "fmt"

func maps() {
	//To create an empty map, use the builtin make: make(map[key-type]val-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	//The builtin delete removes map element whose key is "k2" from a map.
	delete(m, "k2")
	//map[k1:7]
	fmt.Println("map:", m)

	//The optional second return value ( prs in this case) when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _. prs is false because the key "k2" is not present in the map.
	_, prs := m["k2"]
	//prs: false
	fmt.Println("prs:", prs)

	//Declare and initialize a new map in the same line with this syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
