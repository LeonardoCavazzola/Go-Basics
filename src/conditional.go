package main

import "fmt"

func ifElse() {
	if true {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func switchCase() {
	a := 1
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}
}
