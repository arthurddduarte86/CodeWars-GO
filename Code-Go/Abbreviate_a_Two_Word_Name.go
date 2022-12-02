/*
Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot separating them.

It should look like this:

Sam Harris => S.H

patrick feeney => P.F
*/

package kata

import "strings"

func AbbrevName(name string) string{
  nameSliced := strings.Split(name, " ")
  return strings.ToUpper(string(nameSliced[0][0])) + "." + strings.ToUpper(string(nameSliced[1][0]))
}




package kata
import "strings"

func AbbrevName(name string) string{
    var parts []string
    for _, part := range strings.Split(name, " ") {
        parts = append(parts, strings.ToUpper(part[:1]))
    }
    return strings.Join(parts, ".")
}