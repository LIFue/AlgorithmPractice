package algorithm

type Heap[T Comparable] struct {
	arr []T
}

func NewHeap[T Comparable](arr []T) *Heap[T] {
	h := &Heap[T]{
		arr: make([]T, len(arr)),
	}
	copy(h.arr, arr)
	return h
}

// 堆的基础操作之一：向下调整
func (heap *Heap[T]) DownardAdjust(parent int) {
	left := 2*parent + 1
	if left >= len(heap.arr) {
		return
	}

	minIndex := parent
	if heap.arr[left] < heap.arr[minIndex] {
		minIndex = left
	}

	right := 2*parent + 2
	if right < len(heap.arr) && heap.arr[right] < heap.arr[minIndex] {
		minIndex = right
	}

	if minIndex != parent {
		heap.arr[minIndex], heap.arr[parent] = heap.arr[parent], heap.arr[minIndex]
		heap.DownardAdjust(minIndex)
	}
}

// 堆的基础操作之一：向下调整
func (heap *Heap[T]) UpwardAdjust(child int) {
	if child == 0 || child >= len(heap.arr) {
		return
	}
	parent := (child - 1) / 2
	left := parent*2 + 1
	right := parent*2 + 2

	minIndex := parent
	if left < len(heap.arr) && heap.arr[left] < heap.arr[minIndex] {
		minIndex = left
	}
	if right < len(heap.arr) && heap.arr[right] < heap.arr[minIndex] {
		minIndex = right
	}

	if minIndex != parent {
		heap.arr[parent], heap.arr[minIndex] = heap.arr[minIndex], heap.arr[parent]
		heap.UpwardAdjust(parent)
	}
}

func (heap *Heap[T]) Build() {
	if len(heap.arr) <= 1 {
		return
	}

	n := len(heap.arr)
	lastRightNodeIndex := n / 2 * 2

	for i := lastRightNodeIndex; i > 0; i -= 2 {
		subParent := (i - 1) / 2
		heap.DownardAdjust(subParent)
	}
}

func (heap *Heap[T]) Push(ele T) {
	heap.arr = append(heap.arr, ele)
	heap.UpwardAdjust(len(heap.arr) - 1)
}

func (heap *Heap[T]) Pop() T {
	// fmt.Println("before: ", heap.arr)
	var value T
	if len(heap.arr) <= 0 {
		return value
	}

	value = heap.arr[0]
	heap.arr[0] = heap.arr[len(heap.arr)-1]
	heap.arr = heap.arr[:len(heap.arr)-1]
	// fmt.Println("after: ", heap.arr)
	heap.DownardAdjust(0)
	return value
}

func (heap Heap[T]) Size() int {
	return len(heap.arr)
}
