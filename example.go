//go:build run

package main

import (
	"fmt"

	"github.com/hymkor/go-generics-list"
)

func main() {
	L := list.New[rune]()

	for _, c := range "ABCDEFG" {
		L.PushBack(c)
	}

	fmt.Print("Ascend:")
	for p := L.Front(); p != nil; p = p.Next() {
		fmt.Printf(" '%c'", p.Value)
	}
	fmt.Println()

	fmt.Print("Descend:")
	for p := L.Back(); p != nil; p = p.Prev() {
		fmt.Printf(" '%c'", p.Value)
	}
	fmt.Println()
}
