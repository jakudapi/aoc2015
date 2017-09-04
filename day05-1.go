/*--- Day 5: Doesn't He Have Intern-Elves For This? ---

Santa needs help figuring out which strings in his text file are naughty or nice.

A nice string is one with all of the following properties:

    It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
    It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
    It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.

For example:

    ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
    aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
    jchzalrnumimnmhp is naughty because it has no double letter.
    haegwjzuvuyypxyu is naughty because it contains the string xy.
    dvszwmarrgswjxmb is naughty because it contains only one vowel.

How many strings are nice?

Your puzzle answer was 258.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

var DOUBLE_LTR, _ = regexp.Compile(`aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz`)

//golang doesn't backsearch!
var BAD_STRING, _ = regexp.Compile(`ab|cd|pq|xy`)
var VOWELS, _ = regexp.Compile(`a|e|i|o|u`)

func double_letter(s string) bool {
	if len(DOUBLE_LTR.FindString(s)) == 0 {
		return false
	} else {
		return true
	}
}

func bad_string(s string) bool {
	if len(BAD_STRING.FindString(s)) == 0 {
		return false
	} else {
		return true
	}
}

func three_vowels(s string) bool {
	vowel_array := VOWELS.FindAllString(s, -1)
	if vowel_array == nil {
		return false
	} else if len(vowel_array) >= 3 {
		return true
	} else {
		return false
	}
}

func main() {
	start := time.Now()

	file, err := os.Open("day05input.txt")
	if err != nil {
		panic("couldn't open or find file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0

	for scanner.Scan() {
		input_str := scanner.Text()
		if !bad_string(input_str) && three_vowels(input_str) && double_letter(input_str) {
			counter += 1
		}
	}

	fmt.Println(counter)
	duration := time.Now().Sub(start)
	fmt.Print("This took ", duration)

}
