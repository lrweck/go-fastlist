package fastlist

import "sync"

type FastList struct {
	sync.RWMutex
	elementData []interface{}
	size        int
	safe        bool
}

// Creates a new fastlist. Pass true as a parameter to create a concurrency-safe
// list, or false otherwise.
// Returns a FastList
func NewFastList(isSafe bool) FastList {
	return FastList{
		elementData: make([]interface{}, 0),
		size:        0,
		safe:        isSafe,
	}
}

// Adds any element to the list and returns true if added succesfuly
func (fl *FastList) Add(element interface{}) bool {

	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}

	fl.elementData = append(fl.elementData, element)
	fl.size++
	return true
}

// Returns the element at the specified index or nil if not found
func (fl *FastList) Get(index int) interface{} {

	if fl.size <= 0 || index > fl.size-1 || index < 0 {
		return nil
	}

	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}

	return fl.elementData[index]
}

// Returns the last element from the list and removes it.
// Returns nil if the list is empty
func (fl *FastList) RemoveLast() interface{} {

	if fl.size == 0 {
		return nil
	}

	// should it be before the size == 0 check?
	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}

	fl.size--
	elem := fl.elementData[fl.size]
	fl.elementData = fl.elementData[:fl.size]
	return elem
}

// Returns all the elements
func (fl *FastList) GetAll() []interface{} {
	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}
	return fl.elementData
}

// Removes the element at index i
// Most efficient at removing from the end of the list
// Returns true if successfully removed and false if not
func (fl *FastList) RemoveElement(element interface{}) bool {
	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}
	for index := fl.size - 1; index >= 0; index-- {
		if element == fl.elementData[index] {
			remElement(fl, index)
			return true
		}
	}
	return false
}

func remElement(fl *FastList, index int) {
	if index == fl.size-1 {
		fl.elementData = fl.elementData[:index]
	} else {
		copy(fl.elementData[index:], fl.elementData[index+1:])
		fl.elementData = fl.elementData[:fl.size-1]
	}
	fl.size--
}

// Clears the list (removes all elements)
func (fl *FastList) Clear() {
	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}
	fl.elementData = fl.elementData[:0]
	fl.size = 0
}

// Returns the size of the list
func (fl *FastList) Size() int {
	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}
	return fl.size
}

// Sets an element at the index specified and returns the element
// that was located at the specified location.
// Returns nil if the specified location was empty or did not exist
func (fl *FastList) Set(index int, element interface{}) interface{} {

	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}

	for fl.size <= index {
		fl.elementData = append(fl.elementData, nil)
		fl.size++
	}

	old := fl.elementData[index]
	fl.elementData[index] = element
	return old
}

// Removes an element at the specified index and returns it.
// Returns nil if element does not exist
func (fl *FastList) RemoveIndex(index int) interface{} {
	if fl.size <= 0 || index > fl.size-1 || index < 0 {
		return nil
	}

	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}

	old := fl.elementData[index]

	remElement(fl, index)

	return old
}
