package main

import (
	"fmt"
)

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)

	i3, ok := i.(string)
	if !ok {
		fmt.Println("i is not a string",i)
	}
	fmt.Println("i3 is : \"", i3, "\"")
}
