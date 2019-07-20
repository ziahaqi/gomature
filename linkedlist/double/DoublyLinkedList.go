package main

import (
	"fmt"
)

type DoublyLinkedList struct {
	head *Node
	tail *Node
	Size int
}
type Node struct {
	next   *Node
	prev   *Node
	Person Person
}

type Person struct {
	Name string
	Age  int
}

func main() {
	list := DoublyLinkedList{
		head: nil,
		Size: 0,
	}

	p := Person{
		Name: "A",
		Age:  18,
	}
	q := Person{
		Name: "B",
		Age:  19,
	}
	r := Person{
		Name: "C",
		Age:  24,
	}
	s := Person{
		Name: "D",
		Age:  28,
	}
	addToFront(p, &list)
	addToFront(q, &list)
	addToEnd(r, &list)
	addToEnd(s, &list)
	PrintList(&list)
	fmt.Println("\nsize:", list.Size)

	fmt.Println("\n==remove front==")
	removeFromFront(&list)
	PrintList(&list)
	fmt.Println("\nsize:", list.Size)

	fmt.Println("\n==remove end==")
	removeFromEnd(&list)
	PrintList(&list)
	fmt.Println("\nsize:", list.Size)

	found := findByName("A", &list)
	if found != nil {
		fmt.Println("person found:", *found)
	} else {
		fmt.Println("person notfound:")
	}
}
func addToFront(person Person, list *DoublyLinkedList) {
	newNode := Node{
		Person: person,
	}
	newNode.next = list.head

	if list.head == nil {
		list.tail = &newNode
	} else {
		list.head.prev = &newNode
		newNode.next = list.head
	}

	list.head = &newNode
	list.Size++
}

func addToEnd(person Person, list *DoublyLinkedList) {
	newNode := Node{
		Person: person,
	}

	if list.tail == nil {
		list.head = &newNode
	} else {
		list.tail.next = &newNode
		newNode.prev = list.tail
	}

	list.tail = &newNode
	list.Size++
}

func removeFromFront(list *DoublyLinkedList) {
	if list.head == nil {
		return
	}

	current := list.head
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.next.prev = nil
	}

	list.head = list.head.next
	current.next = nil

	list.Size--
}

func removeFromEnd(list *DoublyLinkedList) {
	current := list.tail

	if list.tail.prev == nil {
		list.head = nil
	} else {
		list.tail.prev.next = nil
	}
	list.tail = list.tail.prev
	current.prev = nil
	list.Size--
}

func findByName(name string, list *DoublyLinkedList) *Person {
	current := list.head
	if current == nil {
		return nil
	}

	if current.Person.Name == name {
		return &current.Person
	}
	for current.next != nil {
		current = current.next
		if current.Person.Name == name {
			return &current.Person
		}
	}

	return nil
}

func PrintList(list *DoublyLinkedList) {
	current := list.head

	if current == nil {
		fmt.Println("emtpy")
		return
	}

	fmt.Print("node:", current.Person, "->")
	for current.next != nil {
		current = current.next
		fmt.Print("node:", current.Person, "<->")
	}
	fmt.Print("nil")
}
