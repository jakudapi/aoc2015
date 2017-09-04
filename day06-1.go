/*
--- Day 6: Probably a Fire Hazard ---

Because your neighbors keep defeating you in the holiday house decorating contest year after year, you've decided to deploy one million lights in a 1000x1000 grid.

Furthermore, because you've been especially nice this year, Santa has mailed you instructions on how to display the ideal lighting configuration.

Lights in your grid are numbered from 0 to 999 in each direction; the lights at each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions include whether to turn on, turn off, or toggle various inclusive ranges given as coordinate pairs. Each coordinate pair represents opposite corners of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore refers to 9 lights in a 3x3 square. The lights all start turned off.

To defeat your neighbors this year, all you have to do is set up your lights by doing the instructions Santa sent you in order.

For example:

    turn on 0,0 through 999,999 would turn on (or leave on) every light.
    toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off the ones that were on, and turning on the ones that were off.
    turn off 499,499 through 500,500 would turn off (or leave off) the middle four lights.

After following the instructions, how many lights are lit?

Your puzzle answer was 400410.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func print_grid(grid *[1000][1000]bool) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
}

type Instruction struct {
	Type  string
	X1    int
	Y1    int
	X2    int
	Y2    int
	value bool
}

func parse_instr(s string) Instruction {
	split := strings.Split(s, " ")
	if len(split) == 4 {
		coord1 := strings.Split(split[1], ",")
		coord2 := strings.Split(split[3], ",")
		x1, _ := strconv.Atoi(coord1[0])
		y1, _ := strconv.Atoi(coord1[1])
		x2, _ := strconv.Atoi(coord2[0])
		y2, _ := strconv.Atoi(coord2[1])
		return Instruction{"toggle", x1, y1, x2, y2, false}
	} else {
		coord1 := strings.Split(split[2], ",")
		coord2 := strings.Split(split[4], ",")
		x1, _ := strconv.Atoi(coord1[0])
		y1, _ := strconv.Atoi(coord1[1])
		x2, _ := strconv.Atoi(coord2[0])
		y2, _ := strconv.Atoi(coord2[1])
		onoff := false
		if split[1] == "on" {
			onoff = true
		}
		return Instruction{"turn", x1, y1, x2, y2, onoff}
	}
}

func exec(grid *[1000][1000]bool, instruc Instruction) {
	if instruc.Type == "toggle" {
		for x := instruc.X1; x <= instruc.X2; x++ {
			for y := instruc.Y1; y <= instruc.Y2; y++ {
				grid[y][x] = !grid[y][x]
			}

		}
	} else {
		for x := instruc.X1; x <= instruc.X2; x++ {
			for y := instruc.Y1; y <= instruc.Y2; y++ {
				grid[y][x] = instruc.value
			}
		}
	}
}

func count(grid *[1000][1000]bool) int {
	counter := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[y][x] == true {
				counter += 1
			}
		}
	}
	return counter
}

func main() {
	start := time.Now()

	file, err := os.Open("day6input.txt")
	if err != nil {
		panic("couldn't open or find file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [1000][1000]bool

	for scanner.Scan() {
		instruction := parse_instr(scanner.Text())
		//fmt.Println(instruction)
		exec(&grid, instruction)
	}

	fmt.Println(count(&grid))
	duration := time.Now().Sub(start)
	fmt.Println("This took ", duration)

}
