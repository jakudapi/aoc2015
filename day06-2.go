/*
--- Day 6: Probably a Fire Hazard ---

--- Part Two ---

You just finish implementing your winning light pattern when you realize you mistranslated Santa's message from Ancient Nordic Elvish.

The light grid you bought actually has individual brightness controls; each light can have a brightness of zero or more. The lights all start at zero.

The phrase turn on actually means that you should increase the brightness of those lights by 1.

The phrase turn off actually means that you should decrease the brightness of those lights by 1, to a minimum of zero.

The phrase toggle actually means that you should increase the brightness of those lights by 2.

What is the total brightness of all lights combined after following Santa's instructions?

For example:

    turn on 0,0 through 0,0 would increase the total brightness by 1.
    toggle 0,0 through 999,999 would increase the total brightness by 2000000.

Your puzzle answer was 15343601.
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

func print_grid(grid *[1000][1000]int) {
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
	value int
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
		return Instruction{"toggle", x1, y1, x2, y2, 2}
	} else {
		coord1 := strings.Split(split[2], ",")
		coord2 := strings.Split(split[4], ",")
		x1, _ := strconv.Atoi(coord1[0])
		y1, _ := strconv.Atoi(coord1[1])
		x2, _ := strconv.Atoi(coord2[0])
		y2, _ := strconv.Atoi(coord2[1])
		onoff := -1
		if split[1] == "on" {
			onoff = 1
		}
		return Instruction{"turn", x1, y1, x2, y2, onoff}
	}
}

func exec(grid *[1000][1000]int, instruc Instruction) {
	if instruc.Type == "toggle" {
		for x := instruc.X1; x <= instruc.X2; x++ {
			for y := instruc.Y1; y <= instruc.Y2; y++ {
				grid[y][x] += 2
			}

		}
	} else {
		for x := instruc.X1; x <= instruc.X2; x++ {
			for y := instruc.Y1; y <= instruc.Y2; y++ {
				grid[y][x] += instruc.value
				if grid[y][x] <= 0 {
					grid[y][x] = 0
				}
			}
		}
	}
}

func count(grid *[1000][1000]int) int {
	counter := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			counter += grid[y][x]
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
	var grid [1000][1000]int

	for scanner.Scan() {
		instruction := parse_instr(scanner.Text())
		//fmt.Println(instruction)
		exec(&grid, instruction)
	}

	fmt.Println(count(&grid))
	duration := time.Now().Sub(start)
	fmt.Println("This took ", duration)

}
