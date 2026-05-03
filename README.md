# Rbtree  [![GoDoc](https://godoc.org/github.com/HuKeping/rbtree?status.svg)](https://godoc.org/github.com/HuKeping/rbtree)

This is an implementation of Red-Black tree written by Golang which does **not** support `duplicate keys`.

## Installation

With a healthy Go language installed, simply run `go get github.com/HuKeping/rbtree`

## Example
All you have to do is to implement a comparison function `Less() bool` for your Item
which will be store in the Red-Black tree, here are some examples.
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

#### A quite interesting case for `struct` items.
	package main
	
	import (
		"fmt"
		"github.com/HuKeping/rbtree"
		"time"
	)
	
	type Var struct {
		Expiry time.Time `json:"expiry,omitempty"`
		ID     string    `json:"id",omitempty`
	}
	
	// We will order the node by `Time`
	func (x Var) Less(than rbtree.Item) bool {
		return x.Expiry.Before(than.(Var).Expiry)
	}
	
	func main() {
		rbt := rbtree.New()
	
		var1 := Var{
			Expiry: time.Now().Add(time.Second * 10),
			ID:     "var1",
		}
		var2 := Var{
			Expiry: time.Now().Add(time.Second * 20),
			ID:     "var2",
		}
		var3 := Var{
			Expiry: var2.Expiry,
			ID:     "var2-dup",
		}
		var4 := Var{
			Expiry: time.Now().Add(time.Second * 40),
			ID:     "var4",
		}
		var5 := Var{
			Expiry: time.Now().Add(time.Second * 50),
			ID:     "var5",
		}
	
		rbt.Insert(var1)
		rbt.Insert(var2)
		rbt.Insert(var3)
		rbt.Insert(var4)
		rbt.Insert(var5)
	
		tmp := Var{
			Expiry: var4.Expiry,
			ID:     "This field is not the key factor",
		}
	
		// var4 and var5 were expected
		rbt.Ascend(rbt.Get(tmp), Print)
	}
	
	func Print(item rbtree.Item) bool {
		i, ok := item.(Var)
		if !ok {
			return false
		}
		fmt.Printf("%+v\n", i)
		return true
	}

## Type-Safe Generic API (Go 1.18+)

For Go 1.18 and later, a generic implementation is available with better type safety and performance. The original `Rbtree` API will continue to be maintained.

```go
package main

import (
	"fmt"
	"github.com/HuKeping/rbtree"
)

type MyID int64

func (x MyID) Less(y MyID) bool {
	return x < y
}

func main() {
	tree := rbtree.NewGeneric[MyID]()

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	// Get returns (value, found)
	if val, ok := tree.Get(2); ok {
		fmt.Println(val)
	}

	// Contains for existence check
	fmt.Println(tree.Contains(2)) // true

	// ForEach for full traversal
	tree.ForEach(func(id MyID) bool {
		fmt.Println(id)
		return true
	})

	// Clear all elements
	tree.Clear()
}
```

### Key Differences from Legacy API

| Operation | Legacy (`Rbtree`) | Generic (`GenericRbtree[T]`) |
|-----------|-------------------|------------------------------|
| Create | `rbtree.New()` | `rbtree.NewGeneric[T]()` |
| Get | `Get(key) Item` | `Get(key) (T, bool)` |
| Delete | `Delete(key) Item` | `Delete(key) (T, bool)` |
| Min/Max | `Min() Item` | `Min() (T, bool)` |
| Contains | N/A | `Contains(key) bool` |
| Clear | N/A | `Clear()` |

The generic version provides:
- **Compile-time type safety** - no runtime type assertions
- **Better performance** - 14-31% faster in benchmarks
- **Zero allocations** for `Get` operation
- **Explicit success indication** via `(value, bool)` return
