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

func TestNew(t *testing.T) {
	dequeCompare(t, 0, NewDeque(), &Deque{[]interface{}{}, 0, 0})
	dequeCompare(t, 1, NewDeque(1, 2, 3), &Deque{[]interface{}{1, 2, 3}, 3, 0})
	{
		data := []interface{}{1, 2, 3}
		d := NewDeque(data...)
		dequeCompare(t, 2, d, &Deque{[]interface{}{1, 2, 3}, 3, 0})
		data[0] = 7
		arrayCompare(t, 2, d.data, []interface{}{1, 2, 3})
	}
	{
		data := make([]interface{}, 3, 5)
		data[0], data[1], data[2] = 1, 2, 3
		d := NewDeque(data...)
		dequeCompare(t, 3, d, &Deque{[]interface{}{1, 2, 3}, 3, 0})
		data[0] = 7
		arrayCompare(t, 3, d.data, []interface{}{1, 2, 3})
	}
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
		dequeCompare(t, i, test.Deque, test.Want)
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
		dequeCompare(t, i, test.Deque, test.Want)
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
			dequeCompare(t, i, test.Deque, test.Want)
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
			dequeCompare(t, i, test.Deque, test.Want)
		})
	}
}

type concatTest struct {
	Deque  *Deque
	Concat *Deque
	Want   *Deque
}

func TestConcatFront(t *testing.T) {
	tests := []concatTest{
		{
			Deque:  &Deque{[]interface{}{nil, nil, 1, 5, 9, nil, nil, nil}, 3, 2},
			Concat: &Deque{[]interface{}{nil, nil, nil, 3, 1, 4, nil, nil}, 3, 3},
			Want:   &Deque{[]interface{}{1, 4, 1, 5, 9, nil, nil, 3}, 6, 7},
		},
		{
			Deque:  &Deque{[]interface{}{nil, nil, 4, 5}, 2, 2},
			Concat: &Deque{[]interface{}{3, nil, nil, nil, 1, 2}, 3, 4},
			Want:   &Deque{[]interface{}{4, 5, nil, nil, nil, 1, 2, 3}, 5, 5},
		},
		{
			Deque:  &Deque{[]interface{}{5, 6}, 2, 0},
			Concat: &Deque{[]interface{}{3, 4, 1, 2}, 4, 2},
			Want:   &Deque{[]interface{}{5, 6, nil, nil, nil, nil, 1, 2, 3, 4}, 6, 6},
		},
	}

	for i, test := range tests {
		test.Deque.ConcatFront(test.Concat)
		dequeCompare(t, i, test.Deque, test.Want)
	}
}

func TestConcatBack(t *testing.T) {
	tests := []concatTest{
		{
			Deque:  &Deque{[]interface{}{nil, nil, nil, 3, 1, 4, nil, nil}, 3, 3},
			Concat: &Deque{[]interface{}{nil, nil, 1, 5, 9, nil, nil, nil}, 3, 2},
			Want:   &Deque{[]interface{}{9, nil, nil, 3, 1, 4, 1, 5}, 6, 3},
		},
		{
			Deque:  &Deque{[]interface{}{nil, nil, 1, 2}, 2, 2},
			Concat: &Deque{[]interface{}{5, nil, nil, nil, 3, 4}, 3, 4},
			Want:   &Deque{[]interface{}{1, 2, 3, 4, 5, nil, nil, nil}, 5, 0},
		},
		{
			Deque:  &Deque{[]interface{}{1, 2}, 2, 0},
			Concat: &Deque{[]interface{}{5, 6, 3, 4}, 4, 2},
			Want:   &Deque{[]interface{}{1, 2, 3, 4, 5, 6, nil, nil, nil, nil}, 6, 0},
		},
	}

	for i, test := range tests {
		test.Deque.ConcatBack(test.Concat)
		dequeCompare(t, i, test.Deque, test.Want)
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
				shallowCompare(t, i, test.Deque.Front(), test.Array[0])
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
				shallowCompare(t, i, test.Deque.Back(), test.Array[len(test.Array)-1])
			})
		}
	}
}

func TestAt(t *testing.T) {
	for testIndex, test := range arrayTests {
		for i, want := range test.Array {
			checkPanic(t, testIndex, nil, func() {
				shallowCompare(t, testIndex, test.Deque.At(i), want)
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
		arrayCompare(t, i, test.Deque.Array(), test.Array)
	}
}

func TestCopy(t *testing.T) {
	for i, test := range arrayTests {
		data := make([]interface{}, test.Deque.Cap())
		copy(data, test.Array)
		want := &Deque{data, len(test.Array), 0}
		dequeCompare(t, i, test.Deque.Copy(), want)
	}
}

func TestReserve(t *testing.T) {
	tests := []struct {
		Deque *Deque
		Want  *Deque
		N     int
	}{
		{
			Deque: &Deque{[]interface{}{}, 0, 0},
			Want:  &Deque{[]interface{}{}, 0, 0},
			N:     0,
		},
		{
			Deque: &Deque{[]interface{}{}, 0, 0},
			Want:  &Deque{[]interface{}{nil, nil}, 0, 0},
			N:     2,
		},
		{
			Deque: &Deque{[]interface{}{2, 1}, 2, 1},
			Want:  &Deque{[]interface{}{2, 1}, 2, 1},
			N:     2,
		},
		{
			Deque: &Deque{[]interface{}{3, nil, 1, 2}, 3, 2},
			Want:  &Deque{[]interface{}{1, 2, 3, nil, nil, nil, nil, nil}, 3, 0},
			N:     8,
		},
	}

	for i, test := range tests {
		test.Deque.Reserve(test.N)
		dequeCompare(t, i, test.Deque, test.Want)
	}
}

func TestReset(t *testing.T) {
	for i, test := range arrayTests {
		wantCap := test.Deque.Cap()
		test.Deque.Reset()
		shallowCompare(t, i, test.Deque.Len(), 0)
		shallowCompare(t, i, test.Deque.Cap(), wantCap)
	}
}

var lenCapTests = []struct {
	Deque *Deque
	Len   int
	Cap   int
}{
	{Deque: &Deque{[]interface{}{}, 0, 0}, Len: 0, Cap: 0},
	{Deque: &Deque{[]interface{}{3, nil, 1, 2}, 3, 2}, Len: 3, Cap: 4},
	{Deque: &Deque{[]interface{}{4, 1, 3, 1}, 0, 2}, Len: 0, Cap: 4},
}

func TestLen(t *testing.T) {
	for i, test := range lenCapTests {
		shallowCompare(t, i, test.Deque.Len(), test.Len)
	}
}

func TestCap(t *testing.T) {
	for i, test := range lenCapTests {
		shallowCompare(t, i, test.Deque.Cap(), test.Cap)
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		Deque  *Deque
		String string
	}{
		{
			Deque:  &Deque{[]interface{}{1, 2, 3, 4, 5}, 5, 0},
			String: "[1 2 3 4 5]",
		},
		{
			Deque:  &Deque{[]interface{}{3, nil, nil, "A", "B", "C", 1, 2}, 6, 3},
			String: "[A B C 1 2 3]",
		},
		{
			Deque:  &Deque{[]interface{}{}, 0, 0},
			String: "[]",
		},
		{
			Deque:  &Deque{[]interface{}{4, nil, 3, 1}, 0, 2},
			String: "[]",
		},
	}

	for i, test := range tests {
		shallowCompare(t, i, test.Deque.String(), test.String)
	}
}

func shallowCompare(t *testing.T, testIndex int, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("test %d: got %v, want %v", testIndex, got, want)
	}
}

func deepCompare(t *testing.T, testIndex int, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("test %d: got %v, want %v", testIndex, got, want)
	}
}

func dequeCompare(t *testing.T, testIndex int, got, want *Deque) {
	t.Helper()
	if got.len != want.len || got.front != want.front || !arrayEqual(got.data, want.data) {
		t.Errorf("test %d: got {%v %d %d}, want {%v %d %d}",
			testIndex, got.data, got.len, got.front, want.data, want.len, want.front)
	}
}

func arrayCompare(t *testing.T, testIndex int, got, want []interface{}) {
	t.Helper()
	if !arrayEqual(got, want) {
		t.Errorf("test %d: got %v, want %v", testIndex, got, want)
	}
}

func arrayEqual(a, b []interface{}) bool {
	if len(a) != len(b) || cap(a) != cap(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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
