package algorithm

func TopkByPartition[T Comparable](arr []T, k int) []T {
	// fmt.Println("start")
	result := make([]T, k)

	temp := make([]T, len(arr))
	copy(temp, arr)

	partition[T](temp, k, result)

	return result
}

func partition[T Comparable](arr []T, k int, result []T) {
	if k == 0 || len(arr) == 0 {
		return
	}

	if len(arr) <= k {
		copy(result[:], arr[:])
		return
	}

	pivot := 0
	i, j := 1, len(arr)-1
	for i < j {
		for arr[j] > arr[pivot] && j > i {
			j--
		}

		for arr[i] < arr[pivot] && i < j {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	if arr[i] < arr[pivot] {
		arr[i], arr[pivot] = arr[pivot], arr[i]
		pivot = i
	}

	m := len(arr) - pivot
	n := m - 1

	if n > k {
		partition[T](arr[pivot+1:], k, result)
	} else if n == k {
		// result = append(result, arr[i:]...)
		copy(result[:], arr[pivot+1:])
	} else if m == k {
		// result = append(result, arr[i-1:]...)
		copy(result[:], arr[pivot:])
	} else {
		// result = append(result, arr[i-1:]...)
		copy(result[:], arr[pivot:])
		partition[T](arr[:pivot], k-m, result[m:])
	}
}

func TopkByHeap[T Comparable](arr []T, k int) []T {
	if len(arr) <= k {
		return arr
	}

	heap := NewHeap(arr[:k])
	heap.Build()
	for _, ele := range arr[k:] {
		top := heap.Top()
		if ele > top {
			heap.ReplaceTop(ele)
		}
	}

	return heap.GetAll()
}
