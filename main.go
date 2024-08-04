package main

import (
	"flag"
	"fmt"
)

func main() {
	firstName := flag.String("word", "default value", "description")
	age := flag.Int("age", 23, "age of user")
	flag.Parse()
	fmt.Println("name", *firstName, *age)
}
