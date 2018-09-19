package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "set your name to be greeted")
	flag.Parse()
	if name == "" {
		fmt.Println("Hello world!")
	} else {
		fmt.Printf("Hello %s!\n", name)
	}

}
