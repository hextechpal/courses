package ll

type Node[T any] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}
