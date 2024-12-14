package datastructures

type Counter[T comparable] map[T]int

func NewCounter[T comparable]() Counter[T] {
	return make(Counter[T])
}

func (c Counter[T]) Inc(key T) {
	c[key]++
}

func (c Counter[T]) Dec(key T) {
	c[key]--
	if c[key] == 0 {
		delete(c, key)
	}
}

func (c Counter[T]) Get(key T) int {
	return c[key]
}
