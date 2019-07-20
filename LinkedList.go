package gomature

import (
	"fmt"
)

type SinglyLinkedList struct {
	head *Node
	Size int
}
type Node struct {
	next   *Node
	Person Person
}

type Person struct {
	Name string
	Age int
}

func main() {
	list := SinglyLinkedList{
		head: nil,
		Size: 0,
	}

	p := Person{
		Name: "edi supono",
		Age: 18,
	}
	q := Person{
		Name: "cipluk",
		Age:19,
	}
	r := Person{
		Name: "Palui",
		Age: 24,
	}
	s := Person{
		Name: "Fulan",
		Age:28,
	}
	add(p, &list)
	add(q, &list)
	add(r, &list)
	add(s, &list)

	PrintList(&list)
	fmt.Println("\nsize:", list.Size)
	removeFromFront(&list)
	fmt.Println("\n==after removal==")
	PrintList(&list)
	fmt.Println("\nsize:", list.Size)

	 found := findByName("cipluk", &list)

	 if(found != nil){
		 fmt.Println("person found:", *found)
	 }else{
		 fmt.Println("person notfound:")
	 }

}
func add(person Person, list *SinglyLinkedList) {
	newNode := Node{
		Person: person,
	}
	newNode.next = list.head
	list.head = &newNode
	list.Size++
}

func removeFromFront(list * SinglyLinkedList){
	current := list.head
	list.head = current.next
	current.next = nil
	list.Size--
}

func findByName(name string, list *SinglyLinkedList) *Person{
	current := list.head
	if current == nil {
		return nil
	}

	if current.Person.Name == name{
		return &current.Person
	}
	for current.next != nil {
		current = current.next
		if current.Person.Name == name{
			return &current.Person
		}
	}

	return nil
}

func PrintList(list *SinglyLinkedList) error {
	current := list.head

	if current == nil {
		fmt.Println("emtpy")
		return nil
	}

	fmt.Print("node:", current.Person, "->")
	for current.next != nil {
		current = current.next
		fmt.Print("node:", current.Person, "->")
	}
	fmt.Print("nil")

	return nil
}
