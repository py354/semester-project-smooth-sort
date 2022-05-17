package sort

func HeapSort(arr []int) {
	heap := newMaxHeap(arr)
	heap.sort(len(arr))
}

type maxHeap struct {
	arr []int
}

func newMaxHeap(arr []int) *maxHeap {
	return &maxHeap{arr: arr}
}

func (m *maxHeap) leftChildIndex(index int) int {
	return 2*index + 1
}

func (m *maxHeap) rightChildIndex(index int) int {
	return 2*index + 2
}

func (m *maxHeap) swap(first, second int) {
	temp := m.arr[first]
	m.arr[first] = m.arr[second]
	m.arr[second] = temp
}

func (m *maxHeap) leaf(index int, size int) bool {
	if index >= (size/2) && index <= size {
		return true
	}
	return false
}

func (m *maxHeap) downHeapify(current int, size int) {
	if m.leaf(current, size) {
		return
	}
	largest := current
	leftChildIndex := m.leftChildIndex(current)
	rightRightIndex := m.rightChildIndex(current)
	if leftChildIndex < size && m.arr[leftChildIndex] > m.arr[largest] {
		largest = leftChildIndex
	}
	if rightRightIndex < size && m.arr[rightRightIndex] > m.arr[largest] {
		largest = rightRightIndex
	}
	if largest != current {
		m.swap(current, largest)
		m.downHeapify(largest, size)
	}
	return
}

func (m *maxHeap) buildMinHeap(size int) {
	for index := (size / 2) - 1; index >= 0; index-- {
		m.downHeapify(index, size)
	}
}

func (m *maxHeap) sort(size int) {
	m.buildMinHeap(size)
	for i := size - 1; i > 0; i-- {
		m.swap(0, i)
		m.downHeapify(0, i)
	}
}
