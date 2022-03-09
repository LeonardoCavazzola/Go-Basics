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

func writeFile() {
	file, _ := os.OpenFile("./wrote.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	_, _ = file.WriteString("abc\n")
	_ = file.Close()
}
