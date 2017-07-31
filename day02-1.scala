/*--- Day 2: I Was Told There Would Be No Math ---

The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

For example:

    A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
    A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.

All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
*/

import scala.io.Source
import scala.collection.mutable.ListBuffer
import scala.math.min

val startTime = System.nanoTime()

val gifts = ListBuffer[String]()
var counter: Int = 0

for (item <- Source.fromFile("day02input.txt").getLines){
  gifts.append(item)
}

val dimensions = gifts.map(_.split('x'))
for (gift <-dimensions){
	val l = gift(0).toInt
	val w = gift(1).toInt
	val h = gift(2).toInt
	counter += 2*(l*w + l*h + h*w) + min(min(l*w, l*h), h*w)
}

val endTime = System.nanoTime()

println(counter)
println(s"This took ${(endTime-startTime)/1000} microseconds")

