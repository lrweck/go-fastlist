package fastlist

type FastList struct {
	elementData []interface{}
	size        int
}

func NewFastList() FastList {
	return FastList{
		make([]interface{}, 0),
		0,
	}
}

func (fl *FastList) Add(element interface{}) bool {
	fl.elementData = append(fl.elementData, element)
	fl.size++
	return true
}

func (fl *FastList) Get(index int) interface{} {
	return fl.elementData[index]
}

func (fl *FastList) RemoveLast() interface{} {
	fl.size--
	elem := fl.elementData[fl.size]
	fl.elementData = fl.elementData[:fl.size]
	return elem
}

func (fl *FastList) GetAll() []interface{} {
	return fl.elementData
}

func (fl *FastList) RemoveElement(element interface{}) bool {
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

func (fl *FastList) Clear() {

	fl.elementData = fl.elementData[:0]
	fl.size = 0
}

func (fl *FastList) Size() int {
	return fl.size
}

func (fl *FastList) Set(index int, element interface{}) interface{} {

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

	old := fl.elementData[index]

	numMoved := fl.size - index - 1
	if numMoved > 0 {
		copy(fl.elementData[:index], fl.elementData[:index-1])
	}
	fl.size--
	return old
}
