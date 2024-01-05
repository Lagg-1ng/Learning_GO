package main

import "fmt"

const (
	one = iota + 1 // укажите здесь формулу с iota
	three
	five
	seven
	nine
	eleven
)

func third() {
	fmt.Println(one, three, five, seven, nine, eleven)
}
