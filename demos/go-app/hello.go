package main

import (
	"example/hello/greetings"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, world")
	fmt.Println(greetings.Hello("Jim"))
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d bool = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	const n = 500000000

    const g = 3e20 / n
    fmt.Println(g)

    fmt.Println(int64(g))

    fmt.Println(math.Sin(n))
}
