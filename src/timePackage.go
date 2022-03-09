package main

import (
	"fmt"
	"time"
)

func sleeping() {
	fmt.Println("Sleeping")
	time.Sleep(5 * time.Second)
	fmt.Println("Returning")
}

func now() {
	now1 := time.Now()
	fmt.Println(now1)
	formatted := now1.Format("02/01/2006 15:04:05")
	fmt.Println(formatted)
}
