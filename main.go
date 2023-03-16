package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	val   int // store value
	count int // using count to store duplicate value
	next  *node
	prev  *node
}

// CircleList is a circle list
// will add a new element and random assign head
type CircleList struct {
	// this list is a circle list
	head   *node
	length int
}

func (l *CircleList) add(n *node) {

	// will add a new element
	// and assign random head
	if l.length == 0 {

		// length 0 => init
		l.head = n
		l.head.next = n
		l.head.prev = n

	} else if l.length == 1 {

		// length 1 => add to head
		l.head.next = n
		l.head.prev = n
		n.prev = l.head
		n.next = l.head

	} else {
		// add to head
		l.head.next.prev = n
		n.next = l.head.next

		l.head.next = n
		n.prev = l.head
	}

	// random assign l head
	if rand.Int()%2 == 0 {
		l.head = l.head.next
	} else {
		l.head = l.head.prev
	}

	l.length++
}

func (l *CircleList) remove(n *node) bool {
	// in case of nil
	if n == nil {
		return false
	}

	// in case of length 1, remove head
	if l.length == 1 {
		l.head = nil
		l.length = 0
		return true
	}

	// case removed node is head, change head
	if n == l.head {
		l.head = n.next
	}

	// remove node
	n.prev.next = n.next
	n.next.prev = n.prev

	n.next = nil
	n.prev = nil

	l.length--

	return true
}

// O1st is struct will do add, remove and remove random with complexity O(1)
type O1st struct {
	list        CircleList    // store value
	length      int           // count
	valueToNode map[int]*node // map value to node in case need quick access node without search
}

func (o *O1st) add(number int) {
	// in case of duplicate
	if _, ok := o.valueToNode[number]; ok {
		o.valueToNode[number].count++
		return
	}

	// new node
	n := &node{
		val:   number,
		count: 1,
	}

	// add to node
	o.valueToNode[number] = n
	o.list.add(n)
	o.length++
}

func (o *O1st) remove(number int) bool {
	// in case no value found
	// never appear or already removed
	if n, ok := o.valueToNode[number]; !ok || n == nil {
		return false
	}

	// remove count
	if o.valueToNode[number].count > 1 {
		o.valueToNode[number].count--
		return true
	}

	// remove node
	o.list.remove(o.valueToNode[number])
	o.valueToNode[number] = nil
	o.length--
	return true
}

func (o *O1st) removeRandom() int {
	// because new element is added to with random order
	// so remove random element is also random
	val := o.list.head.val
	o.remove(val)
	return val
}

func newO1() *O1st {
	return &O1st{
		list:        CircleList{},
		length:      0,
		valueToNode: make(map[int]*node),
	}
}

func main() {
	/* explain
		the struct O1st include:
		- list: CircleList
				- can store a duplicate value
				- can add a new element and random assign head
		- length: number of number add to the struct
		- valueToNode: map[int]*node
				- use to quick access node in case need to remove or update node

	complexity:
		add: O(1) : just append next to o1st.list.head and random assign head
		remove: O(1) : just remove node from list and update valueToNode
		removeRandom: O(1) : just remove o1st.list.head, because when add new element, head is random node, so remove head is also random
	*/

	o1 := newO1()
	// change any value you want to test below

	o1.add(1)
	o1.add(2)
	o1.add(3)
	o1.add(4)
	o1.add(5)
	o1.add(6)

	fmt.Println("remove random: ", o1.removeRandom())

	fmt.Println("remove 4:", o1.remove(4))
	fmt.Println("remove 6:", o1.remove(6))
}
