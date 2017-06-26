// Copyright 2017 George Tankersley. All rights reserved.

// Constants used in the implementation of GF(2^255-19) field arithmetic.
package radix51

const (
	// The vaule 2^51-1, used in carry propagation
	maskLow51Bits = uint64(1)<<51 - 1
)

var (
	FieldZero FieldElement = [5]uint64{0, 0, 0, 0, 0}
	FieldOne  FieldElement = [5]uint64{1, 0, 0, 0, 0}
)
