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

// Adds any element to the list
func (fl *FastList) Add(element interface{}) bool {

	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}

	fl.elementData = append(fl.elementData, element)
	fl.size++
	return true
}

// Returns the element at index i
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

// Returns the last element from the list and removes it
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
func (fl *FastList) RemoveElement(element interface{}) bool {
	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}
	for index := fl.size - 1; index >= 0; index-- {
		if element == fl.elementData[index] {
			if index == fl.size-1 {
				fl.elementData = fl.elementData[:index]
			} else {
				copy(fl.elementData[index:], fl.elementData[index+1:])
				fl.elementData = fl.elementData[:fl.size-1]
			}
			fl.size--
			return true
		}
	}
	return false
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

func (fl *FastList) removeIndex(index int) interface{} {
	if fl.size == 0 {
		return nil
	}

	if fl.safe {
		fl.Lock()
		defer fl.Unlock()
	}

	old := fl.elementData[index]

	numMoved := fl.size - index - 1
	if numMoved > 0 {
		copy(fl.elementData[:index], fl.elementData[:index-1])
	}
	fl.size--
	return old
}
