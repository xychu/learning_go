package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(i)
		fmt.Print(s)
		fmt.Println(i)
	}
}

func main() {
	say("R U ready?")
	go say("hello")
	go say("world")
	say("www")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
