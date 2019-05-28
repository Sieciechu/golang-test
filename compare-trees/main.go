package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkRecursive(t, ch)
	close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	if nil != t.Left {
		WalkRecursive(t.Left, ch)
	}

	ch <- t.Value

	if nil != t.Right {
		WalkRecursive(t.Right, ch)
	}
}

func printValues(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	areSame := false


	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 && !ok2 {
			break;
		}

		fmt.Printf("%d == %d ->%v\n", v1, v2, v1==v2)
		if v1 != v2 {
			areSame = false
			break
		}

		areSame = true
	}


	return areSame

}

func main() {
	t1:= tree.New(1)
	t2:= tree.New(2)

	areSame := Same(t1, t2)
	fmt.Printf("Same(t1, t2) ->%v\n", areSame)
	fmt.Printf("t1 == t2 ->%v\n", t1==t2)

}
