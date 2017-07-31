/*
--- Day 1: Not Quite Lisp ---

--- Part Two ---

Now, given the same instructions, find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.

For example:

    ) causes him to enter the basement at character position 1.
    ()()) causes him to enter the basement at character position 5.

What is the position of the character that causes Santa to first enter the basement?


*/
import scala.io.Source


val inputFile = "day01input.txt"
val instructions: String = Source.fromFile(inputFile).getLines.next
var position: Int = 0
var counter: Int = 0
var flag: Boolean = false


def get_position(instructions: String): Int = {
  for (direction <- instructions){ direction match {
  case '(' => position += 1
  case ')' => position -= 1
  }
  if (flag == false) counter += 1
  if (position < 0 && flag == false){
    flag = true
    } 
  }
  counter
}

println(get_position(instructions))
