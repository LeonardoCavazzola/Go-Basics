package main

import (
	"fmt"
	"net/http"
)

func methodGet() {
	googleUri := "https://google.com"
	response, _ := http.Get(googleUri)
	fmt.Println(response.Status)
}
