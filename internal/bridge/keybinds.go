package bridge

type Incrementable interface {
	increment()
	decrement()
}

// The only reason this is used is because i fancy using interfaces
func Increment_cursor(i Incrementable) {
	i.increment()
}
func Decrement_cusror(i Incrementable) {
	i.decrement()
}
