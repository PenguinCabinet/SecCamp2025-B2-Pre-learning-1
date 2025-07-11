package main

import (
	"fmt"

	"golang.org/x/tour/tree"
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

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	ch2 := make(chan int)
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for {
		e1, flag1 := <-ch1
		e2, flag2 := <-ch2
		if flag1 != flag2 {
			return false
		}
		if e1 != e2 {
			return false
		}
		if flag1 && flag2 {
			break
		}
	}
	return true
}

func main() {
	/*
		ch:=make(chan int)
		go func(){
			Walk(tree.New(1),ch)
			close(ch)
		}()
		for e:=range ch{
			fmt.Printf("%v\n",e)
		}
	*/
	fmt.Printf("tree.New(1), tree.New(1), %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("tree.New(1), tree.New(2), %v\n", Same(tree.New(1), tree.New(2)))
}
