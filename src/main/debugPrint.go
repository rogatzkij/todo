package main

import "fmt"

const (
	DEBUG = true
)

func debug(msg string) {

	if DEBUG {
		fmt.Println(msg)
	}
}
