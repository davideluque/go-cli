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
	numRoll := flag.Int("n", 1, "number of rolls. default is 1")

	flag.Parse()

	matched, _ := regexp.MatchString("d[0-9]+", *dice)

	if matched {
		rolls := rollDice(dice, numRoll)

		// to avoid adding functions to our call stack, we will collect the rolls
		// and print them all at once
		printRolls(rolls)
	} else {
		fmt.Println("Invalid dice")
		fmt.Println("Usage: dice -d d6")
	}
}

func rollDice(dice *string, numRolls *int) []int {
	rolls := make([]int, *numRolls)

	// get the number of sides
	diceSides := regexp.MustCompile("[0-9]+").FindString(*dice)

	sides, _ := strconv.Atoi(diceSides)

	for i := 0; i < *numRolls; i++ {
		roll := rand.Intn(sides) + 1
		rolls[i] = roll
	}

	return rolls
}

func printRolls(rolls []int) {
	for i, roll := range rolls {
		fmt.Println("Roll #", i+1, ":", roll)
	}
}
