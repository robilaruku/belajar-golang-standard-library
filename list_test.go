package belajar_golang_standard_library

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	var data *list.List = list.New()

	data.PushBack("Robi")
	data.PushBack("Laruku")
	data.PushBack("Dream")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Robi

	next := head.Next() // Laruku
	fmt.Println(next.Value)

	next = next.Next() // Dream
	fmt.Println(next.Value)

	fmt.Println()

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
