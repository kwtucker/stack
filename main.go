package main

import (
	"fmt"
	"time"

	"github.com/kwtucker/Stack/src"
)

var stack = &src.Stack{
	Reversed:     false,
	FunctionName: true,
	Path:         true,
	Marker:       "int",
}

func main() {
	one()
}

func one() {
	fmt.Println(stack.StackIt(2))
	time.Sleep(time.Second * 1)
	two()
}

func two() {
	stack.Marker = "datetime"
	fmt.Println(stack.StackIt(3))
	time.Sleep(time.Second * 1)
	three()
}

func three() {
	fmt.Println(stack.StackIt(4))
	time.Sleep(time.Second * 1)
	four()
}

func four() {
	stack.Marker = "*"
	fmt.Println(stack.StackIt(5))
}
