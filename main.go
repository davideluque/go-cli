package main

import (
	"flag"
	"fmt"
	"regexp"
)

func main() {
	dice := flag.String("d", "d6", "dice to roll. e.g. d6, d10, d20, etc. default is d6")

	flag.Parse()

	matched, _ := regexp.MatchString("d[0-9]+", *dice)

	if matched {
		diceSides := regexp.MustCompile("[0-9]+").FindString(*dice)
		fmt.Println(diceSides)
	} else {
		fmt.Println("Invalid dice")
	}
}
