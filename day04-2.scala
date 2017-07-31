/*
--- Day 4: The Ideal Stocking Stuffer ---
Now find one that starts with six zeroes.
*/

import java.security.MessageDigest
import javax.xml.bind.DatatypeConverter

val startTime = System.nanoTime()

val input: String = "bgvyzdsv"
var test: Int = 0
var byteResult = MessageDigest.getInstance("MD5").digest("abcdef000".getBytes)

var hexResult: String = DatatypeConverter.printHexBinary(byteResult)

while (hexResult.slice(0,6) != "000000"){
	test += 1
	//println(test)
	byteResult = MessageDigest.getInstance("MD5").digest((input+test.toString).getBytes)
	hexResult = DatatypeConverter.printHexBinary(byteResult)
	//println(hexResult, hexResult.slice(0,6))

}
val endTime = System.nanoTime()
println(s"The first number that works is: ${test}, for a result of ${hexResult}")
println(s"This took ${(endTime-startTime)/1000000} milliseconds")
