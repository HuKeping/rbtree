// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

type Iterator func(i Item) bool

// Ascend will call iterator once for each element greater than pivot
// in ascending order. It will stop whenever the iterator returns false.
func (t *Rbtree) Ascend(pivot Item, iterator Iterator) {
	t.ascend(t.root, pivot, iterator)
}

func (t *Rbtree) ascend(h *Node, pivot Item, iterator Iterator) bool {
	if h == t.NIL {
		return true
	}

	if !less(h.Item, pivot) {
		if !t.ascend(h.Left, pivot, iterator) {
			return false
		}
		if !iterator(h.Item) {
			return false
		}
	}

	return t.ascend(h.Right, pivot, iterator)
}

func (t *Rbtree) Min() Item {
	x := t.min(t.root)

	if x == t.NIL {
		return nil
	}

	return x.Item
}
