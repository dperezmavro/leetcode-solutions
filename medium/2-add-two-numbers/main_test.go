package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	for n := addTwoNumbers(sToN([]int{2, 4, 3}), sToN([]int{5, 6, 4})); n.Next != nil; n = n.Next {
		log.Println(n.Val)
	}
}
