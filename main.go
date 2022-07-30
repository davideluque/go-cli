package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	dice := flag.String("d", "d6", "dice to roll. e.g. d6, d10, d20, etc. default is d6")

	flag.Parse()

	matched, _ := regexp.MatchString("d[0-9]+", *dice)

	if matched {
		diceSides := regexp.MustCompile("[0-9]+").FindString(*dice)

		// Cast diceSides to int and roll the dice
		sides, _ := strconv.Atoi(diceSides)

		roll := rand.Intn(sides) + 1

		fmt.Println("You rolled a", roll)
	} else {
		fmt.Println("Invalid dice")
		fmt.Println("Usage: dice -d d6")
	}
}
