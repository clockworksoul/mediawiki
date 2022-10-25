package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	r, err := getModulePage("clientlogin")
	if err != nil {
		panic(err)
	}

	if m, err := parseModulePage(r); err != nil {
		panic(err)
	} else {
		b, _ := json.MarshalIndent(m, "", "  ")
		fmt.Println(string(b))

		fmt.Println("------------")
		fmt.Println(Generate(m))
	}
}
