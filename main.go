package main

import (
	"fmt"
	"strings"

	"github.com/01-edu/z01"
)

func printstring(str string) {
	for i := 0; i < len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}
}

func printslice(str string) {
	list := []string{}
	for _, v := range str {
		list = append(list, string(v))
	}
	fmt.Println(list)
}

// functions

func add(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Error division by 0"
	}
	return a / b, "nil"
}

// Pointers & struct

// Struct

// pointer

// Linked List
type Node struct {
	Val  int
	Next *Node
}

// Trees and Graphs
type TreeNode struct {
	val         int
	Left, Right *TreeNode
}

func main() {

	// linkedlist test
	n1 := &Node{Val: 1, Next: nil}
	n2 := &Node{Val: 2, Next: nil}
	n1.Next = n2

	// Tree testing
	root := &TreeNode{10, &TreeNode{5, nil, nil}, &TreeNode{15, nil, nil}}
	fmt.Println(root.Left.val)
	fmt.Println(root.Right.val)

	// Graph( Adjacency List )
	// graph := map[string][]string{
	//	"A": {"B", "C"},
	//	"B": {"A", "D"},
	//	"C": {"A", "D"},
	// }

	// numbers strings, Runes, booleans
	var n int = 9
	var f float64 = 3.14
	var b bool = true
	var s string = "Hello Ade"

	fmt.Print(n, f, b, s)

	// runes and unicode
	var r rune = 'A'
	fmt.Println(r)

	fmt.Printf("%c\n", r)

	str := "Go 🐱"
	for i, ch := range str {
		fmt.Println(i, ch, string(ch))
	}

	strr := "Go is awesome"

	fmt.Println(strings.ToUpper(strr))

	// Arrays, Slices, Maps
	arr := [3]int{} // array are list with fixed size
	fmt.Println(arr)

	slic := []int{1, 2, 3} // slice has variable length
	slic = append(slic, 4)
	fmt.Println(slic)

	// Map

	scores := map[string]int{"Alice": 80, "Ade": 100}
	scores["Eve"] = 85

	for k, v := range scores {
		fmt.Println(k, v)
	}

	// control structures

	x := 10

	if x <= 9 {
		fmt.Println("small")
	} else {
		fmt.Println("big")
	}

	// loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	nums := []int{2, 4, 6, 8}

	for i, v := range nums {
		fmt.Println(i, v)
	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		if v%2 == 0 {
			fmt.Println(k, "is even")
		} else {
			fmt.Println(k, "is not even")
		}
	}

	store := []int{}
	store = append(store, nums[0])

	for i := 1; i < (len(nums)); i++ {
		store = append(store, nums[i]+store[i-1])
	}
	fmt.Println(store)

	// both pointers in same direction
	numbs := []int{1, 1, 2, 3, 3, 4, 4, 5, 5, 5, 5, 6, 6, 7, 7}
	slow := 0
	duplicates := []int{}

	for fast := 1; fast < len(numbs); fast++ {
		if numbs[fast] != numbs[slow] {
			slow++
			numbs[slow] = numbs[fast]
		} else if numbs[fast] == numbs[slow] {
			duplicates = append(duplicates, numbs[slow])
		}
	}
	fmt.Println(slow)
	fmt.Println(numbs[:slow+1])
	fmt.Println(duplicates)

}
