// Copyright 2015, Hu Keping. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/HuKeping/rbtree"
)

func main() {
	rbt := rbtree.New()

	m := 0
	n := 10

	for m < n {
		rbt.Insert(rbtree.Int(m))
		m++
	}

	m = 0
	for m < n {
		if m%2 == 0 {
			rbt.Delete(rbtree.Int(m))
		}
		m++
	}

	rbt.Ascend(rbt.Min(), Print)
}

func Print(item rbtree.Item) bool {
	i, ok := item.(rbtree.Int)
	if !ok {
		return false
	}
	fmt.Println(i)
	return true
}
