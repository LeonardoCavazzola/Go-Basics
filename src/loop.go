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

func forI() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func forRange() {
	numbers := []string{"zero", "one", "two", "three", "four"}
	for i, item := range numbers {
		fmt.Println("index =", i, "| item =", item)
	}
}
