package main

import(
	"fmt"
)

var x = 45
var y = "Hello World!"

func main(){
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	fmt.Println(y)
	fmt.Printf("y is of type %T\n", y)
}
