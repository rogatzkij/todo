package main

import "fmt"

const (
	DEBUG = true
)

func Debug(msg string) {

	if DEBUG {
		fmt.Println(msg)
	}
}
