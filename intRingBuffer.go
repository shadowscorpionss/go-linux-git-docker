package main

import (
	"errors"
)

const (
	BUFFER_OVERFLOW_ERROR = "Buffer Overflow"
	EMPTY_BUFFER_ERROR    = "Empty Buffer"
)

type IntRingBuffer struct {
	head     *IntBufferNode //Head node
	tail     *IntBufferNode //Tail node
	size     int
	capacity int
}

// Constructor
func NewIntRingBuffer(capacity int) *IntRingBuffer {
	return &IntRingBuffer{capacity: capacity, size: 0}
}

// Addition of Elements
func (cb *IntRingBuffer) Add(item int) (err error) {
	//default result
	err = nil

	// Size of buffer increases as elements
	// are added to the Linked List
	cb.size++

	// Checking if the buffer is full
	if cb.size == cb.capacity {
		err = errors.New(BUFFER_OVERFLOW_ERROR)
		return
	}

	// Checking if the buffer is empty
	if cb.head == nil {
		cb.head = NewIntBufferNode(item)
		cb.tail = cb.head
		return
	}

	// Node element to be linked
	temp := NewIntBufferNode(item)

	// Referencing the last element to the head node
	temp.Next(cb.head)

	// Updating the tail reference to the
	// latest node added
	cb.tail.Next(temp)

	// Updating the tail to the latest node added
	cb.tail = temp

	return
}

// Retrieving the head element
func (cb *IntRingBuffer) Get() (res int, err error) {
	// Getting the element
	res, err = cb.Peek()
	if err != nil {
		return
	}

	// Updating the head pointer
	cb.head = cb.head.next

	// Updating the tail reference to
	// the new head pointer
	cb.tail.Next(cb.head)

	// Decrementing the size
	cb.size--
	if cb.IsEmpty() {
		cb.innerClear()
	}

	return
}

// Retrieving the head element without deleting
func (cb *IntRingBuffer) Peek() (res int, err error) {
	//default result
	res, err = 0, nil

	// Checking if the buffer is empty
	if cb.IsEmpty() {
		err = errors.New(EMPTY_BUFFER_ERROR)
		return
	}
	// Getting the element
	res = cb.head.data

	return
}

// For checking if the buffer is empty
func (cb *IntRingBuffer) IsEmpty() bool {
	return cb.size == 0
}

// For retrieving the size of the buffer
func (cb *IntRingBuffer) Count() int {
	return cb.size
}

// retrieving all
func (cb *IntRingBuffer) GetAll() (res []int) {
	//default result
	res = nil
	var (
		err  error
		item int
	)

	if cb.IsEmpty() {
		return
	}

	for {
		item, err = cb.Get()
		if err != nil {
			break
		}
		res = append(res, item)
	}

	return
}

// Removing any references present
// when the buffer becomes empty
func (cb *IntRingBuffer) innerClear() {
	cb.head = nil
	cb.tail = nil
}

// Clears buffer
func (cb *IntRingBuffer) Clear() {
	cb.size = 0
	cb.innerClear()
}
