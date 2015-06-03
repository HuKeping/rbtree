# Rbtree

This is an implementation of Red-Black tree written by Golang.

## Installation

With a healthy Go language installed, simply run `go get github.com/HuKeping/rbtree`

## Example

#### A simple case for `int` items.
	package main
	
	import (
		"fmt"
		"github.com/HuKeping/rbtree"
	)
	
	func main() {
		rbt := rbtree.New()
	
		m := 0
		n := 10
	
		for m < n {
			rbt.Insert(rbtree.Int(m))
			m++
		}
	
		m = 0
		for m < n {
			if m%2 == 0 {
				rbt.Delete(rbtree.Int(m))
			}
			m++
		}

		// 1, 3, 5, 7, 9 were expected.
		rbt.Ascend(rbt.Min(), Print)
	}
	
	func Print(item rbtree.Item) bool {
		i, ok := item.(rbtree.Int)
		if !ok {
			return false
		}
		fmt.Println(i)
		return true
	}

#### A simple case for `string` items.
	package main
	
	import (
		"fmt"
		"github.com/HuKeping/rbtree"
	)
	
	func main() {
		rbt := rbtree.New()
	
		rbt.Insert(rbtree.String("Hello"))
		rbt.Insert(rbtree.String("World"))

		rbt.Ascend(rbt.Min(), Print)
	}
	
	func Print(item rbtree.Item) bool {
		i, ok := item.(rbtree.String)
		if !ok {
			return false
		}
		fmt.Println(i)
		return true
	}
