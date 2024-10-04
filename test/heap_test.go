package test

import (
	"algorithm_practice/algorithm"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := algorithm.NewHeap([]int{34, 2, 67, 56, 90, 1, -10})
	heap.Build()
	heap.Push(66)
	heap.Push(11)
	for heap.Size() > 0 {
		value := heap.Pop()
		fmt.Println("value: ", value)
	}
}
