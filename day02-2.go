/*--- Day 2: I Was Told There Would Be No Math ---
--- Part Two ---

The elves are also running low on ribbon. Ribbon is all the same width, so they only have to worry about the length they need to order, which they would again like to be exact.

The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face. Each present also requires a bow made out of ribbon as well; the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present. Don't ask how they tie the bow, though; they'll never tell.

For example:

    A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
    A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.

How many total feet of ribbon should they order?

Your puzzle answer was 3783758.
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

func get_min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func get_amount(dims []string) int {
	dim1, _ := strconv.Atoi(dims[0])
	dim2, _ := strconv.Atoi(dims[1])
	dim3, _ := strconv.Atoi(dims[2])
  side1 := 2 * (dim1 + dim2)
  side2 := 2 * (dim1 + dim3)
  side3 := 2 * (dim3 + dim2)
	return dim1 * dim2 * dim3 + get_min(side1, get_min(side2, side3))
}

func main() {
	start := time.Now()

	file, err := os.Open("day02input.txt")
	if err != nil {
		panic("couldn't open or find file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0

	for scanner.Scan() {
		dims := strings.Split(scanner.Text(), "x")
		counter += get_amount(dims)
	}

  fmt.Println(counter)
	duration := time.Now().Sub(start)
	fmt.Print("This took ", duration)

}
