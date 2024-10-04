package test

import (
	"algorithm_practice/algorithm"
	"testing"
)

func TestJaccardForSortedCollection(t *testing.T) {
	type args[T algorithm.Comparable] struct {
		collection1 []T
		collection2 []T
	}
	tests := []struct {
		name string
		args args[int]
		want float64
	}{
		{
			name: "test",
			args: args[int]{
				collection1: []int{1, 3, 4, 7, 9, 12, 23},
				collection2: []int{2, 3, 5, 9, 23, 27, 31},
			},
			want: 0.27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := algorithm.JaccardForSortedCollection(tt.args.collection1, tt.args.collection2); got != tt.want {
				t.Errorf("JaccardForSortedCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
