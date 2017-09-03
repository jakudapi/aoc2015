/*--- Day 2: I Was Told There Would Be No Math ---

The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

For example:

    A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
    A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.

All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
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
  side1 := dim1 * dim2
  side2 := dim1 * dim3
  side3 := dim2 * dim3
	return 2 * (side1 + side2 + side3) + get_min(side1, get_min(side2, side3))
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
