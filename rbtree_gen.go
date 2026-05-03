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
