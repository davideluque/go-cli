package main

import (
	"flag";
	"fmt"
)

func main() {
	dice := flag.String("d", "d6", "dice to roll. e.g. d6, d10, d20, etc. default is d6")

	flag.Parse()
	
	fmt.Println("Rolling", *dice)
}