package shared

type Comparator[T any] interface {
	Compare(x, y T) int
}

type IntComparator struct {
}

func (i *IntComparator) Compare(x, y int) int {
	if x == y {
		return 0
	} else if x < y {
		return -1
	} else {
		return 1
	}
}
