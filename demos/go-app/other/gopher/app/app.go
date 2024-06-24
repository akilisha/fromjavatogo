package main

import (
	"example/hello/greetings"
  "example/hello/other/gopher/greet"
	"fmt"
)

func main(){
  fmt.Println(greetings.Hello("Bosco"))
  greet.Hello()
}

