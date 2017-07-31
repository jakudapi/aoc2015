/*
--- Day 3: Perfectly Spherical Houses in a Vacuum ---

--- Part Two ---

The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the same starting house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the previous year.

This year, how many houses receive at least one present?

For example:

    ^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
    ^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
    ^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.

*/


import scala.io.Source
import scala.collection.mutable

val houses = mutable.Map((0,0) -> 2)
val directions  = Source.fromFile("day03input.txt").mkString
var santaX: Int = 0
var santaY: Int = 0
var robotX: Int = 0
var robotY: Int = 0
var santasTurn: Boolean = true
var newLoc = (0,0)

for (direction <- directions){
	if (santasTurn){
		newLoc = direction match {
			case '^' => (santaX, santaY+1)
			case 'v' => (santaX, santaY-1)
			case '>' => (santaX+1, santaY)
			case '<' => (santaX-1, santaY)
		}
		santaX = newLoc._1 //update Santa's x position
		santaY = newLoc._2 //update Santa's y position
	}
	else {
		newLoc = direction match {
			case '^' => (robotX, robotY+1)
			case 'v' => (robotX, robotY-1)
			case '>' => (robotX+1, robotY)
			case '<' => (robotX-1, robotY)
		}
	robotX = newLoc._1 //update robot's x position
	robotY = newLoc._2 //update robot's y position
	}
	houses(newLoc) = 1 + houses.getOrElse(newLoc,0) // deliver gift
	santasTurn = !santasTurn

}

println(s"Houses which got presents: ${houses.size}")