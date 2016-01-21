// Copyright 2015, Hu Keping . All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

// This file contains most of the methods that can be used
// by the user. Anyone who wants to look for some API about
// the rbtree, this is the right place.

// Number of nodes in the tree.
func (t *Rbtree) Len() uint { return t.count }

func (t *Rbtree) Insert(item Item) {
	if item == nil {
		return
	}

	// Always insert a RED node
	t.insert(&Node{t.NIL, t.NIL, t.NIL, RED, item})
}

//InsertOrGet inserts or retrieves the item in the tree. If the
//item is already in the tree then the return value will be that.
//If the item is not in the tree the return value will be the item
//you put in.
func (t *Rbtree) InsertOrGet(item Item) Item {
	if item == nil {
		return nil
	}

	return t.insert(&Node{t.NIL, t.NIL, t.NIL, RED, item}).Item
}

func (t *Rbtree) Delete(item Item) Item {
	if item == nil {
		return nil
	}

	// The `color` field here is nobody
	return t.delete(&Node{t.NIL, t.NIL, t.NIL, RED, item}).Item
}

func (t *Rbtree) Get(item Item) Item {
	if item == nil {
		return nil
	}

	// The `color` field here is nobody
	ret := t.search(&Node{t.NIL, t.NIL, t.NIL, RED, item})
	if ret == nil {
		return nil
	}

	return ret.Item
}

//TODO: This is for debug, delete it in the future
func (t *Rbtree) Search(item Item) *Node {

	return t.search(&Node{t.NIL, t.NIL, t.NIL, RED, item})
}

func (t *Rbtree) Min() Item {
	x := t.min(t.root)

	if x == t.NIL {
		return nil
	}

	return x.Item
}

func (t *Rbtree) Max() Item {
	x := t.max(t.root)

	if x == t.NIL {
		return nil
	}

	return x.Item
}
