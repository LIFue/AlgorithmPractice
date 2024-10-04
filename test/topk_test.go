package test

import (
	"algorithm_practice/algorithm"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestTopkByPartition(t *testing.T) {
	type args[T algorithm.Comparable] struct {
		arr []T
		k   int
	}
	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "test1",
			args: args[int]{
				arr: []int{3, 1, 9, 5, 2, 7},
				k:   3,
			},
			want: []int{5, 9, 7},
		},
		{
			name: "test2",
			args: args[int]{
				arr: []int{1},
				k:   1,
			},
			want: []int{1},
		},
		{
			name: "test3",
			args: args[int]{
				arr: []int{5, 9, 5, 8, 9, 3, 2, 10, 39, 54, -1, 233, 44, 32, 8},
				k:   2,
			},
			want: []int{54, 233},
		},
		{
			name: "test4",
			args: args[int]{
				arr: []int{5, 9, 5, 8, 9, 3, 2, 10, 39, 54, -1, 233, 44, 32, 8},
				k:   10,
			},
			want: []int{9, 8, 10, 39, 54, 9, 233, 44, 32, 8},
		},
		{
			name: "test5",
			args: args[int]{
				arr: []int{5, 9, 5, 8, 9, 3, 2, 10, 39, 54, -1, 233, 44, 32, 8},
				k:   11,
			},
			want: []int{5, 9, 8, 10, 39, 54, 9, 233, 44, 32, 8},
		},
		{
			name: "test6",
			args: args[int]{
				arr: []int{5, 9, 5, 8, 9, 3, 2, 10, 39, 54, -1, 233, 44, 32, 8},
				k:   12,
			},
			want: []int{5, 9, 8, 5, 10, 39, 54, 9, 233, 44, 32, 8},
		},
		{
			name: "test7",
			args: args[int]{
				arr: []int{5, 9, 5, 8, 9, 3, 2, 10, 39, 54, -1, 233, 44, 32, 8},
				k:   14,
			},
			want: []int{3, 2, 5, 5, 9, 8, 10, 39, 54, 9, 233, 44, 32, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := algorithm.TopkByPartition(tt.args.arr, tt.args.k)
			// slices.Sort[int](got)
			// slices.Sort[int](tt.want)
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})

			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i] < tt.want[j]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopkByPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopkByPartitionByRandList(t *testing.T) {

	for i := 0; i < 10000; i++ {
		listLen := rand.Int31n(10000) + 1
		k := rand.Int31n(listLen) + 1

		arr := make([]int32, listLen)
		for j := 0; j < int(listLen); j++ {
			arr[j] = rand.Int31() - 5000
		}

		topK := algorithm.TopkByPartition(arr, int(k))

		min := topK[0]
		for _, t := range topK {
			if min > t {
				min = t
			}
		}

		small, equal := 0, 0
		for _, ele := range arr {
			if ele < min {
				small++
				continue
			}
			if ele == min {
				equal++
			}
		}

		if small > int(listLen)-int(k) || small+equal < int(listLen)-int(k)+1 {
			t.Errorf("\nlen(topk): %d, k: %d \n len(list): %d, list: %d\nsmall: %d, equal: %d\n listLen - k: %d, ", len(topK), k, len(arr), listLen, small, equal, listLen-k)
			t.Fail()
		}
	}
}

func TestHeapTopkByPartitionByRandList(t *testing.T) {

	for i := 0; i < 1000; i++ {
		listLen := rand.Int31n(10000) + 1
		k := rand.Int31n(listLen) + 1

		// listLen := 20
		// k := 10

		arr := make([]int32, listLen)
		for j := 0; j < int(listLen); j++ {
			arr[j] = rand.Int31() - 5000
		}

		topK := algorithm.TopkByHeap(arr, int(k))

		min := topK[0]
		for _, t := range topK {
			if min > t {
				min = t
			}
		}

		small, equal := 0, 0
		for _, ele := range arr {
			if ele < min {
				small++
				continue
			}
			if ele == min {
				equal++
			}
		}

		if small > int(listLen)-int(k) || small+equal < int(listLen)-int(k)+1 {
			t.Errorf("\nlen(topk): %d, k: %d \n len(list): %d, list: %d\nsmall: %d, equal: %d\n listLen - k: %d, ", len(topK), k, len(arr), listLen, small, equal, listLen-k)
			t.Fail()
		}
	}
}
