/*
--- Day 10: Elves Look, Elves Say ---
Neat, right? You might also enjoy hearing John Conway talking about this sequence (that's Conway of Conway's Game of Life fame).

Now, starting again with the digits in your puzzle input, apply this process 50 times. What is the length of the new result?

Your puzzle answer was 5103798.
*/

def lookSay(input: StringBuilder): StringBuilder = {
  var output = new StringBuilder
  var counter: Int = 0
  var current = 'x'

  for (character <- input){
    if (current == 'x'){
      current = character
      counter = 1
    }
    else if (character == current) counter += 1
    else {
      output += counter.toString()(0)
      output += current
      counter = 1
      current = character
    }
  }
  output += counter.toString()(0)
  output += current
  //this version takes 50-60 milliseconds for 40 cycles !
}

val startTime = System.nanoTime()
val CYCLES = 50
var puzzleInput: StringBuilder = new StringBuilder("1113122113")

for (cycle <- 1 to CYCLES) {
  puzzleInput = lookSay(puzzleInput)
  //println(puzzleInput)
}

println(s"Length is ${puzzleInput.length}")
val endTime = System.nanoTime()

println(s"took ${(endTime-startTime)/1000000} milliseconds")

