package bst

type Valuer interface {
	GetValue() int
}
type Comparer interface {
	Compare(c Valuer) int
	Valuer
}
