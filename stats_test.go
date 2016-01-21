// Copyright 2015, Hu Keping. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

import (
	"testing"
)

func TestDeleteReturnValue(t *testing.T) {
	rbt := New()

	rbt.Insert(String("go"))
	rbt.Insert(String("lang"))

	if rbt.Len() != 2 {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), 2)
	}

	// go should be in the rbtree
	deletedGo := rbt.Delete(String("go"))
	if deletedGo != String("go") {
		t.Errorf("expect %v, got %v", "go", deletedGo)
	}

	// C should not be in the rbtree
	deletedC := rbt.Delete(String("C"))
	if deletedC != nil {
		t.Errorf("expect %v, got %v", nil, deletedC)
	}
}
