package deque

import (
	"reflect"
	"testing"
)

type pushTest struct {
	Deque *Deque
	Want  *Deque
	Val   interface{}
}

func TestPushFront(t *testing.T) {
	tests := []pushTest{
		{
			Deque: &Deque{[]interface{}{nil, 1, 4, 1}, 3, 1},
			Want:  &Deque{[]interface{}{3, 1, 4, 1}, 4, 0},
			Val:   3,
		},
		{
			Deque: &Deque{[]interface{}{2, 3, nil, nil}, 2, 0},
			Want:  &Deque{[]interface{}{2, 3, nil, 1}, 3, 3},
			Val:   1,
		},
		{
			Deque: &Deque{[]interface{}{1, 0, 3, 2}, 4, 2},
			Want:  &Deque{[]interface{}{3, 2, 1, 0, nil, nil, nil, 4}, 5, 7},
			Val:   4,
		},
	}

	for i, test := range tests {
		test.Deque.PushFront(test.Val)
		deepEqual(t, i, test.Deque, test.Want)
	}
}

func TestPushBack(t *testing.T) {
	tests := []pushTest{
		{
			Deque: &Deque{[]interface{}{3, 1, 4, nil}, 3, 0},
			Want:  &Deque{[]interface{}{3, 1, 4, 1}, 4, 0},
			Val:   1,
		},
		{
			Deque: &Deque{[]interface{}{2, nil, nil, 1}, 2, 3},
			Want:  &Deque{[]interface{}{2, 3, nil, 1}, 3, 3},
			Val:   3,
		},
		{
			Deque: &Deque{[]interface{}{2, 1, 4, 3}, 4, 2},
			Want:  &Deque{[]interface{}{4, 3, 2, 1, 0, nil, nil, nil}, 5, 0},
			Val:   0,
		},
	}

	for i, test := range tests {
		test.Deque.PushBack(test.Val)
		deepEqual(t, i, test.Deque, test.Want)
	}
}

type popTest struct {
	Deque *Deque
	Want  *Deque
	Panic interface{}
}

func TestPopFront(t *testing.T) {
	tests := []popTest{
		{
			Deque: &Deque{[]interface{}{6, 4, 2, nil}, 3, 0},
			Want:  &Deque{[]interface{}{nil, 4, 2, nil}, 2, 1},
		},
		{
			Deque: &Deque{[]interface{}{3, 4, 1, 2}, 4, 2},
			Want:  &Deque{[]interface{}{3, 4, nil, 2}, 3, 3},
		},
		{
			Deque: &Deque{[]interface{}{nil, nil}, 0, 1},
			Panic: ErrEmpty,
		},
	}

	for i, test := range tests {
		checkPanic(t, i, test.Panic, func() {
			test.Deque.PopFront()
			deepEqual(t, i, test.Deque, test.Want)
		})
	}
}

func TestPopBack(t *testing.T) {
	tests := []popTest{
		{
			Deque: &Deque{[]interface{}{nil, 2, 7, 1}, 3, 1},
			Want:  &Deque{[]interface{}{nil, 2, 7, nil}, 2, 1},
		},
		{
			Deque: &Deque{[]interface{}{3, 4, 1, 2}, 4, 2},
			Want:  &Deque{[]interface{}{3, nil, 1, 2}, 3, 2},
		},
		{
			Deque: &Deque{[]interface{}{nil, nil}, 0, 1},
			Panic: ErrEmpty,
		},
	}

	for i, test := range tests {
		checkPanic(t, i, test.Panic, func() {
			test.Deque.PopBack()
			deepEqual(t, i, test.Deque, test.Want)
		})
	}
}

var arrayTests = []struct {
	Deque *Deque
	Array []interface{}
}{
	{
		Deque: &Deque{[]interface{}{1, 2, 3, nil}, 3, 0},
		Array: []interface{}{1, 2, 3},
	},
	{
		Deque: &Deque{[]interface{}{3, 4, 1, 2}, 4, 2},
		Array: []interface{}{1, 2, 3, 4},
	},
	{
		Deque: &Deque{[]interface{}{8, nil, nil, 2, 7, 1, 8, 2}, 6, 3},
		Array: []interface{}{2, 7, 1, 8, 2, 8},
	},
	{
		Deque: &Deque{[]interface{}{}, 0, 0},
		Array: []interface{}{},
	},
	{
		Deque: &Deque{[]interface{}{4, nil, 3, 1}, 0, 2},
		Array: []interface{}{},
	},
}

func TestFront(t *testing.T) {
	for i, test := range arrayTests {
		if len(test.Array) == 0 {
			checkPanic(t, i, ErrEmpty, func() {
				test.Deque.Front()
			})
		} else {
			checkPanic(t, i, nil, func() {
				shallowEqual(t, i, test.Deque.Front(), test.Array[0])
			})
		}
	}
}

func TestBack(t *testing.T) {
	for i, test := range arrayTests {
		if len(test.Array) == 0 {
			checkPanic(t, i, ErrEmpty, func() {
				test.Deque.Back()
			})
		} else {
			checkPanic(t, i, nil, func() {
				shallowEqual(t, i, test.Deque.Back(), test.Array[len(test.Array)-1])
			})
		}
	}
}

func TestAt(t *testing.T) {
	for testIndex, test := range arrayTests {
		for i, want := range test.Array {
			checkPanic(t, testIndex, nil, func() {
				shallowEqual(t, testIndex, test.Deque.At(i), want)
			})
		}
		checkPanic(t, testIndex, ErrIndexBounds, func() {
			test.Deque.At(-1)
		})
		checkPanic(t, testIndex, ErrIndexBounds, func() {
			test.Deque.At(len(test.Array))
		})
	}
}

func TestArray(t *testing.T) {
	for i, test := range arrayTests {
		deepEqual(t, i, test.Deque.Array(), test.Array)
	}
}

func TestCopy(t *testing.T) {
	for i, test := range arrayTests {
		data := make([]interface{}, test.Deque.Cap())
		copy(data, test.Array)
		want := &Deque{data, len(test.Array), 0}
		deepEqual(t, i, test.Deque.Copy(), want)
	}
}

func shallowEqual(t *testing.T, testIndex int, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("test %d: got %v, want %v", testIndex, got, want)
	}
}

func deepEqual(t *testing.T, testIndex int, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("test %d: got %v, want %v", testIndex, got, want)
	}
}

func checkPanic(t *testing.T, testIndex int, want interface{}, mightPanic func()) {
	t.Helper()
	defer func() {
		t.Helper()
		if r := recover(); r != want {
			t.Errorf("test %d: got panic %v, want panic %v", testIndex, r, want)
		}
	}()
	mightPanic()
}
