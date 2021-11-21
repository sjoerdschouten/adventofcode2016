package solutions

import (
	"fmt"

	"github.com/sjoerdschouten/adventofcode2016/utils"
)

func Solve1() {
	currentNumber := 5
	var bathroomCode []int
	input, _ := utils.ReadLines("./input/input")

	for _, instructions := range input {
		//fmt.Println(instructions)
		for _, instruction := range instructions {

			//fmt.Printf("InstructionChar: %s\n", string(instruction))
			switch instruction {
			case 'U':
				// handle up
				if currentNumber == 1 || currentNumber == 2 || currentNumber == 3 {
					continue
				} else {
					currentNumber -= 3
				}
			case 'D':
				// handle down
				if currentNumber == 7 || currentNumber == 8 || currentNumber == 9 {
					continue
				} else {
					currentNumber += 3
				}
			case 'R':
				// handle right
				if currentNumber == 3 || currentNumber == 6 || currentNumber == 9 {
					continue
				} else {
					currentNumber += 1
				}
			case 'L':
				// handle left
				if currentNumber == 1 || currentNumber == 4 || currentNumber == 7 {
					continue
				} else {
					currentNumber -= 1
				}
			}
		}
		bathroomCode = append(bathroomCode, currentNumber)
	}
	fmt.Printf("The code for part 1 is: %v\n", bathroomCode)
}

func Solve2() {
	currentChar := 'B'
	var bathroomCode []rune
	input, _ := utils.ReadLines("./input/input")

	for _, instructions := range input {
		//fmt.Println(instructions)
		for _, instruction := range instructions {

			//fmt.Printf("current position: %q\n", currentChar)
			//fmt.Printf("InstructionChar: %s\n", string(instruction))
			switch instruction {
			case 'U':
				// handle up
				if currentChar == '5' || currentChar == '2' || currentChar == '1' || currentChar == '4' || currentChar == '9' {
					continue
				} else if currentChar == '3' {
					currentChar = '1'
				} else if currentChar == 'A' {
					currentChar = '6'
				} else if currentChar == 'B' {
					currentChar = '7'
				} else if currentChar == 'C' {
					currentChar = '8'
				} else if currentChar == 'D' {
					currentChar = 'B'
				} else {
					currentChar -= 4
				}
			case 'D':
				// handle down
				if currentChar == '5' || currentChar == '9' || currentChar == 'A' || currentChar == 'C' || currentChar == 'D' {
				} else if currentChar == '1' {
					currentChar = '3'
				} else if currentChar == '6' {
					currentChar = 'A'
				} else if currentChar == '7' {
					currentChar = 'B'
				} else if currentChar == '8' {
					currentChar = 'C'
				} else if currentChar == 'B' {
					currentChar = 'D'
				} else {
					currentChar += 4
				}
			case 'R':
				// handle right
				if currentChar == '1' || currentChar == '4' || currentChar == '9' || currentChar == 'C' || currentChar == 'D' {
				} else if currentChar == 'A' {
					currentChar = 'B'
				} else if currentChar == 'B' {
					currentChar = 'C'
				} else {
					currentChar += 1
				}
			case 'L':
				// handle left
				if currentChar == '1' || currentChar == '2' || currentChar == '5' || currentChar == 'A' || currentChar == 'D' {

				} else if currentChar == 'B' {
					currentChar = 'A'

				} else if currentChar == 'C' {
					currentChar = 'B'
				} else {
					currentChar -= 1
				}
			}
		}
		//fmt.Printf("Line done, adding %q to code \n", currentChar)
		bathroomCode = append(bathroomCode, currentChar)
	}
	fmt.Printf("The code for part 2 is: %q", bathroomCode)
}
