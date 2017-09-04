/*
--- Day 3: Perfectly Spherical Houses in a Vacuum ---

Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

For example:

    > delivers presents to 2 houses: one at the starting location, and one to the east.
    ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
    ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

*/

package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

/* This could be done more simply but I wanted to play around with golang structs and methods */
type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) Left() {
	c.X--
}

func (c *Coordinate) Up() {
	c.Y++
}

func (c *Coordinate) Down() {
	c.Y--
}

func (c *Coordinate) Right() {
	c.X++
}

func main() {
	start := time.Now()

	file, err := ioutil.ReadFile("day03input.txt")
	if err != nil {
		panic("couldn't open or find file")
	}

	data := string(file)

	pos := Coordinate{0, 0}
	tracker := make(map[Coordinate]int)
	tracker[Coordinate{0, 0}] = 1

	for _, r := range data {
		c := rune(r)
		switch c {
		case 'v':
			pos.Down()
		case '>':
			pos.Right()
		case '^':
			pos.Up()
		case '<':
			pos.Left()
		}
		tracker[pos] += 1
	}

	fmt.Println("Houses with gift:", len(tracker))
	duration := time.Now().Sub(start)
	fmt.Print("This took ", duration)

}
