package main

import (
	"fmt"

	"github.com/BrendoRochaDel/algorithms_go/linearStructures/queues"
	"github.com/BrendoRochaDel/algorithms_go/linearStructures/stacks"
)

func main() {
	fmt.Println("===============STACKS===============")
	stack := stacks.NewStack()

	stack.Push("example")
	fmt.Println(stack.Peek())
	stack.Pop()
	value, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	fmt.Println("\n===============QUEUES==============")
	queue := queues.NewQueue()

	queue.Enqueue("example1")
	queue.Enqueue("example2")
	queue.Enqueue("example3")
	fmt.Println(queue.Dequeue())
	queue.Enqueue("example4")
	err = queue.Enqueue("exampleError")
	fmt.Println(err)
}
