package main

import (
	"fmt"
	"log"
)

func main() {
	// Your CLI tool's entry point
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	// Your tool's main logic goes here
	fmt.Println("My CLI tool is running")
	return nil
}
