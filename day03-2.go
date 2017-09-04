/*
--- Day 3: Perfectly Spherical Houses in a Vacuum ---
--- Part Two ---

The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the same starting house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the previous year.

This year, how many houses receive at least one present?

For example:

    ^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
    ^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
    ^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.

Your puzzle answer was 2631.

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

	pos_santa := Coordinate{0, 0}
  pos_robot := Coordinate{0, 0}
	tracker := make(map[Coordinate]int)
	tracker[Coordinate{0, 0}] = 1  //starting house gets a present

  turn_tracker := true

	for _, r := range data {
		c := rune(r)
		switch c {
		case 'v':
      if turn_tracker{
        pos_santa.Down()
      } else {
        pos_robot.Down()
      }
		case '^':
      if turn_tracker{
        pos_santa.Up()
      } else {
        pos_robot.Up()
      }
    case '<':
      if turn_tracker{
        pos_santa.Left()
      } else {
        pos_robot.Left()
      }
    case '>':
      if turn_tracker{
        pos_santa.Right()
      } else {
        pos_robot.Right()
      }
		
		}
		if turn_tracker{
      tracker[pos_santa] += 1
    } else {
      tracker[pos_robot] += 1
    }
    turn_tracker = !turn_tracker  // switch between robot and Santa
	}

	fmt.Println("Houses with gift:", len(tracker))
	duration := time.Now().Sub(start)
	fmt.Print("This took ", duration)

}
