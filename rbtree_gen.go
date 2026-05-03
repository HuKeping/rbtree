// Copyright 2015, Hu Keping. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rbtree implements operations on Red-Black tree with Go 1.18+ generics.
package rbtree

// Lessable is the constraint for types that can be used in GenericRbtree.
// Types must implement Less(T) bool method for comparison.
type Lessable[T any] interface {
	Less(T) bool
}

// GenericRbtree is a Red-Black tree implementation using generics (Go 1.18+).
// It provides the same functionality as Rbtree but with type safety.
type GenericRbtree[T Lessable[T]] struct {
	nilNode *genericNode[T]
	root    *genericNode[T]
	count   uint
}

// genericNode represents a node in the generic red-black tree.
type genericNode[T Lessable[T]] struct {
	left   *genericNode[T]
	right  *genericNode[T]
	parent *genericNode[T]
	color  uint
	value  T
}

// NewGeneric creates a new generic red-black tree.
func NewGeneric[T Lessable[T]]() *GenericRbtree[T] {
	return new(GenericRbtree[T]).Init()
}

// Init initializes the generic red-black tree.
func (t *GenericRbtree[T]) Init() *GenericRbtree[T] {
	nilNode := &genericNode[T]{nil, nil, nil, BLACK, *new(T)}
	return &GenericRbtree[T]{
		nilNode: nilNode,
		root:    nilNode,
		count:   0,
	}
}

// Len returns the number of elements in the tree.
func (t *GenericRbtree[T]) Len() uint {
	return t.count
}

// Insert adds a value to the tree. If the value already exists, it is not inserted.
func (t *GenericRbtree[T]) Insert(value T) {
	t.insert(&genericNode[T]{t.nilNode, t.nilNode, t.nilNode, RED, value})
}

// InsertOrGet inserts the value if not present, otherwise returns the existing value.
// Returns the value in the tree (either the newly inserted one or the existing one).
func (t *GenericRbtree[T]) InsertOrGet(value T) T {
	return t.insert(&genericNode[T]{t.nilNode, t.nilNode, t.nilNode, RED, value}).value
}

// Get retrieves the value associated with the given key.
// Returns the value and true if found, zero value and false otherwise.
func (t *GenericRbtree[T]) Get(key T) (T, bool) {
	var zero T
	node := t.search(key)
	if node == t.nilNode {
		return zero, false
	}
	return node.value, true
}

// Contains reports whether the key exists in the tree.
func (t *GenericRbtree[T]) Contains(key T) bool {
	return t.search(key) != t.nilNode
}

// Delete removes the value from the tree.
// Returns the deleted value and true if it was present, zero value and false otherwise.
func (t *GenericRbtree[T]) Delete(key T) (T, bool) {
	var zero T
	node := t.delete(key)
	if node == t.nilNode {
		return zero, false
	}
	return node.value, true
}

// Min returns the minimum value in the tree.
// Returns the zero value if the tree is empty.
func (t *GenericRbtree[T]) Min() (T, bool) {
	var zero T
	if t.root == t.nilNode {
		return zero, false
	}
	node := t.min(t.root)
	if node == t.nilNode {
		return zero, false
	}
	return node.value, true
}

// Max returns the maximum value in the tree.
// Returns the zero value if the tree is empty.
func (t *GenericRbtree[T]) Max() (T, bool) {
	var zero T
	if t.root == t.nilNode {
		return zero, false
	}
	node := t.max(t.root)
	if node == t.nilNode {
		return zero, false
	}
	return node.value, true
}

// Clear removes all elements from the tree.
// It clears all node references to allow immediate GC.
func (t *GenericRbtree[T]) Clear() {
	t.clearTree(t.root)
	t.root = t.nilNode
	t.count = 0
}

// clearTree recursively clears all node references to help GC
func (t *GenericRbtree[T]) clearTree(x *genericNode[T]) {
	if x == t.nilNode {
		return
	}
	t.clearTree(x.left)
	t.clearTree(x.right)
	// Clear references to help GC
	x.left = nil
	x.right = nil
	x.parent = nil
}

// --- Internal methods ---

func (t *GenericRbtree[T]) leftRotate(x *genericNode[T]) {
	if x.right == t.nilNode {
		return
	}

	y := x.right
	x.right = y.left
	if y.left != t.nilNode {
		y.left.parent = x
	}
	y.parent = x.parent

	if x.parent == t.nilNode {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (t *GenericRbtree[T]) rightRotate(x *genericNode[T]) {
	if x.left == t.nilNode {
		return
	}

	y := x.left
	x.left = y.right
	if y.right != t.nilNode {
		y.right.parent = x
	}
	y.parent = x.parent

	if x.parent == t.nilNode {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.right = x
	x.parent = y
}

func (t *GenericRbtree[T]) insert(z *genericNode[T]) *genericNode[T] {
	x := t.root
	y := t.nilNode

	for x != t.nilNode {
		y = x
		if z.value.Less(x.value) {
			x = x.left
		} else if x.value.Less(z.value) {
			x = x.right
		} else {
			return x
		}
	}

	z.parent = y
	if y == t.nilNode {
		t.root = z
	} else if z.value.Less(y.value) {
		y.left = z
	} else {
		y.right = z
	}

	t.count++
	t.insertFixup(z)
	return z
}

func (t *GenericRbtree[T]) insertFixup(z *genericNode[T]) {
	for z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rightRotate(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *GenericRbtree[T]) min(x *genericNode[T]) *genericNode[T] {
	if x == t.nilNode {
		return t.nilNode
	}

	for x.left != t.nilNode {
		x = x.left
	}

	return x
}

func (t *GenericRbtree[T]) max(x *genericNode[T]) *genericNode[T] {
	if x == t.nilNode {
		return t.nilNode
	}

	for x.right != t.nilNode {
		x = x.right
	}

	return x
}

func (t *GenericRbtree[T]) search(key T) *genericNode[T] {
	p := t.root

	for p != t.nilNode {
		if p.value.Less(key) {
			p = p.right
		} else if key.Less(p.value) {
			p = p.left
		} else {
			break
		}
	}

	return p
}

func (t *GenericRbtree[T]) successor(x *genericNode[T]) *genericNode[T] {
	if x == t.nilNode {
		return t.nilNode
	}

	if x.right != t.nilNode {
		return t.min(x.right)
	}

	y := x.parent
	for y != t.nilNode && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

func (t *GenericRbtree[T]) delete(key T) *genericNode[T] {
	z := t.search(key)

	if z == t.nilNode {
		return t.nilNode
	}

	var y *genericNode[T]
	var x *genericNode[T]

	if z.left == t.nilNode || z.right == t.nilNode {
		y = z
	} else {
		y = t.successor(z)
	}

	if y.left != t.nilNode {
		x = y.left
	} else {
		x = y.right
	}

	x.parent = y.parent

	if y.parent == t.nilNode {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != z {
		z.value = y.value
	}

	if y.color == BLACK {
		t.deleteFixup(x)
	}

	t.count--

	return z
}

func (t *GenericRbtree[T]) deleteFixup(x *genericNode[T]) {
	for x != t.root && x.color == BLACK {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.right.color == BLACK {
					w.left.color = BLACK
					w.color = RED
					t.rightRotate(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				w.right.color = BLACK
				t.leftRotate(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					t.leftRotate(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				w.left.color = BLACK
				t.rightRotate(x.parent)
				x = t.root
			}
		}
	}
	x.color = BLACK
}

// AscendRange iterates over elements in the range [ge, lt) in ascending order.
// The iteration stops when iterator returns false.
func (t *GenericRbtree[T]) AscendRange(ge, lt T, iterator func(T) bool) {
	t.ascendRange(t.root, ge, lt, iterator)
}

// Ascend iterates over all elements >= pivot in ascending order.
// The iteration stops when iterator returns false.
func (t *GenericRbtree[T]) Ascend(pivot T, iterator func(T) bool) {
	t.ascend(t.root, pivot, iterator)
}

// Descend iterates over all elements <= pivot in descending order.
// The iteration stops when iterator returns false.
func (t *GenericRbtree[T]) Descend(pivot T, iterator func(T) bool) {
	t.descend(t.root, pivot, iterator)
}

// ForEach iterates over all elements in ascending order.
// The iteration stops when iterator returns false.
func (t *GenericRbtree[T]) ForEach(iterator func(T) bool) {
	t.forEach(t.root, iterator)
}

func (t *GenericRbtree[T]) forEach(x *genericNode[T], iterator func(T) bool) bool {
	if x == t.nilNode {
		return true
	}
	if !t.forEach(x.left, iterator) {
		return false
	}
	if !iterator(x.value) {
		return false
	}
	return t.forEach(x.right, iterator)
}

func (t *GenericRbtree[T]) ascend(x *genericNode[T], pivot T, iterator func(T) bool) bool {
	if x == t.nilNode {
		return true
	}

	if !x.value.Less(pivot) {
		if !t.ascend(x.left, pivot, iterator) {
			return false
		}
		if !iterator(x.value) {
			return false
		}
	}

	return t.ascend(x.right, pivot, iterator)
}

func (t *GenericRbtree[T]) descend(x *genericNode[T], pivot T, iterator func(T) bool) bool {
	if x == t.nilNode {
		return true
	}

	if !pivot.Less(x.value) {
		if !t.descend(x.right, pivot, iterator) {
			return false
		}
		if !iterator(x.value) {
			return false
		}
	}

	return t.descend(x.left, pivot, iterator)
}

func (t *GenericRbtree[T]) ascendRange(x *genericNode[T], ge, lt T, iterator func(T) bool) bool {
	if x == t.nilNode {
		return true
	}

	if !x.value.Less(lt) {
		return t.ascendRange(x.left, ge, lt, iterator)
	}
	if x.value.Less(ge) {
		return t.ascendRange(x.right, ge, lt, iterator)
	}

	if !t.ascendRange(x.left, ge, lt, iterator) {
		return false
	}
	if !iterator(x.value) {
		return false
	}
	return t.ascendRange(x.right, ge, lt, iterator)
}
