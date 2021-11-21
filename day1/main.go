package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/sjoerdschouten/adventofcode2016/utils"
)

type Point struct {
	X int
	Y int
}

type Position struct {
	point Point
	// 0 = north, 1 = east, 2 = south, 3 = west
	direction int
	previouslyVisited map[Point]bool
	foundHQ bool
	totalDistance int
	distanceToHQ int
}

func (position *Position) turnRight() {
	if position.direction == 3 {
		position.direction = 0
	} else {
		position.direction++
	}
}

func (position *Position) turnLeft() {
	if position.direction == 0 {
		position.direction = 3
	} else {
		position.direction--
	}
}

func (position *Position) move(blocks int) {
	if position.direction == 0 {
		position.point.Y += blocks
	} else if position.direction == 1 {
		position.point.X += blocks
	} else if position.direction == 2 {
		position.point.Y -= blocks
	} else if position.direction == 3 {
		position.point.X -= blocks
	}
	position.addCurrentLocationToVisited(position.point)
}

func (position *Position) addCurrentLocationToVisited(currentLocation Point) {
	if position.previouslyVisited[currentLocation] {
		println(int(math.Abs(float64(currentLocation.X)) + math.Abs(float64(currentLocation.Y))))
		position.determineHeadquarters()
	}
	position.previouslyVisited[currentLocation] = true
}

func (position *Position) determineHeadquarters() {
	position.calculateDistance()
	position.foundHQ = true
	position.distanceToHQ = position.totalDistance

}

func (position *Position) calculateDistance() {
	position.totalDistance = int(math.Abs(float64(position.point.X)) + math.Abs(float64(position.point.Y)))
}

func (position *Position) calcDestination(instructions []string) *Position {
	for _, instruction := range instructions {
		if instruction[0] == 'R' {
			position.turnRight()
		} else {
			position.turnLeft()
		}
		blocks, _ := strconv.Atoi(instruction[1:])
		position.move(blocks)
		if position.foundHQ {
			continue
		}
	}
	return position
}

func splitInput(input string) (result []string) {
	return strings.Split(input, ",")
}

func main() {
	input := splitInput(strings.ReplaceAll(utils.ReadEntireFile("./input.txt"), " ", ""))
	position := &Position {
		point: Point{
			X: 0,
			Y: 0,
		},
		direction: 0,
		previouslyVisited: make(map[Point]bool),
		foundHQ: false,
	}
	position.calcDestination(input)
	position.calculateDistance()
	fmt.Printf("Distance to HQ %d\n", position.distanceToHQ)
	fmt.Printf("totalDistance %d\n", position.totalDistance)
	// fmt.Printf("%+v", position)
}
