package main

import (
	"fmt"
	"os"
)

func main() {
	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)
	const s int = 10
	var x speed = speed(s)
	fmt.Println(x)
	fmt.Println(os.Args)

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	people := [5]string{"one", "two"}
	for index, person := range people {
		fmt.Println(index, person)
	}
}
