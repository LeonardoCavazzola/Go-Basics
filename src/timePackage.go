package main

import (
	"fmt"
	"time"
)

func tryTime() {
	fmt.Println("Sleeping")
	time.Sleep(5 * time.Second)
	fmt.Println("Returning")
}
