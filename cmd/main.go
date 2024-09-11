package main

import (
	"fmt"
	"os"

	"github.com/BoburF/bhsml/src/bhsml-asm"
)

func main() {
	file, err := os.Open("index.bhsml")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fo, err := os.Create("ast.json")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	err = basm.Assemble(file, fo)
	fmt.Println("JSON data written to ast.json successfully")
}
