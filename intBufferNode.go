package main


// Buffer using a Linked List 
// A Generic Node class is used to create a Linked List
type IntBufferNode struct {
    // Data Stored in each Node of the Linked List
    data int
    // Pointer to the next node in the Linked List
    next *IntBufferNode    
}

// Node class constructor used to initializes
    // the data in each Node
func NewIntBufferNode(data int) *IntBufferNode{
    return &IntBufferNode{data: data}
}

func (bn *IntBufferNode) Next(next *IntBufferNode){
    bn.next= next
}

func (bn *IntBufferNode) Clear(){
    bn.next=nil
}