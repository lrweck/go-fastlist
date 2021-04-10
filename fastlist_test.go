package fastlist

import (
	"testing"
)

// type testFastList struct {
// 	elem []interface{}
// 	size int
// }

// func init() {
// 	fl.Add(1)
// 	fl.Add("a")
// 	fl.Add("x")
// 	fl.Add(12456)

// }

func TestGetAll(t *testing.T) {

	fl := NewFastList()

	fl.Add(1)
	fl.Add("a")
	fl.Add("x")
	fl.Add(123456)

	var expected []interface{}
	expected = append(expected, 1, "a", "x", 123456)
	index := 0

	for _, j := range fl.GetAll() {
		// t.Errorf("fl: %v - %d", j, i)
		// t.Errorf("ex: %v - %d", expected[index], index)
		if j != expected[index] {
			t.Errorf("Expected %v, but got %v\n", expected[index], j)
			index++
		}
		index++
	}
}

func TestGet(t *testing.T) {

	fl := NewFastList()

	fl.Add(1)
	fl.Add("a")
	fl.Add("x")
	fl.Add(123456)

	var expected []interface{}
	expected = append(expected, 1, "a", "x", 123456)

	for i, j := range expected {

		if j != fl.Get(i) {
			t.Errorf("Expected %v, but got %v\n", j, fl.Get(i))
		}
	}
}

func TestRemoveLast(t *testing.T) {

	fl := NewFastList()

	fl.Add(1)
	fl.Add("a")
	fl.Add("x")
	fl.Add(123456)

	var expected []interface{}
	expected = append(expected, 1, "a", "x", 123456)

	for i := len(expected) - 1; i >= 0; i-- {
		if x := fl.RemoveLast(); expected[i] != x {
			t.Errorf("Expected %v, but got %v\n", expected[i], x)
		}
	}

	if fl.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", fl.Size())
	}

}

func TestSize(t *testing.T) {

	fl := NewFastList()

	fl.Add(1)
	fl.Add("a")
	fl.Add("x")
	fl.Add(123456)

	var expected []interface{}
	expected = append(expected, 1, "a", "x", 123456)

	sizeTrack := len(expected)
	for i := len(expected) - 1; i >= 0; i-- {
		fl.RemoveLast()
		sizeTrack--
		if sizeTrack != fl.Size() {
			t.Errorf("Expected size %d, but got %d", sizeTrack, fl.Size())
		}
	}

	if fl.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", fl.Size())
	}

}

func TestClear(t *testing.T) {
	fl := NewFastList()

	fl.Add(1)
	fl.Add("a")
	fl.Add("x")
	fl.Add(123456)

	fl.Clear()

	if fl.Size() != 0 {
		t.Errorf("Expected cleared slice")
	}

}
