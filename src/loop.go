package main

import "fmt"

func whileTrueAndBreak() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
}
