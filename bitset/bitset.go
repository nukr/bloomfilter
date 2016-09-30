// Copyright 2011 Will Fitzgerald. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package bitset implements bitsets.

	It provides methods for making a bitset of an arbitrary
	upper limit, setting and testing bit locations, and clearing
	bit locations as well as the entire set.

	bitsets are implmeented as arrays of uint64s, so it may be best
	to limit the upper size to a multiple of 64. It is not an error
	to set, test, or clear a bit within a 64 bit boundary.

	Example use:

	b := MakeBitSet(64000)
	b.SetBit(1000)
	if b.Bit(1000) {
		b.ClearBit(1000)
	}
*/
package bitset

// for MaxUint64
import (
	"fmt"
	"math"
)

// Bitset are arrays of uint64.
type Bitset []uint64

// MakeBitSet with an upper limit on size. Note this is the
// number of bits, not the number of uint64s, which is a kind of
// implementation detail.
func MakeBitSet(maxSize uint) Bitset {
	if maxSize%64 == 0 {
		s := make(Bitset, maxSize/64)
		s.Clear()
		return s
	}
	s := make(Bitset, maxSize/64+1)
	s.Clear()
	return s
}

// Bit Test whether bit i is set.
func (set Bitset) Bit(i uint) bool {
	return ((set[i/64] & (1 << (i % 64))) != 0)
}

// SetBit i to 1
func (set Bitset) SetBit(i uint) {
	set[i/64] |= (1 << (i % 64))
}

// ClearBit i to 0
func (set Bitset) ClearBit(i uint) {
	set[i/64] &= (1 << (i % 64)) ^ math.MaxUint64
}

// Clear entire Bitset
func (set Bitset) Clear() {
	for i := range set {
		set[i] = 0
	}
}

// Show ...
func (set Bitset) Show() {
	for i := range set {
		fmt.Printf("%064b\n", set[i])
	}
}
