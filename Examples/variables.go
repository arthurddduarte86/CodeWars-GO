package main

import "fmt"

var ab int
var integerNum = 2
var P, Q, R int = 2, 3, 4

//  var c := 3;  :=  Can only be used inside functions

func main() {
	var student1 string = "Arthur Duarte" //type is string
	var student2 = "Maria"                //type is inferred
	x := 2                                //type is inferred
	var a string                          //Variable Declaration Without Initial Value
	var b int                             //Variable Declaration Without Initial Value
	var c bool                            //Variable Declaration Without Initial Value

	fmt.Println("student1 ", student1)
	fmt.Println("student2 ", student2)
	fmt.Println(x)
	fmt.Println("a (string) ", a)
	fmt.Println("b ", b)
	fmt.Println("c ", c)

	fmt.Println(ab)
	fmt.Println(integerNum)
	fmt.Printf("P, Q, R são %v %v e %v\n", P, Q, R)

	//fmt.Println(c)

}
