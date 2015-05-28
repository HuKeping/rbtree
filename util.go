// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

type Int int

func (x Int) Less(than Item) bool {
	return x < than.(Int)
}

type String string

func (x String) Less(than Item) bool {
	return x < than.(String)
}
