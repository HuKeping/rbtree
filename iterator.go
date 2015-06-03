// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

type Iterator func(i Item) bool

// Ascend will call iterator once for each element greater or equal than pivot
// in ascending order. It will stop whenever the iterator returns false.
func (t *Rbtree) Ascend(pivot Item, iterator Iterator) {
	t.ascend(t.root, pivot, iterator)
}

func (t *Rbtree) ascend(x *Node, pivot Item, iterator Iterator) bool {
	if x == t.NIL {
		return true
	}

	if !less(x.Item, pivot) {
		if !t.ascend(x.Left, pivot, iterator) {
			return false
		}
		if !iterator(x.Item) {
			return false
		}
	}

	return t.ascend(x.Right, pivot, iterator)
}

// Descend will call iterator once for each element less or equal than pivot
// in descending order. It will stop whenever the iterator returns false.
func (t *Rbtree) Descend(pivot Item, iterator Iterator) {
	t.descend(t.root, pivot, iterator)
}

func (t *Rbtree) descend(x *Node, pivot Item, iterator Iterator) bool {
	if x == nil {
		return true
	}

	if !less(pivot, x.Item) {
		if !t.descend(x.Right, pivot, iterator) {
			return false
		}
		if !iterator(x.Item) {
			return false
		}
	}

	return t.descend(x.Left, pivot, iterator)
}

func (t *Rbtree) Min() Item {
	x := t.min(t.root)

	if x == t.NIL {
		return nil
	}

	return x.Item
}
