package algorithm

func Partition[T Comparable](arr []T) {
	if len(arr) <= 1 {
		return
	}

	pivot := 0
	i, j := 1, len(arr)-1
	for i < j {
		for arr[j] >= arr[pivot] && j > i {
			j--
		}

		for arr[i] <= arr[pivot] && i < j {
			i++
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	if arr[i] < arr[pivot] {
		arr[i], arr[pivot] = arr[pivot], arr[i]
	}

	Partition[T](arr[:i])
	Partition[T](arr[i:])
}
