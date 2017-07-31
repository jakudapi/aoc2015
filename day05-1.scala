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
*/

import scala.io.Source
//import scala.util.matching.Regex

var counter: Int = 0
val pattern1 = raw"ab|cd|pq|xy".r
val pattern2 = raw"([a-z])\1{1,}".r // duplicate letters anywhere
val isVowels = Set('a', 'e', 'i', 'o', 'u')

val input = Source.fromFile("day05input.txt").getLines.toArray

for (word <- input){
  val vowels = word.count(isVowels(_)) //count vowels in word

  pattern1.findFirstIn(word) match {
    case Some(x) => println(s"Found ${x} in ${word}. NAUGHTY!")
    case None => pattern2.findFirstIn(word) match {
              case Some(x) if vowels >= 3 => println(s"${word} is a NICE word!") ; counter += 1
              case _ => println(s"${word} is NAUGHTY!")
    }
    
  }
}
println(counter)