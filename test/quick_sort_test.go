package test

import (
	"algorithm_practice/algorithm"
	"testing"
)

func TestPartition(t *testing.T) {
	type args[T algorithm.Comparable] struct {
		arr []T
	}
	tests := []struct {
		name string
		args args[int]
	}{
		{
			name: "test1",
			args: args[int]{
				arr: []int{1, 2, 3, 4},
			},
		},
		{
			name: "test2",
			args: args[int]{
				arr: []int{2, 1},
			},
		},
		{
			name: "test3",
			args: args[int]{
				arr: []int{3, 2, 1},
			},
		},
		{
			name: "test4",
			args: args[int]{
				arr: []int{-1, 3, 10, 1, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			algorithm.Partition(tt.args.arr)
			t.Logf("arr: %+v", tt.args.arr)
		})
	}
}
