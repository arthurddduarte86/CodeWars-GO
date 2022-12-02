/*
Write a function that accepts an integer n and a string s as parameters, and returns a string of s repeated exactly n times.

Examples (input -> output)
6, "I"     -> "IIIIII"
5, "Hello" -> "HelloHelloHelloHelloHello"
*/


package kata

func RepeatStr(repetitions int, value string) string {
  var message string = ""
  for counter:=0; counter<repetitions; counter++ {message+=value}
  return message
}

//
//

package kata
import "strings"

func RepeatStr(repetitions int, value string) string {
  return strings.Repeat(value, repetitions)
}