package main

import (
	"fmt"
	"os"

	basm "github.com/BoburF/bhsml/src/bhsml-asm"
	"github.com/BoburF/browser/renderer"
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
	if err != nil {
		fmt.Println("Error assembling bhsml:", err)
		return
	}

	fmt.Println("JSON data written to ast.json successfully")

	renderer.Render()
}

