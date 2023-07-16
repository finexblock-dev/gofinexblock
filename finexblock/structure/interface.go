package structure

type Heap[T any] interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Pop() T
	Push(x T)
}