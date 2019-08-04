package deque

import (
	"errors"
	"fmt"
	"strings"
)

// Deque is a double-ended queue implemented as using circular array.
type Deque struct {
	data  []interface{}
	len   int
	front int
}

// ErrEmpty is panicked when attempting to access or remove an element
// in an empty deque.
var ErrEmpty = errors.New("deque: deque is empty")

// ErrIndexBounds is panicked by At with an invalid index.
var ErrIndexBounds = errors.New("deque: index out of range")

// PushFront pushes the element x onto the front of the deque.
// The complexity is amortized O(1).
func (d *Deque) PushFront(x interface{}) {
	d.realloc(1)
	d.front = (d.front + cap(d.data) - 1) % cap(d.data)
	d.data[d.front] = x
	d.len++
}

// PushBack pushes the element x onto the back of the deque.
// The complexity is amortized O(1).
func (d *Deque) PushBack(x interface{}) {
	d.realloc(1)
	d.data[(d.front+d.len)%cap(d.data)] = x
	d.len++
}

// PopFront removes and returns the first element from the deque.
// The complexity is O(1).
// It panics if the deque is empty.
func (d *Deque) PopFront() interface{} {
	d.checkEmpty()
	d.len--
	x := d.data[d.front]
	d.data[d.front] = nil
	d.front = (d.front + 1) % cap(d.data)
	return x
}

// PopBack removes and returns the last element from the deque.
// The complexity is O(1).
// It panics if the deque is empty.
func (d *Deque) PopBack() interface{} {
	d.checkEmpty()
	d.len--
	back := (d.front + d.len) % cap(d.data)
	x := d.data[back]
	d.data[back] = nil
	return x
}

// ConcatFront pushes all elements of deque d2 to the front of deque d.
func (d *Deque) ConcatFront(d2 *Deque) {
	d.realloc(d2.len)
	d.front = (d.front + cap(d.data) - d2.len) % cap(d.data)
	for i := 0; i < d2.len; i++ {
		d.data[(d.front+i)%cap(d.data)] = d2.data[(d2.front+i)%cap(d2.data)]
	}
	d.len += d2.len
}

// ConcatBack pushes all elements of deque d2 to the back of deque d.
func (d *Deque) ConcatBack(d2 *Deque) {
	d.realloc(d2.len)
	for i := 0; i < d2.len; i++ {
		d.data[(d.front+d.len+i)%cap(d.data)] = d2.data[(d2.front+i)%cap(d2.data)]
	}
	d.len += d2.len
}

// Front returns the first element of the deque.
// It panics if the deque is empty.
func (d *Deque) Front() interface{} {
	d.checkEmpty()
	return d.data[d.front]
}

// Back returns the last element of the deque.
// It panics if the deque is empty.
func (d *Deque) Back() interface{} {
	d.checkEmpty()
	return d.data[(d.front+d.len+cap(d.data)-1)%cap(d.data)]
}

// At returns the element at index i in the deque.
// It panics if i is out of bounds.
func (d *Deque) At(i int) interface{} {
	if i < 0 || i >= d.len {
		panic(ErrIndexBounds)
	}
	return d.data[(d.front+i)%cap(d.data)]
}

// Array copies all elements of deque d into an array.
func (d *Deque) Array() []interface{} {
	a := make([]interface{}, d.len)
	d.fillArray(a)
	return a
}

// Copy makes a copy of deque d with the same elements and capacity.
func (d *Deque) Copy() *Deque {
	data := make([]interface{}, cap(d.data))
	d.fillArray(data)
	return &Deque{data, d.len, 0}
}

// Reserve resizes the underlying array of the deque to at least n.
func (d *Deque) Reserve(n int) {
	if n <= cap(d.data) {
		return
	}
	data := make([]interface{}, n)
	for i := 0; i < d.len; i++ {
		data[i] = d.data[(d.front+i)%cap(d.data)]
	}
	d.data = data
	d.front = 0
}

// Reset resets the deque to be empty, but it retains the underlying array.
func (d *Deque) Reset() {
	d.len = 0
	d.front = 0
}

// Len returns the number of elements in the deque.
func (d *Deque) Len() int {
	return d.len
}

// Cap returns the number of elements allocated in the deque.
func (d *Deque) Cap() int {
	return cap(d.data)
}

func (d *Deque) String() string {
	var b strings.Builder
	b.WriteByte('[')
	for i := 0; i < d.len; i++ {
		if i != 0 {
			b.WriteByte(' ')
		}
		b.WriteString(fmt.Sprint(d.data[(d.front+i)%cap(d.data)]))
	}
	b.WriteByte(']')
	return b.String()
}

// realloc enlarges d to fit an additional n items.
func (d *Deque) realloc(n int) {
	switch {
	case d.len+n <= cap(d.data):
	case d.len+n <= cap(d.data)*2:
		d.Reserve(cap(d.data) * 2)
	default:
		d.Reserve(d.len + n*2)
	}
}

// checkEmpty panics if d is empty.
func (d *Deque) checkEmpty() {
	if d.len == 0 {
		panic(ErrEmpty)
	}
}

// fillArray copies all elements of deque to array a starting at index 0 of a.
func (d *Deque) fillArray(a []interface{}) {
	for i := 0; i < d.len; i++ {
		a[i] = d.data[(d.front+i)%cap(d.data)]
	}
}
