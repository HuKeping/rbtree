// Copyright 2015, Hu Keping. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

import (
	"reflect"
	"testing"
)

// IntWithLess implements the Less method for int values.
type IntWithLess int

func (x IntWithLess) Less(other IntWithLess) bool {
	return int(x) < int(other)
}

// StringWithLess implements the Less method for string values.
type StringWithLess string

func (x StringWithLess) Less(other StringWithLess) bool {
	return string(x) < string(other)
}

func TestGenericInsertAndDelete(t *testing.T) {
	rbt := NewGeneric[IntWithLess]()

	n := 1000
	for i := 0; i < n; i++ {
		rbt.Insert(IntWithLess(i))
	}
	if rbt.Len() != uint(n) {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), n)
	}

	for i := n - 1; i >= 0; i-- {
		val, ok := rbt.Delete(IntWithLess(i))
		if !ok {
			t.Errorf("Delete(%d) failed", i)
		}
		if val != IntWithLess(i) {
			t.Errorf("Delete(%d) returned %d", i, val)
		}
	}
	if rbt.Len() != 0 {
		t.Errorf("tree.Len() = %d, expect 0", rbt.Len())
	}
}

func TestGenericInsertOrGet(t *testing.T) {
	rbt := NewGeneric[IntWithLess]()

	rbt.Insert(IntWithLess(1))
	rbt.Insert(IntWithLess(2))
	rbt.Insert(IntWithLess(3))

	// Insert existing value
	existing := rbt.InsertOrGet(IntWithLess(1))
	if existing != IntWithLess(1) {
		t.Errorf("InsertOrGet(1) = %d, expect 1", existing)
	}

	// Insert new value
	newVal := rbt.InsertOrGet(IntWithLess(4))
	if newVal != IntWithLess(4) {
		t.Errorf("InsertOrGet(4) = %d, expect 4", newVal)
	}

	if rbt.Len() != 4 {
		t.Errorf("tree.Len() = %d, expect 4", rbt.Len())
	}
}

func TestGenericGet(t *testing.T) {
	rbt := NewGeneric[IntWithLess]()

	rbt.Insert(IntWithLess(1))
	rbt.Insert(IntWithLess(2))
	rbt.Insert(IntWithLess(3))

	// Get existing
	val, ok := rbt.Get(IntWithLess(1))
	if !ok {
		t.Errorf("Get(1) should exist")
	}
	if val != IntWithLess(1) {
		t.Errorf("Get(1) = %d, expect 1", val)
	}

	// Get non-existing
	_, ok = rbt.Get(IntWithLess(100))
	if ok {
		t.Errorf("Get(100) should not exist")
	}
}

func TestGenericContains(t *testing.T) {
	rbt := NewGeneric[IntWithLess]()

	rbt.Insert(IntWithLess(1))
	rbt.Insert(IntWithLess(2))

	if !rbt.Contains(IntWithLess(1)) {
		t.Errorf("Contains(1) should be true")
	}
	if rbt.Contains(IntWithLess(100)) {
		t.Errorf("Contains(100) should be false")
	}
}

func TestGenericDelete(t *testing.T) {
	rbt := NewGeneric[IntWithLess]()

	rbt.Insert(IntWithLess(1))
	rbt.Insert(IntWithLess(2))

	// Delete existing
	val, ok := rbt.Delete(IntWithLess(1))
	if !ok {
		t.Errorf("Delete(1) should succeed")
	}
	if val != IntWithLess(1) {
		t.Errorf("Delete(1) returned %d, expect 1", val)
	}

	// Delete non-existing
	_, ok = rbt.Delete(IntWithLess(100))
	if ok {
		t.Errorf("Delete(100) should fail")
	}
}

func TestGenericMin(t *testing.T) {
	rbt := NewGeneric[IntWithLess]()

	rbt.Insert(IntWithLess(3))
	rbt.Insert(IntWithLess(1))
	rbt.Insert(IntWithLess(2))

	val, ok := rbt.Min()
	if !ok {
		t.Errorf("Min() should exist")
	}
	if val != IntWithLess(1) {
		t.Errorf("Min() = %d, expect 1", val)
	}

	// Empty tree
	empty := NewGeneric[IntWithLess]()
	_, ok = empty.Min()
	if ok {
		t.Errorf("Min() on empty tree should return false")
	}
}

func TestGenericMax(t *testing.T) {
	rbt := NewGeneric[StringWithLess]()

	rbt.Insert(StringWithLess("a"))
	rbt.Insert(StringWithLess("h"))
	rbt.Insert(StringWithLess("z"))

	val, ok := rbt.Max()
	if !ok {
		t.Errorf("Max() should exist")
	}
	if val != StringWithLess("z") {
		t.Errorf("Max() = %v, expect z", val)
	}
}

func TestGenericClear(t *testing.T) {
	rbt := NewGeneric[IntWithLess]()

	rbt.Insert(IntWithLess(1))
	rbt.Insert(IntWithLess(2))
	rbt.Insert(IntWithLess(3))

	if rbt.Len() != 3 {
		t.Errorf("Len() = %d, expect 3", rbt.Len())
	}

	rbt.Clear()

	if rbt.Len() != 0 {
		t.Errorf("Len() after Clear() = %d, expect 0", rbt.Len())
	}

	_, ok := rbt.Min()
	if ok {
		t.Errorf("Min() after Clear() should return false")
	}
}

func TestGenericAscend(t *testing.T) {
	rbt := NewGeneric[StringWithLess]()

	rbt.Insert(StringWithLess("a"))
	rbt.Insert(StringWithLess("b"))
	rbt.Insert(StringWithLess("c"))
	rbt.Insert(StringWithLess("d"))

	rbt.Delete(StringWithLess("a"))

	var ret []StringWithLess
	// Ascend from "b" should return b, c, d
	rbt.Ascend(StringWithLess("b"), func(s StringWithLess) bool {
		ret = append(ret, s)
		return true
	})

	expected := []StringWithLess{"b", "c", "d"}
	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("expected %v but got %v", expected, ret)
	}
}

func TestGenericDescend(t *testing.T) {
	rbt := NewGeneric[IntWithLess]()

	for i := IntWithLess(0); i < 10; i++ {
		rbt.Insert(i)
	}

	var ret []IntWithLess
	// Descend from 5 should return elements <= 5 in descending order
	rbt.Descend(IntWithLess(5), func(i IntWithLess) bool {
		ret = append(ret, i)
		return true
	})

	expected := []IntWithLess{5, 4, 3, 2, 1, 0}
	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("expected %v but got %v", expected, ret)
	}
}

func TestGenericAscendRange(t *testing.T) {
	rbt := NewGeneric[StringWithLess]()

	strings := []StringWithLess{"a", "b", "c", "aa", "ab", "ac", "abc", "acb", "bac"}
	for _, v := range strings {
		rbt.Insert(v)
	}

	var ret []StringWithLess
	rbt.AscendRange(StringWithLess("ab"), StringWithLess("b"), func(s StringWithLess) bool {
		ret = append(ret, s)
		return true
	})

	expected := []StringWithLess{"ab", "abc", "ac", "acb"}
	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("expected %v but got %v", expected, ret)
	}
}

func TestGenericForEach(t *testing.T) {
	rbt := NewGeneric[IntWithLess]()

	rbt.Insert(IntWithLess(3))
	rbt.Insert(IntWithLess(1))
	rbt.Insert(IntWithLess(2))

	var ret []IntWithLess
	rbt.ForEach(func(i IntWithLess) bool {
		ret = append(ret, i)
		return true
	})

	expected := []IntWithLess{1, 2, 3}
	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("expected %v but got %v", expected, ret)
	}
}

func TestGenericInsertDup(t *testing.T) {
	rbt := NewGeneric[StringWithLess]()

	rbt.Insert(StringWithLess("go"))
	rbt.Insert(StringWithLess("go"))
	rbt.Insert(StringWithLess("go"))

	if rbt.Len() != 1 {
		t.Errorf("tree.Len() = %d, expect 1", rbt.Len())
	}
}
