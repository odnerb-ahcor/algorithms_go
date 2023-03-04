package main

import (
	"fmt"

	"github.com/BrendoRochaDel/algorithms_go/linearStructures/stacks"
)

func main() {
	fmt.Println("===============STACKS===============")
	stack := stacks.NewStack()

	stack.Push("Brendo")
	fmt.Println(stack.Peek())
	fmt.Println(stack.Size())
	stack.Pop()
	value, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	fmt.Println("\n===============QUEUES==============")
}
