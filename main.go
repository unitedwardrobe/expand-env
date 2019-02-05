package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func expandContent(content []byte) {
	fmt.Printf("%s", os.ExpandEnv(string(content)))
}

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() > 0 {
		data, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			panic(err)
		}

		expandContent(data)
		return
	}

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s [filename]\n", os.Args[0])
		os.Exit(128)
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	expandContent(content)
}
