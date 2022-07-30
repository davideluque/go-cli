package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	dice := flag.String("d", "d6", "dice to roll. e.g. d6, d10, d20, etc. default is d6")
	numRoll := flag.Int("n", 1, "number of rolls. default is 1")
	sum := flag.Bool("s", false, "sum the rolls. default is false")
	advantage := flag.Bool("adv", false, "roll with advantage. default is false")
	disadvantage := flag.Bool("dis", false, "roll with disadvantage. default is false")

	flag.Parse()

	matched, _ := regexp.MatchString("d[0-9]+", *dice)

	if matched {
		rolls := rollDice(dice, numRoll)

		// to avoid adding functions to our call stack, we will collect the rolls
		// and print them all at once
		printRolls(rolls)

		if *sum {
			fmt.Println("Sum: ", sumDice(rolls))
		}

		if *advantage {
			fmt.Println("Advantage: ", rollWithAdvantage(rolls))
		}

		if *disadvantage {
			fmt.Println("Disadvantage: ", rollWithDisadvantage(rolls))
		}
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

func sumDice(rolls []int) int {
	sum := 0

	for _, roll := range rolls {
		sum += roll
	}

	return sum
}

func rollWithAdvantage(rolls []int) int {
	// if we have more than one roll, we need to get the highest
	if len(rolls) > 1 {
		// sort the rolls
		sort.Ints(rolls)

		// return the last roll
		return rolls[len(rolls)-1]
	}

	// otherwise, we just return the only roll
	return rolls[0]
}

func rollWithDisadvantage(rolls []int) int {
	// if we have more than one roll, we need to get the lowest
	if len(rolls) > 1 {
		// sort the rolls
		sort.Ints(rolls)

		// return the first roll
		return rolls[0]
	}

	// otherwise, we just return the only roll
	return rolls[0]
}
