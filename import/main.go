package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Syntax: import <module>")
		os.Exit(1)
		return
	}

	p := os.Args[1]

	r, err := getModulePage(p)
	if err != nil {
		panic(err)
	}

	m, err := parseModulePage(r)
	if err != nil {
		panic(err)
	}

	text, err := Generate(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(text)
}
