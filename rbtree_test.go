// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

import (
	"testing"
)

func TestInsertInt(t *testing.T) {
	rbt := New()

	rbt.Insert(Int(1))
	rbt.Insert(Int(2))
	rbt.Insert(Int(3))
	rbt.Insert(Int(4))
	rbt.Insert(Int(5))

	if rbt.Len() != 5 {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), 5)
	}
}

func TestInsertString(t *testing.T) {
	rbt := New()

	rbt.Insert(String("go"))
	rbt.Insert(String("lang"))

	if rbt.Len() != 2 {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), 2)
	}
}

// Test for duplicate
func TestInsertDup(t *testing.T) {
	rbt := New()

	rbt.Insert(String("go"))
	rbt.Insert(String("go"))
	rbt.Insert(String("go"))

	if rbt.Len() != 1 {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), 1)
	}
}

func TestDelete(t *testing.T) {
	rbt := New()

	rbt.Insert(Int(1))
	rbt.Insert(Int(2))
	rbt.Insert(Int(3))

	rbt.Delete(Int(1))
	rbt.Delete(Int(2))
	rbt.Delete(Int(3))
	if rbt.Len() != 0 {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), 3)
	}
}
