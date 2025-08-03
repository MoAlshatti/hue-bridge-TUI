package bridge

type Incrementable interface {
	Increment()
	Decrement()
}

// The only reason this is used is because i fancy using interfaces
func Increment_cursor(i Incrementable) {
	i.Increment()
}
func Decrement_cusror(i Incrementable) {
	i.Decrement()
}
