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

	if m, err := parseModulePage(r); err != nil {
		panic(err)
	} else {
		if text, err := Generate(m); err != nil {
			panic(err)
		} else {
			fmt.Println(text)
		}
	}
}
