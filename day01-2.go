/*
--- Day 1: Not Quite Lisp ---

--- Part Two ---

Now, given the same instructions, find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.

For example:

    ) causes him to enter the basement at character position 1.
    ()()) causes him to enter the basement at character position 5.

What is the position of the character that causes Santa to first enter the basement?


*/

package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	start := time.Now()
	data, err := ioutil.ReadFile("day01input.txt")
	if err != nil {
		panic("can't find input file!")
	}
	var counter int8 = 0
	for i, r := range data {
		c := rune(r)
		switch c {
		case '(':
			counter++
		case ')':
			counter--
		}
		if counter == -1 {
			fmt.Println(i + 1)
			break}
	}

	duration := time.Now().Sub(start)
	fmt.Println("This took ", duration)
}
