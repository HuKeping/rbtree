// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

import (
	"reflect"
	"testing"
)

func TestInsertAndDelete(t *testing.T) {
	rbt := New()

	m := 0
	n := 1000
	for m < n {
		rbt.Insert(Int(m))
		m++
	}
	if rbt.Len() != uint(n) {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), n)
	}

	for m > 0 {
		rbt.Delete(Int(m))
		m--
	}
	if rbt.Len() != 1 {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), 1)
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

func TestDescend(t *testing.T) {
	rbt := New()

	m := 0
	n := 10
	for m < n {
		rbt.Insert(Int(m))
		m++
	}

	var ret []Item

	rbt.Descend(Int(1), func(i Item) bool {
		ret = append(ret, i)
		return true
	})
	expected := []Item{Int(1), Int(0)}
	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("expected %v but got %v", expected, ret)
	}

	ret = nil
	rbt.Descend(Int(10), func(i Item) bool {
		ret = append(ret, i)
		return true
	})
	expected = []Item{Int(9), Int(8), Int(7), Int(6), Int(5), Int(4), Int(3), Int(2), Int(1), Int(0)}
	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("expected %v but got %v", expected, ret)
	}
}

func TestGet(t *testing.T) {
	rbt := New()

	rbt.Insert(Int(1))
	rbt.Insert(Int(2))
	rbt.Insert(Int(3))

	no := rbt.Get(Int(100))
	ok := rbt.Get(Int(1))

	if no != nil {
		t.Errorf("100 is expect not exists")
	}

	if ok == nil {
		t.Errorf("1 is expect exists")
	}
}

func TestAscend(t *testing.T) {
	rbt := New()

	rbt.Insert(String("a"))
	rbt.Insert(String("b"))
	rbt.Insert(String("c"))
	rbt.Insert(String("d"))

	rbt.Delete(rbt.Min())

	var ret []Item
	rbt.Ascend(rbt.Min(), func(i Item) bool {
		ret = append(ret, i)
		return true
	})

	expected := []Item{String("b"), String("c"), String("d")}
	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("expected %v but got %v", expected, ret)
	}
}
