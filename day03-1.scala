/*
--- Day 3: Perfectly Spherical Houses in a Vacuum ---

Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

For example:

    > delivers presents to 2 houses: one at the starting location, and one to the east.
    ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
    ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

*/


import scala.io.Source
import scala.collection.mutable

val startTime = System.nanoTime()
val houses = mutable.Map((0,0) -> 1)
val directions  = Source.fromFile("day03input.txt").mkString
var locX: Int = 0
var locY: Int = 0

for (direction <- directions){ 
	val newLoc = direction match {
	case '^' => (locX, locY+1)
	case 'v' => (locX, locY-1)
	case '>' => (locX+1, locY)
	case '<' => (locX-1, locY)
	}
	houses(newLoc) = 1 + houses.getOrElse(newLoc,0) // deliver gift
	locX = newLoc._1 //update Santa's x position
	locY = newLoc._2 //update Santa's y position
}
val endTime = System.nanoTime()
println(s"Houses which got presents: ${houses.size}")
println(s"This took ${(endTime-startTime)/1000000} milliseconds")