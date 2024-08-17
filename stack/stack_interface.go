package st

type StackInterface[A any] interface {
	Peek() (A, error)
	Pop() (A, error)
	Push(el A)
	IsEmpty() bool
	Print()
	ToString() string
	Size() int
}
