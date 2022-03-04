package main

import "fmt"

func array() {
	const size = 10
	var array [size]string
	array[0] = "0-pos"
	array[1] = "1-pos"
	array[2] = "2-pos"

	fmt.Println("Array =", array)
	fmt.Println("Not void position = '", array[2], "'")
	fmt.Println("Void position is = '", array[3], "'")
}

func slice() {
	var slice = []string{"0-pos", "1-pos", "2-pos"}
	fmt.Println("Slice =", slice)
	fmt.Println("Length =", len(slice))
	slice[2] = "Replaced"

	slice = append(slice, "Added")
	fmt.Println("Slice =", slice)
	fmt.Println("Length =", len(slice))

}
