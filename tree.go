package main

import (
	"golang.org/x/tour/tree"
	"sort"
	"fmt"
)
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walker(t *tree.Tree) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := Walker(t1), Walker(t2)
	slice1 := make([]int, cap(c1))
	slice2 := make([]int, cap(c2))
	for val1 := range c1 {
		slice1 = append(slice1, val1)
	}
	for val2 := range c2 {
		slice2 = append(slice2, val2)
	}
	fmt.Println(slice1)
	fmt.Println(slice2)
	sort.Sort(sort.IntSlice(slice1))
	sort.Sort(sort.IntSlice(slice2))
	min1 := slice1[0]
	min2 := slice2[0]
	max1 := slice1[len(slice1)-1]
	max2 := slice2[len(slice2)-1]
	success := false
	if min1 == min2 {
		if max1 == max2 {
			success = true
		}
	}
	return success
}

func main() {
	fmt.Println(Same(tree.New(3), tree.New(1)))
}