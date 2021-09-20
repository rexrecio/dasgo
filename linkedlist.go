//Circular linkedlist implementation
package dasgo

import "sync"

type node struct {
	next *node
	data interface{}
}
type LinkedList struct {
	tail     *node
	IsAppend bool //add new element at the end of list?
	mutex    sync.Mutex
}

//User define function to call for each element in the list
//First paramater is the data element, second parameter is additional in case we want to pass more info
type udf func(interface{}, interface{})

func (list *LinkedList) IsEmpty() bool {
	if list.tail == nil {
		return true
	} else {
		return false
	}
}
func (list *LinkedList) Add(newdata interface{}) {
	newnode := new(node)
	newnode.data = newdata
	list.mutex.Lock()
	if list.IsEmpty() {
		newnode.next = newnode
		list.tail = newnode
	} else {
		newnode.next = list.tail.next
		list.tail.next = newnode

		//If you want to add the new element at the end of the list
		if list.IsAppend {
			list.tail = list.tail.next
		}
	}
	list.mutex.Unlock()
}
func (list *LinkedList) Count() int {
	if list.IsEmpty() {
		return 0
	}

	var count int
	list.mutex.Lock()
	for current := list.tail; ; {
		current = current.next
		count++
		if current == list.tail {
			break
		}
	}
	list.mutex.Unlock()
	return count
}
func (list *LinkedList) SetAppend() {
	list.IsAppend = true
}
func (list *LinkedList) Traverse(f udf, param interface{}) {
	if list.IsEmpty() {
		return
	}
	list.mutex.Lock()
	head := list.tail.next
	for current := head; ; {
		f(current.data, param)
		current = current.next
		if current == head {
			break
		}
	}
	list.mutex.Unlock()
}
func (list *LinkedList) DeleteHead() interface{} {
	//Deleting the head is faster than deleting the tail because we need not traverse the whole list
	if list.IsEmpty() {
		return nil
	}
	list.mutex.Lock()
	head := list.tail.next
	returndata := head.data
	if head == list.tail {
		list.tail = nil
	} else {
		list.tail.next = head.next
	}
	list.mutex.Unlock()
	return returndata
}
func (list *LinkedList) DeleteTail() interface{} {
	//Deleting the tail is slower than deleteing the head because we need to traverse the whole list
	if list.IsEmpty() {
		return nil
	}
	list.mutex.Lock()
	returndata := list.tail.data
	/*locate the node prior to the tail*/
	previous := list.tail
	for ; previous.next != list.tail; previous = previous.next {
	}
	if previous == list.tail {
		//If only one element in the list
		list.tail = nil
	} else {
		//General case of more than one element
		previous.next = list.tail.next
		list.tail = previous
	}
	list.mutex.Unlock()
	return returndata
}
