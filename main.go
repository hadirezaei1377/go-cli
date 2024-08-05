package main

import (
	"flag"
	"fmt"
)

func main() {
	// flags definition
	name := flag.String("name", "World", "a name to say hello to")
	age := flag.Int("age", 0, "your age")

	// parse flags
	flag.Parse()

	// use flags
	fmt.Printf("Hello, %s!\n", *name)
	fmt.Printf("You are %d years old.\n", *age)
}
