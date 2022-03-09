package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func bufioReader() {
	file, _ := os.Open("./example.txt")

	reader := bufio.NewReader(file)
	line, _, _ := reader.ReadLine()
	fmt.Println(string(line))

	_ = file.Close()
}

func readFile() {
	file, err := ioutil.ReadFile("./example.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(file))
}
