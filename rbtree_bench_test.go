// Copyright 2015, Hu Keping. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

import (
	"testing"
)

// Benchmark for legacy Rbtree with Int type

func BenchmarkLegacyInsert(b *testing.B) {
	rbt := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Insert(Int(i))
	}
}

func BenchmarkLegacyGet(b *testing.B) {
	rbt := New()
	// Setup: insert 1000 elements
	for i := 0; i < 1000; i++ {
		rbt.Insert(Int(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Get(Int(i % 1000))
	}
}

func BenchmarkLegacyDelete(b *testing.B) {
	rbt := New()
	// Setup: insert 1000 elements
	for i := 0; i < 1000; i++ {
		rbt.Insert(Int(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Delete(Int(i % 1000))
		// Re-insert to keep tree populated
		rbt.Insert(Int(i % 1000))
	}
}

func BenchmarkLegacyMin(b *testing.B) {
	rbt := New()
	// Setup: insert 1000 elements
	for i := 0; i < 1000; i++ {
		rbt.Insert(Int(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Min()
	}
}

func BenchmarkLegacyAscend(b *testing.B) {
	rbt := New()
	// Setup: insert 1000 elements
	for i := 0; i < 1000; i++ {
		rbt.Insert(Int(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count := 0
		rbt.Ascend(Int(0), func(item Item) bool {
			count++
			return count < 100 // Only iterate first 100
		})
	}
}

// Benchmark for GenericRbtree with IntWithLess type

func BenchmarkGenericInsert(b *testing.B) {
	rbt := NewGeneric[IntWithLess]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Insert(IntWithLess(i))
	}
}

func BenchmarkGenericGet(b *testing.B) {
	rbt := NewGeneric[IntWithLess]()
	// Setup: insert 1000 elements
	for i := 0; i < 1000; i++ {
		rbt.Insert(IntWithLess(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Get(IntWithLess(i % 1000))
	}
}

func BenchmarkGenericDelete(b *testing.B) {
	rbt := NewGeneric[IntWithLess]()
	// Setup: insert 1000 elements
	for i := 0; i < 1000; i++ {
		rbt.Insert(IntWithLess(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Delete(IntWithLess(i % 1000))
		// Re-insert to keep tree populated
		rbt.Insert(IntWithLess(i % 1000))
	}
}

func BenchmarkGenericMin(b *testing.B) {
	rbt := NewGeneric[IntWithLess]()
	// Setup: insert 1000 elements
	for i := 0; i < 1000; i++ {
		rbt.Insert(IntWithLess(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Min()
	}
}

func BenchmarkGenericAscend(b *testing.B) {
	rbt := NewGeneric[IntWithLess]()
	// Setup: insert 1000 elements
	for i := 0; i < 1000; i++ {
		rbt.Insert(IntWithLess(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count := 0
		rbt.Ascend(IntWithLess(0), func(item IntWithLess) bool {
			count++
			return count < 100 // Only iterate first 100
		})
	}
}

// Mixed workload benchmark

func BenchmarkLegacyMixed(b *testing.B) {
	rbt := New()
	// Setup: insert 1000 elements
	for i := 0; i < 1000; i++ {
		rbt.Insert(Int(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Get(Int(i % 1000))
		rbt.Insert(Int(1000 + i))
		if i%10 == 0 {
			rbt.Delete(Int(i % 1000))
		}
	}
}

func BenchmarkGenericMixed(b *testing.B) {
	rbt := NewGeneric[IntWithLess]()
	// Setup: insert 1000 elements
	for i := 0; i < 1000; i++ {
		rbt.Insert(IntWithLess(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Get(IntWithLess(i % 1000))
		rbt.Insert(IntWithLess(1000 + i))
		if i%10 == 0 {
			rbt.Delete(IntWithLess(i % 1000))
		}
	}
}
