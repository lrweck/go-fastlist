package fastlist

import (
	"testing"
)

func TestGetAll(t *testing.T) {

	fl := NewFastList(true)

	fl.Add(1)
	fl.Add("a")
	fl.Add("x")
	fl.Add(123456)

	var expected []interface{}
	expected = append(expected, 1, "a", "x", 123456)
	index := 0

	for _, j := range fl.GetAll() {
		if j != expected[index] {
			t.Errorf("Expected %v, but got %v\n", expected[index], j)
			index++
		}
		index++
	}

}

func TestSet(t *testing.T) {

	fl := NewFastList(true)

	fl.Add(1)
	fl.Add("a")
	fl.Add("x")
	fl.Add(123456)

	var expected []interface{}
	expected = append(expected, 123456, "x", "a", 1)

	for i, v := range expected {
		tmp := fl.Get(i)

		if x := fl.Set(i, v); x != tmp {
			t.Errorf("Expected %v, but got %v\n", tmp, x)
		}
	}

	x := fl.Set(8, "cba")

	if x == nil {
		if tmp := fl.Get(8); tmp != "cba" {
			t.Errorf("Expected %s, but got %v\n", "cba", tmp)
		}
	}

}

func TestGet(t *testing.T) {

	fl := NewFastList(true)

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

	// Index bigger than max index pos
	_ = fl.Get(len(expected))

	// Negative index
	_ = fl.Get(-1)

}

func TestRemoveLast(t *testing.T) {

	fl := NewFastList(true)

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

	// Test if it's possible to remove a non-existing elem
	_ = fl.RemoveLast()

	if fl.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", fl.Size())
	}

}

func TestRemoveElement(t *testing.T) {

	fl := NewFastList(true)

	fl.Add(1)
	fl.Add("a")
	fl.Add("x")
	fl.Add(123456)

	var expected []interface{}
	expected = append(expected, 1, "a", "x", 123456)

	for _, v := range expected {
		ok := fl.RemoveElement(v)
		if !ok {
			t.Errorf("Error removing element %v\n", v)
		}
	}

	if fl.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", fl.Size())
	}

	hm := fl.RemoveElement("hm")

	if hm {
		t.Errorf("Should not have found this...")
	}

}

func TestRemoveIndex(t *testing.T) {

	fl := NewFastList(true)

	fl.Add(1)
	fl.Add("a")
	fl.Add("x")
	fl.Add(123456)

	var expected []interface{}
	expected = append(expected, 1, "a", "x", 123456)

	// for i, v := range expected {
	// 	if el := fl.RemoveIndex(i); el != v {
	// 		t.Errorf("Expected %v, but got %v\n", v, el)
	// 	}
	// }
	for times := len(expected) - 1; times >= 0; times-- {
		if tmp := fl.RemoveIndex(times); tmp != expected[times] {
			t.Errorf("Expected %v, but got %v\n", expected[times], tmp)
		}
	}

	if fl.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", fl.Size())
	}

	hm := fl.RemoveIndex(18)

	if hm != nil {
		t.Errorf("Should not have found element at index 18...")
	}

	hm = fl.RemoveIndex(-1)

	if hm != nil {
		t.Errorf("Should not have element at index -1...")
	}

}

func TestSize(t *testing.T) {

	fl := NewFastList(true)

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
	fl := NewFastList(true)

	fl.Add(1)
	fl.Add("a")
	fl.Add("x")
	fl.Add(123456)

	fl.Clear()

	if fl.Size() != 0 {
		t.Errorf("Expected cleared slice")
	}

}

// Benchmarks Add Method
func BenchmarkAdd(b *testing.B) {
	fl := NewFastList(true)

	for i := 0; i < b.N; i++ {
		fl.Add(i)
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	fl := NewFastList(false)

	for i := 0; i < 5000; i++ {
		fl.Add(i)
	}

	for i := 0; i < b.N; i++ {
		fl.RemoveElement(i)
	}

	for i := b.N; i >= 0; i-- {
		fl.RemoveElement(i)
	}

}

func BenchmarkRemoveLast(b *testing.B) {
	fl := NewFastList(false)

	for i := 0; i < 5000000; i++ {
		fl.Add(i)
	}

	for i := 0; i < b.N; i++ {
		_ = fl.RemoveLast()
	}

}
