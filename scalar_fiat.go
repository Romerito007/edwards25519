// Code generated by Fiat Cryptography. DO NOT EDIT.
//
// Autogenerated: word_by_word_montgomery --lang Go --cmovznz-by-mul --relax-primitive-carry-to-bitwidth 32,64 --public-function-case camelCase --public-type-case camelCase --private-function-case camelCase --private-type-case camelCase --doc-text-before-function-name '' --doc-newline-before-package-declaration --doc-prepend-header 'Code generated by Fiat Cryptography. DO NOT EDIT.' --package-name edwards25519 Scalar 64 '2^252 + 27742317777372353535851937790883648493' mul add sub opp nonzero from_montgomery to_montgomery to_bytes from_bytes
//
// curve description: Scalar
//
// machine_wordsize = 64 (from "64")
//
// requested operations: mul, add, sub, opp, nonzero, from_montgomery, to_montgomery, to_bytes, from_bytes
//
// m = 0x1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed (from "2^252 + 27742317777372353535851937790883648493")
//
//
//
// NOTE: In addition to the bounds specified above each function, all
//
//   functions synthesized for this Montgomery arithmetic require the
//
//   input to be strictly less than the prime modulus (m), and also
//
//   require the input to be in the unique saturated representation.
//
//   All functions also ensure that these two properties are true of
//
//   return values.
//
//
//
// Computed values:
//
//   eval z = z[0] + (z[1] << 64) + (z[2] << 128) + (z[3] << 192)
//
//   bytes_eval z = z[0] + (z[1] << 8) + (z[2] << 16) + (z[3] << 24) + (z[4] << 32) + (z[5] << 40) + (z[6] << 48) + (z[7] << 56) + (z[8] << 64) + (z[9] << 72) + (z[10] << 80) + (z[11] << 88) + (z[12] << 96) + (z[13] << 104) + (z[14] << 112) + (z[15] << 120) + (z[16] << 128) + (z[17] << 136) + (z[18] << 144) + (z[19] << 152) + (z[20] << 160) + (z[21] << 168) + (z[22] << 176) + (z[23] << 184) + (z[24] << 192) + (z[25] << 200) + (z[26] << 208) + (z[27] << 216) + (z[28] << 224) + (z[29] << 232) + (z[30] << 240) + (z[31] << 248)
//
//   twos_complement_eval z = let x1 := z[0] + (z[1] << 64) + (z[2] << 128) + (z[3] << 192) in
//
//                            if x1 & (2^256-1) < 2^255 then x1 & (2^256-1) else (x1 & (2^256-1)) - 2^256

package edwards25519

import "math/bits"

type fiatScalarUint1 uint64 // We use uint64 instead of a more narrow type for performance reasons; see https://github.com/mit-plv/fiat-crypto/pull/1006#issuecomment-892625927
type fiatScalarInt1 int64 // We use uint64 instead of a more narrow type for performance reasons; see https://github.com/mit-plv/fiat-crypto/pull/1006#issuecomment-892625927

// The type fiatScalarMontgomeryDomainFieldElement is a field element in the Montgomery domain.
//
// Bounds: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
type fiatScalarMontgomeryDomainFieldElement [4]uint64

// The type fiatScalarNonMontgomeryDomainFieldElement is a field element NOT in the Montgomery domain.
//
// Bounds: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
type fiatScalarNonMontgomeryDomainFieldElement [4]uint64

// fiatScalarCmovznzU64 is a single-word conditional move.
//
// Postconditions:
//   out1 = (if arg1 = 0 then arg2 else arg3)
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0xffffffffffffffff]
//   arg3: [0x0 ~> 0xffffffffffffffff]
// Output Bounds:
//   out1: [0x0 ~> 0xffffffffffffffff]
func fiatScalarCmovznzU64(out1 *uint64, arg1 fiatScalarUint1, arg2 uint64, arg3 uint64) {
	x1 := (uint64(arg1) * 0xffffffffffffffff)
	x2 := ((x1 & arg3) | ((^x1) & arg2))
	*out1 = x2
}

// fiatScalarMul multiplies two field elements in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
//   0 ≤ eval arg2 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) * eval (from_montgomery arg2)) mod m
//   0 ≤ eval out1 < m
//
func fiatScalarMul(out1 *fiatScalarMontgomeryDomainFieldElement, arg1 *fiatScalarMontgomeryDomainFieldElement, arg2 *fiatScalarMontgomeryDomainFieldElement) {
	x1 := arg1[1]
	x2 := arg1[2]
	x3 := arg1[3]
	x4 := arg1[0]
	var x5 uint64
	var x6 uint64
	x6, x5 = bits.Mul64(x4, arg2[3])
	var x7 uint64
	var x8 uint64
	x8, x7 = bits.Mul64(x4, arg2[2])
	var x9 uint64
	var x10 uint64
	x10, x9 = bits.Mul64(x4, arg2[1])
	var x11 uint64
	var x12 uint64
	x12, x11 = bits.Mul64(x4, arg2[0])
	var x13 uint64
	var x14 uint64
	x13, x14 = bits.Add64(x12, x9, uint64(0x0))
	var x15 uint64
	var x16 uint64
	x15, x16 = bits.Add64(x10, x7, uint64(fiatScalarUint1(x14)))
	var x17 uint64
	var x18 uint64
	x17, x18 = bits.Add64(x8, x5, uint64(fiatScalarUint1(x16)))
	x19 := (uint64(fiatScalarUint1(x18)) + x6)
	var x20 uint64
	_, x20 = bits.Mul64(x11, 0xd2b51da312547e1b)
	var x22 uint64
	var x23 uint64
	x23, x22 = bits.Mul64(x20, 0x1000000000000000)
	var x24 uint64
	var x25 uint64
	x25, x24 = bits.Mul64(x20, 0x14def9dea2f79cd6)
	var x26 uint64
	var x27 uint64
	x27, x26 = bits.Mul64(x20, 0x5812631a5cf5d3ed)
	var x28 uint64
	var x29 uint64
	x28, x29 = bits.Add64(x27, x24, uint64(0x0))
	x30 := (uint64(fiatScalarUint1(x29)) + x25)
	var x32 uint64
	_, x32 = bits.Add64(x11, x26, uint64(0x0))
	var x33 uint64
	var x34 uint64
	x33, x34 = bits.Add64(x13, x28, uint64(fiatScalarUint1(x32)))
	var x35 uint64
	var x36 uint64
	x35, x36 = bits.Add64(x15, x30, uint64(fiatScalarUint1(x34)))
	var x37 uint64
	var x38 uint64
	x37, x38 = bits.Add64(x17, x22, uint64(fiatScalarUint1(x36)))
	var x39 uint64
	var x40 uint64
	x39, x40 = bits.Add64(x19, x23, uint64(fiatScalarUint1(x38)))
	var x41 uint64
	var x42 uint64
	x42, x41 = bits.Mul64(x1, arg2[3])
	var x43 uint64
	var x44 uint64
	x44, x43 = bits.Mul64(x1, arg2[2])
	var x45 uint64
	var x46 uint64
	x46, x45 = bits.Mul64(x1, arg2[1])
	var x47 uint64
	var x48 uint64
	x48, x47 = bits.Mul64(x1, arg2[0])
	var x49 uint64
	var x50 uint64
	x49, x50 = bits.Add64(x48, x45, uint64(0x0))
	var x51 uint64
	var x52 uint64
	x51, x52 = bits.Add64(x46, x43, uint64(fiatScalarUint1(x50)))
	var x53 uint64
	var x54 uint64
	x53, x54 = bits.Add64(x44, x41, uint64(fiatScalarUint1(x52)))
	x55 := (uint64(fiatScalarUint1(x54)) + x42)
	var x56 uint64
	var x57 uint64
	x56, x57 = bits.Add64(x33, x47, uint64(0x0))
	var x58 uint64
	var x59 uint64
	x58, x59 = bits.Add64(x35, x49, uint64(fiatScalarUint1(x57)))
	var x60 uint64
	var x61 uint64
	x60, x61 = bits.Add64(x37, x51, uint64(fiatScalarUint1(x59)))
	var x62 uint64
	var x63 uint64
	x62, x63 = bits.Add64(x39, x53, uint64(fiatScalarUint1(x61)))
	var x64 uint64
	var x65 uint64
	x64, x65 = bits.Add64(uint64(fiatScalarUint1(x40)), x55, uint64(fiatScalarUint1(x63)))
	var x66 uint64
	_, x66 = bits.Mul64(x56, 0xd2b51da312547e1b)
	var x68 uint64
	var x69 uint64
	x69, x68 = bits.Mul64(x66, 0x1000000000000000)
	var x70 uint64
	var x71 uint64
	x71, x70 = bits.Mul64(x66, 0x14def9dea2f79cd6)
	var x72 uint64
	var x73 uint64
	x73, x72 = bits.Mul64(x66, 0x5812631a5cf5d3ed)
	var x74 uint64
	var x75 uint64
	x74, x75 = bits.Add64(x73, x70, uint64(0x0))
	x76 := (uint64(fiatScalarUint1(x75)) + x71)
	var x78 uint64
	_, x78 = bits.Add64(x56, x72, uint64(0x0))
	var x79 uint64
	var x80 uint64
	x79, x80 = bits.Add64(x58, x74, uint64(fiatScalarUint1(x78)))
	var x81 uint64
	var x82 uint64
	x81, x82 = bits.Add64(x60, x76, uint64(fiatScalarUint1(x80)))
	var x83 uint64
	var x84 uint64
	x83, x84 = bits.Add64(x62, x68, uint64(fiatScalarUint1(x82)))
	var x85 uint64
	var x86 uint64
	x85, x86 = bits.Add64(x64, x69, uint64(fiatScalarUint1(x84)))
	x87 := (uint64(fiatScalarUint1(x86)) + uint64(fiatScalarUint1(x65)))
	var x88 uint64
	var x89 uint64
	x89, x88 = bits.Mul64(x2, arg2[3])
	var x90 uint64
	var x91 uint64
	x91, x90 = bits.Mul64(x2, arg2[2])
	var x92 uint64
	var x93 uint64
	x93, x92 = bits.Mul64(x2, arg2[1])
	var x94 uint64
	var x95 uint64
	x95, x94 = bits.Mul64(x2, arg2[0])
	var x96 uint64
	var x97 uint64
	x96, x97 = bits.Add64(x95, x92, uint64(0x0))
	var x98 uint64
	var x99 uint64
	x98, x99 = bits.Add64(x93, x90, uint64(fiatScalarUint1(x97)))
	var x100 uint64
	var x101 uint64
	x100, x101 = bits.Add64(x91, x88, uint64(fiatScalarUint1(x99)))
	x102 := (uint64(fiatScalarUint1(x101)) + x89)
	var x103 uint64
	var x104 uint64
	x103, x104 = bits.Add64(x79, x94, uint64(0x0))
	var x105 uint64
	var x106 uint64
	x105, x106 = bits.Add64(x81, x96, uint64(fiatScalarUint1(x104)))
	var x107 uint64
	var x108 uint64
	x107, x108 = bits.Add64(x83, x98, uint64(fiatScalarUint1(x106)))
	var x109 uint64
	var x110 uint64
	x109, x110 = bits.Add64(x85, x100, uint64(fiatScalarUint1(x108)))
	var x111 uint64
	var x112 uint64
	x111, x112 = bits.Add64(x87, x102, uint64(fiatScalarUint1(x110)))
	var x113 uint64
	_, x113 = bits.Mul64(x103, 0xd2b51da312547e1b)
	var x115 uint64
	var x116 uint64
	x116, x115 = bits.Mul64(x113, 0x1000000000000000)
	var x117 uint64
	var x118 uint64
	x118, x117 = bits.Mul64(x113, 0x14def9dea2f79cd6)
	var x119 uint64
	var x120 uint64
	x120, x119 = bits.Mul64(x113, 0x5812631a5cf5d3ed)
	var x121 uint64
	var x122 uint64
	x121, x122 = bits.Add64(x120, x117, uint64(0x0))
	x123 := (uint64(fiatScalarUint1(x122)) + x118)
	var x125 uint64
	_, x125 = bits.Add64(x103, x119, uint64(0x0))
	var x126 uint64
	var x127 uint64
	x126, x127 = bits.Add64(x105, x121, uint64(fiatScalarUint1(x125)))
	var x128 uint64
	var x129 uint64
	x128, x129 = bits.Add64(x107, x123, uint64(fiatScalarUint1(x127)))
	var x130 uint64
	var x131 uint64
	x130, x131 = bits.Add64(x109, x115, uint64(fiatScalarUint1(x129)))
	var x132 uint64
	var x133 uint64
	x132, x133 = bits.Add64(x111, x116, uint64(fiatScalarUint1(x131)))
	x134 := (uint64(fiatScalarUint1(x133)) + uint64(fiatScalarUint1(x112)))
	var x135 uint64
	var x136 uint64
	x136, x135 = bits.Mul64(x3, arg2[3])
	var x137 uint64
	var x138 uint64
	x138, x137 = bits.Mul64(x3, arg2[2])
	var x139 uint64
	var x140 uint64
	x140, x139 = bits.Mul64(x3, arg2[1])
	var x141 uint64
	var x142 uint64
	x142, x141 = bits.Mul64(x3, arg2[0])
	var x143 uint64
	var x144 uint64
	x143, x144 = bits.Add64(x142, x139, uint64(0x0))
	var x145 uint64
	var x146 uint64
	x145, x146 = bits.Add64(x140, x137, uint64(fiatScalarUint1(x144)))
	var x147 uint64
	var x148 uint64
	x147, x148 = bits.Add64(x138, x135, uint64(fiatScalarUint1(x146)))
	x149 := (uint64(fiatScalarUint1(x148)) + x136)
	var x150 uint64
	var x151 uint64
	x150, x151 = bits.Add64(x126, x141, uint64(0x0))
	var x152 uint64
	var x153 uint64
	x152, x153 = bits.Add64(x128, x143, uint64(fiatScalarUint1(x151)))
	var x154 uint64
	var x155 uint64
	x154, x155 = bits.Add64(x130, x145, uint64(fiatScalarUint1(x153)))
	var x156 uint64
	var x157 uint64
	x156, x157 = bits.Add64(x132, x147, uint64(fiatScalarUint1(x155)))
	var x158 uint64
	var x159 uint64
	x158, x159 = bits.Add64(x134, x149, uint64(fiatScalarUint1(x157)))
	var x160 uint64
	_, x160 = bits.Mul64(x150, 0xd2b51da312547e1b)
	var x162 uint64
	var x163 uint64
	x163, x162 = bits.Mul64(x160, 0x1000000000000000)
	var x164 uint64
	var x165 uint64
	x165, x164 = bits.Mul64(x160, 0x14def9dea2f79cd6)
	var x166 uint64
	var x167 uint64
	x167, x166 = bits.Mul64(x160, 0x5812631a5cf5d3ed)
	var x168 uint64
	var x169 uint64
	x168, x169 = bits.Add64(x167, x164, uint64(0x0))
	x170 := (uint64(fiatScalarUint1(x169)) + x165)
	var x172 uint64
	_, x172 = bits.Add64(x150, x166, uint64(0x0))
	var x173 uint64
	var x174 uint64
	x173, x174 = bits.Add64(x152, x168, uint64(fiatScalarUint1(x172)))
	var x175 uint64
	var x176 uint64
	x175, x176 = bits.Add64(x154, x170, uint64(fiatScalarUint1(x174)))
	var x177 uint64
	var x178 uint64
	x177, x178 = bits.Add64(x156, x162, uint64(fiatScalarUint1(x176)))
	var x179 uint64
	var x180 uint64
	x179, x180 = bits.Add64(x158, x163, uint64(fiatScalarUint1(x178)))
	x181 := (uint64(fiatScalarUint1(x180)) + uint64(fiatScalarUint1(x159)))
	var x182 uint64
	var x183 uint64
	x182, x183 = bits.Sub64(x173, 0x5812631a5cf5d3ed, uint64(0x0))
	var x184 uint64
	var x185 uint64
	x184, x185 = bits.Sub64(x175, 0x14def9dea2f79cd6, uint64(fiatScalarUint1(x183)))
	var x186 uint64
	var x187 uint64
	x186, x187 = bits.Sub64(x177, uint64(0x0), uint64(fiatScalarUint1(x185)))
	var x188 uint64
	var x189 uint64
	x188, x189 = bits.Sub64(x179, 0x1000000000000000, uint64(fiatScalarUint1(x187)))
	var x191 uint64
	_, x191 = bits.Sub64(x181, uint64(0x0), uint64(fiatScalarUint1(x189)))
	var x192 uint64
	fiatScalarCmovznzU64(&x192, fiatScalarUint1(x191), x182, x173)
	var x193 uint64
	fiatScalarCmovznzU64(&x193, fiatScalarUint1(x191), x184, x175)
	var x194 uint64
	fiatScalarCmovznzU64(&x194, fiatScalarUint1(x191), x186, x177)
	var x195 uint64
	fiatScalarCmovznzU64(&x195, fiatScalarUint1(x191), x188, x179)
	out1[0] = x192
	out1[1] = x193
	out1[2] = x194
	out1[3] = x195
}

// fiatScalarAdd adds two field elements in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
//   0 ≤ eval arg2 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) + eval (from_montgomery arg2)) mod m
//   0 ≤ eval out1 < m
//
func fiatScalarAdd(out1 *fiatScalarMontgomeryDomainFieldElement, arg1 *fiatScalarMontgomeryDomainFieldElement, arg2 *fiatScalarMontgomeryDomainFieldElement) {
	var x1 uint64
	var x2 uint64
	x1, x2 = bits.Add64(arg1[0], arg2[0], uint64(0x0))
	var x3 uint64
	var x4 uint64
	x3, x4 = bits.Add64(arg1[1], arg2[1], uint64(fiatScalarUint1(x2)))
	var x5 uint64
	var x6 uint64
	x5, x6 = bits.Add64(arg1[2], arg2[2], uint64(fiatScalarUint1(x4)))
	var x7 uint64
	var x8 uint64
	x7, x8 = bits.Add64(arg1[3], arg2[3], uint64(fiatScalarUint1(x6)))
	var x9 uint64
	var x10 uint64
	x9, x10 = bits.Sub64(x1, 0x5812631a5cf5d3ed, uint64(0x0))
	var x11 uint64
	var x12 uint64
	x11, x12 = bits.Sub64(x3, 0x14def9dea2f79cd6, uint64(fiatScalarUint1(x10)))
	var x13 uint64
	var x14 uint64
	x13, x14 = bits.Sub64(x5, uint64(0x0), uint64(fiatScalarUint1(x12)))
	var x15 uint64
	var x16 uint64
	x15, x16 = bits.Sub64(x7, 0x1000000000000000, uint64(fiatScalarUint1(x14)))
	var x18 uint64
	_, x18 = bits.Sub64(uint64(fiatScalarUint1(x8)), uint64(0x0), uint64(fiatScalarUint1(x16)))
	var x19 uint64
	fiatScalarCmovznzU64(&x19, fiatScalarUint1(x18), x9, x1)
	var x20 uint64
	fiatScalarCmovznzU64(&x20, fiatScalarUint1(x18), x11, x3)
	var x21 uint64
	fiatScalarCmovznzU64(&x21, fiatScalarUint1(x18), x13, x5)
	var x22 uint64
	fiatScalarCmovznzU64(&x22, fiatScalarUint1(x18), x15, x7)
	out1[0] = x19
	out1[1] = x20
	out1[2] = x21
	out1[3] = x22
}

// fiatScalarSub subtracts two field elements in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
//   0 ≤ eval arg2 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) - eval (from_montgomery arg2)) mod m
//   0 ≤ eval out1 < m
//
func fiatScalarSub(out1 *fiatScalarMontgomeryDomainFieldElement, arg1 *fiatScalarMontgomeryDomainFieldElement, arg2 *fiatScalarMontgomeryDomainFieldElement) {
	var x1 uint64
	var x2 uint64
	x1, x2 = bits.Sub64(arg1[0], arg2[0], uint64(0x0))
	var x3 uint64
	var x4 uint64
	x3, x4 = bits.Sub64(arg1[1], arg2[1], uint64(fiatScalarUint1(x2)))
	var x5 uint64
	var x6 uint64
	x5, x6 = bits.Sub64(arg1[2], arg2[2], uint64(fiatScalarUint1(x4)))
	var x7 uint64
	var x8 uint64
	x7, x8 = bits.Sub64(arg1[3], arg2[3], uint64(fiatScalarUint1(x6)))
	var x9 uint64
	fiatScalarCmovznzU64(&x9, fiatScalarUint1(x8), uint64(0x0), 0xffffffffffffffff)
	var x10 uint64
	var x11 uint64
	x10, x11 = bits.Add64(x1, (x9 & 0x5812631a5cf5d3ed), uint64(0x0))
	var x12 uint64
	var x13 uint64
	x12, x13 = bits.Add64(x3, (x9 & 0x14def9dea2f79cd6), uint64(fiatScalarUint1(x11)))
	var x14 uint64
	var x15 uint64
	x14, x15 = bits.Add64(x5, uint64(0x0), uint64(fiatScalarUint1(x13)))
	var x16 uint64
	x16, _ = bits.Add64(x7, (x9 & 0x1000000000000000), uint64(fiatScalarUint1(x15)))
	out1[0] = x10
	out1[1] = x12
	out1[2] = x14
	out1[3] = x16
}

// fiatScalarOpp negates a field element in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = -eval (from_montgomery arg1) mod m
//   0 ≤ eval out1 < m
//
func fiatScalarOpp(out1 *fiatScalarMontgomeryDomainFieldElement, arg1 *fiatScalarMontgomeryDomainFieldElement) {
	var x1 uint64
	var x2 uint64
	x1, x2 = bits.Sub64(uint64(0x0), arg1[0], uint64(0x0))
	var x3 uint64
	var x4 uint64
	x3, x4 = bits.Sub64(uint64(0x0), arg1[1], uint64(fiatScalarUint1(x2)))
	var x5 uint64
	var x6 uint64
	x5, x6 = bits.Sub64(uint64(0x0), arg1[2], uint64(fiatScalarUint1(x4)))
	var x7 uint64
	var x8 uint64
	x7, x8 = bits.Sub64(uint64(0x0), arg1[3], uint64(fiatScalarUint1(x6)))
	var x9 uint64
	fiatScalarCmovznzU64(&x9, fiatScalarUint1(x8), uint64(0x0), 0xffffffffffffffff)
	var x10 uint64
	var x11 uint64
	x10, x11 = bits.Add64(x1, (x9 & 0x5812631a5cf5d3ed), uint64(0x0))
	var x12 uint64
	var x13 uint64
	x12, x13 = bits.Add64(x3, (x9 & 0x14def9dea2f79cd6), uint64(fiatScalarUint1(x11)))
	var x14 uint64
	var x15 uint64
	x14, x15 = bits.Add64(x5, uint64(0x0), uint64(fiatScalarUint1(x13)))
	var x16 uint64
	x16, _ = bits.Add64(x7, (x9 & 0x1000000000000000), uint64(fiatScalarUint1(x15)))
	out1[0] = x10
	out1[1] = x12
	out1[2] = x14
	out1[3] = x16
}

// fiatScalarNonzero outputs a single non-zero word if the input is non-zero and zero otherwise.
//
// Preconditions:
//   0 ≤ eval arg1 < m
// Postconditions:
//   out1 = 0 ↔ eval (from_montgomery arg1) mod m = 0
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
// Output Bounds:
//   out1: [0x0 ~> 0xffffffffffffffff]
func fiatScalarNonzero(out1 *uint64, arg1 *[4]uint64) {
	x1 := (arg1[0] | (arg1[1] | (arg1[2] | arg1[3])))
	*out1 = x1
}

// fiatScalarFromMontgomery translates a field element out of the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
// Postconditions:
//   eval out1 mod m = (eval arg1 * ((2^64)⁻¹ mod m)^4) mod m
//   0 ≤ eval out1 < m
//
func fiatScalarFromMontgomery(out1 *fiatScalarNonMontgomeryDomainFieldElement, arg1 *fiatScalarMontgomeryDomainFieldElement) {
	x1 := arg1[0]
	var x2 uint64
	_, x2 = bits.Mul64(x1, 0xd2b51da312547e1b)
	var x4 uint64
	var x5 uint64
	x5, x4 = bits.Mul64(x2, 0x1000000000000000)
	var x6 uint64
	var x7 uint64
	x7, x6 = bits.Mul64(x2, 0x14def9dea2f79cd6)
	var x8 uint64
	var x9 uint64
	x9, x8 = bits.Mul64(x2, 0x5812631a5cf5d3ed)
	var x10 uint64
	var x11 uint64
	x10, x11 = bits.Add64(x9, x6, uint64(0x0))
	var x13 uint64
	_, x13 = bits.Add64(x1, x8, uint64(0x0))
	var x14 uint64
	var x15 uint64
	x14, x15 = bits.Add64(uint64(0x0), x10, uint64(fiatScalarUint1(x13)))
	var x16 uint64
	var x17 uint64
	x16, x17 = bits.Add64(x14, arg1[1], uint64(0x0))
	var x18 uint64
	_, x18 = bits.Mul64(x16, 0xd2b51da312547e1b)
	var x20 uint64
	var x21 uint64
	x21, x20 = bits.Mul64(x18, 0x1000000000000000)
	var x22 uint64
	var x23 uint64
	x23, x22 = bits.Mul64(x18, 0x14def9dea2f79cd6)
	var x24 uint64
	var x25 uint64
	x25, x24 = bits.Mul64(x18, 0x5812631a5cf5d3ed)
	var x26 uint64
	var x27 uint64
	x26, x27 = bits.Add64(x25, x22, uint64(0x0))
	var x29 uint64
	_, x29 = bits.Add64(x16, x24, uint64(0x0))
	var x30 uint64
	var x31 uint64
	x30, x31 = bits.Add64((uint64(fiatScalarUint1(x17)) + (uint64(fiatScalarUint1(x15)) + (uint64(fiatScalarUint1(x11)) + x7))), x26, uint64(fiatScalarUint1(x29)))
	var x32 uint64
	var x33 uint64
	x32, x33 = bits.Add64(x4, (uint64(fiatScalarUint1(x27)) + x23), uint64(fiatScalarUint1(x31)))
	var x34 uint64
	var x35 uint64
	x34, x35 = bits.Add64(x5, x20, uint64(fiatScalarUint1(x33)))
	var x36 uint64
	var x37 uint64
	x36, x37 = bits.Add64(x30, arg1[2], uint64(0x0))
	var x38 uint64
	var x39 uint64
	x38, x39 = bits.Add64(x32, uint64(0x0), uint64(fiatScalarUint1(x37)))
	var x40 uint64
	var x41 uint64
	x40, x41 = bits.Add64(x34, uint64(0x0), uint64(fiatScalarUint1(x39)))
	var x42 uint64
	_, x42 = bits.Mul64(x36, 0xd2b51da312547e1b)
	var x44 uint64
	var x45 uint64
	x45, x44 = bits.Mul64(x42, 0x1000000000000000)
	var x46 uint64
	var x47 uint64
	x47, x46 = bits.Mul64(x42, 0x14def9dea2f79cd6)
	var x48 uint64
	var x49 uint64
	x49, x48 = bits.Mul64(x42, 0x5812631a5cf5d3ed)
	var x50 uint64
	var x51 uint64
	x50, x51 = bits.Add64(x49, x46, uint64(0x0))
	var x53 uint64
	_, x53 = bits.Add64(x36, x48, uint64(0x0))
	var x54 uint64
	var x55 uint64
	x54, x55 = bits.Add64(x38, x50, uint64(fiatScalarUint1(x53)))
	var x56 uint64
	var x57 uint64
	x56, x57 = bits.Add64(x40, (uint64(fiatScalarUint1(x51)) + x47), uint64(fiatScalarUint1(x55)))
	var x58 uint64
	var x59 uint64
	x58, x59 = bits.Add64((uint64(fiatScalarUint1(x41)) + (uint64(fiatScalarUint1(x35)) + x21)), x44, uint64(fiatScalarUint1(x57)))
	var x60 uint64
	var x61 uint64
	x60, x61 = bits.Add64(x54, arg1[3], uint64(0x0))
	var x62 uint64
	var x63 uint64
	x62, x63 = bits.Add64(x56, uint64(0x0), uint64(fiatScalarUint1(x61)))
	var x64 uint64
	var x65 uint64
	x64, x65 = bits.Add64(x58, uint64(0x0), uint64(fiatScalarUint1(x63)))
	var x66 uint64
	_, x66 = bits.Mul64(x60, 0xd2b51da312547e1b)
	var x68 uint64
	var x69 uint64
	x69, x68 = bits.Mul64(x66, 0x1000000000000000)
	var x70 uint64
	var x71 uint64
	x71, x70 = bits.Mul64(x66, 0x14def9dea2f79cd6)
	var x72 uint64
	var x73 uint64
	x73, x72 = bits.Mul64(x66, 0x5812631a5cf5d3ed)
	var x74 uint64
	var x75 uint64
	x74, x75 = bits.Add64(x73, x70, uint64(0x0))
	var x77 uint64
	_, x77 = bits.Add64(x60, x72, uint64(0x0))
	var x78 uint64
	var x79 uint64
	x78, x79 = bits.Add64(x62, x74, uint64(fiatScalarUint1(x77)))
	var x80 uint64
	var x81 uint64
	x80, x81 = bits.Add64(x64, (uint64(fiatScalarUint1(x75)) + x71), uint64(fiatScalarUint1(x79)))
	var x82 uint64
	var x83 uint64
	x82, x83 = bits.Add64((uint64(fiatScalarUint1(x65)) + (uint64(fiatScalarUint1(x59)) + x45)), x68, uint64(fiatScalarUint1(x81)))
	x84 := (uint64(fiatScalarUint1(x83)) + x69)
	var x85 uint64
	var x86 uint64
	x85, x86 = bits.Sub64(x78, 0x5812631a5cf5d3ed, uint64(0x0))
	var x87 uint64
	var x88 uint64
	x87, x88 = bits.Sub64(x80, 0x14def9dea2f79cd6, uint64(fiatScalarUint1(x86)))
	var x89 uint64
	var x90 uint64
	x89, x90 = bits.Sub64(x82, uint64(0x0), uint64(fiatScalarUint1(x88)))
	var x91 uint64
	var x92 uint64
	x91, x92 = bits.Sub64(x84, 0x1000000000000000, uint64(fiatScalarUint1(x90)))
	var x94 uint64
	_, x94 = bits.Sub64(uint64(0x0), uint64(0x0), uint64(fiatScalarUint1(x92)))
	var x95 uint64
	fiatScalarCmovznzU64(&x95, fiatScalarUint1(x94), x85, x78)
	var x96 uint64
	fiatScalarCmovznzU64(&x96, fiatScalarUint1(x94), x87, x80)
	var x97 uint64
	fiatScalarCmovznzU64(&x97, fiatScalarUint1(x94), x89, x82)
	var x98 uint64
	fiatScalarCmovznzU64(&x98, fiatScalarUint1(x94), x91, x84)
	out1[0] = x95
	out1[1] = x96
	out1[2] = x97
	out1[3] = x98
}

// fiatScalarToMontgomery translates a field element into the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = eval arg1 mod m
//   0 ≤ eval out1 < m
//
func fiatScalarToMontgomery(out1 *fiatScalarMontgomeryDomainFieldElement, arg1 *fiatScalarNonMontgomeryDomainFieldElement) {
	x1 := arg1[1]
	x2 := arg1[2]
	x3 := arg1[3]
	x4 := arg1[0]
	var x5 uint64
	var x6 uint64
	x6, x5 = bits.Mul64(x4, 0x399411b7c309a3d)
	var x7 uint64
	var x8 uint64
	x8, x7 = bits.Mul64(x4, 0xceec73d217f5be65)
	var x9 uint64
	var x10 uint64
	x10, x9 = bits.Mul64(x4, 0xd00e1ba768859347)
	var x11 uint64
	var x12 uint64
	x12, x11 = bits.Mul64(x4, 0xa40611e3449c0f01)
	var x13 uint64
	var x14 uint64
	x13, x14 = bits.Add64(x12, x9, uint64(0x0))
	var x15 uint64
	var x16 uint64
	x15, x16 = bits.Add64(x10, x7, uint64(fiatScalarUint1(x14)))
	var x17 uint64
	var x18 uint64
	x17, x18 = bits.Add64(x8, x5, uint64(fiatScalarUint1(x16)))
	var x19 uint64
	_, x19 = bits.Mul64(x11, 0xd2b51da312547e1b)
	var x21 uint64
	var x22 uint64
	x22, x21 = bits.Mul64(x19, 0x1000000000000000)
	var x23 uint64
	var x24 uint64
	x24, x23 = bits.Mul64(x19, 0x14def9dea2f79cd6)
	var x25 uint64
	var x26 uint64
	x26, x25 = bits.Mul64(x19, 0x5812631a5cf5d3ed)
	var x27 uint64
	var x28 uint64
	x27, x28 = bits.Add64(x26, x23, uint64(0x0))
	var x30 uint64
	_, x30 = bits.Add64(x11, x25, uint64(0x0))
	var x31 uint64
	var x32 uint64
	x31, x32 = bits.Add64(x13, x27, uint64(fiatScalarUint1(x30)))
	var x33 uint64
	var x34 uint64
	x33, x34 = bits.Add64(x15, (uint64(fiatScalarUint1(x28)) + x24), uint64(fiatScalarUint1(x32)))
	var x35 uint64
	var x36 uint64
	x35, x36 = bits.Add64(x17, x21, uint64(fiatScalarUint1(x34)))
	var x37 uint64
	var x38 uint64
	x38, x37 = bits.Mul64(x1, 0x399411b7c309a3d)
	var x39 uint64
	var x40 uint64
	x40, x39 = bits.Mul64(x1, 0xceec73d217f5be65)
	var x41 uint64
	var x42 uint64
	x42, x41 = bits.Mul64(x1, 0xd00e1ba768859347)
	var x43 uint64
	var x44 uint64
	x44, x43 = bits.Mul64(x1, 0xa40611e3449c0f01)
	var x45 uint64
	var x46 uint64
	x45, x46 = bits.Add64(x44, x41, uint64(0x0))
	var x47 uint64
	var x48 uint64
	x47, x48 = bits.Add64(x42, x39, uint64(fiatScalarUint1(x46)))
	var x49 uint64
	var x50 uint64
	x49, x50 = bits.Add64(x40, x37, uint64(fiatScalarUint1(x48)))
	var x51 uint64
	var x52 uint64
	x51, x52 = bits.Add64(x31, x43, uint64(0x0))
	var x53 uint64
	var x54 uint64
	x53, x54 = bits.Add64(x33, x45, uint64(fiatScalarUint1(x52)))
	var x55 uint64
	var x56 uint64
	x55, x56 = bits.Add64(x35, x47, uint64(fiatScalarUint1(x54)))
	var x57 uint64
	var x58 uint64
	x57, x58 = bits.Add64(((uint64(fiatScalarUint1(x36)) + (uint64(fiatScalarUint1(x18)) + x6)) + x22), x49, uint64(fiatScalarUint1(x56)))
	var x59 uint64
	_, x59 = bits.Mul64(x51, 0xd2b51da312547e1b)
	var x61 uint64
	var x62 uint64
	x62, x61 = bits.Mul64(x59, 0x1000000000000000)
	var x63 uint64
	var x64 uint64
	x64, x63 = bits.Mul64(x59, 0x14def9dea2f79cd6)
	var x65 uint64
	var x66 uint64
	x66, x65 = bits.Mul64(x59, 0x5812631a5cf5d3ed)
	var x67 uint64
	var x68 uint64
	x67, x68 = bits.Add64(x66, x63, uint64(0x0))
	var x70 uint64
	_, x70 = bits.Add64(x51, x65, uint64(0x0))
	var x71 uint64
	var x72 uint64
	x71, x72 = bits.Add64(x53, x67, uint64(fiatScalarUint1(x70)))
	var x73 uint64
	var x74 uint64
	x73, x74 = bits.Add64(x55, (uint64(fiatScalarUint1(x68)) + x64), uint64(fiatScalarUint1(x72)))
	var x75 uint64
	var x76 uint64
	x75, x76 = bits.Add64(x57, x61, uint64(fiatScalarUint1(x74)))
	var x77 uint64
	var x78 uint64
	x78, x77 = bits.Mul64(x2, 0x399411b7c309a3d)
	var x79 uint64
	var x80 uint64
	x80, x79 = bits.Mul64(x2, 0xceec73d217f5be65)
	var x81 uint64
	var x82 uint64
	x82, x81 = bits.Mul64(x2, 0xd00e1ba768859347)
	var x83 uint64
	var x84 uint64
	x84, x83 = bits.Mul64(x2, 0xa40611e3449c0f01)
	var x85 uint64
	var x86 uint64
	x85, x86 = bits.Add64(x84, x81, uint64(0x0))
	var x87 uint64
	var x88 uint64
	x87, x88 = bits.Add64(x82, x79, uint64(fiatScalarUint1(x86)))
	var x89 uint64
	var x90 uint64
	x89, x90 = bits.Add64(x80, x77, uint64(fiatScalarUint1(x88)))
	var x91 uint64
	var x92 uint64
	x91, x92 = bits.Add64(x71, x83, uint64(0x0))
	var x93 uint64
	var x94 uint64
	x93, x94 = bits.Add64(x73, x85, uint64(fiatScalarUint1(x92)))
	var x95 uint64
	var x96 uint64
	x95, x96 = bits.Add64(x75, x87, uint64(fiatScalarUint1(x94)))
	var x97 uint64
	var x98 uint64
	x97, x98 = bits.Add64(((uint64(fiatScalarUint1(x76)) + (uint64(fiatScalarUint1(x58)) + (uint64(fiatScalarUint1(x50)) + x38))) + x62), x89, uint64(fiatScalarUint1(x96)))
	var x99 uint64
	_, x99 = bits.Mul64(x91, 0xd2b51da312547e1b)
	var x101 uint64
	var x102 uint64
	x102, x101 = bits.Mul64(x99, 0x1000000000000000)
	var x103 uint64
	var x104 uint64
	x104, x103 = bits.Mul64(x99, 0x14def9dea2f79cd6)
	var x105 uint64
	var x106 uint64
	x106, x105 = bits.Mul64(x99, 0x5812631a5cf5d3ed)
	var x107 uint64
	var x108 uint64
	x107, x108 = bits.Add64(x106, x103, uint64(0x0))
	var x110 uint64
	_, x110 = bits.Add64(x91, x105, uint64(0x0))
	var x111 uint64
	var x112 uint64
	x111, x112 = bits.Add64(x93, x107, uint64(fiatScalarUint1(x110)))
	var x113 uint64
	var x114 uint64
	x113, x114 = bits.Add64(x95, (uint64(fiatScalarUint1(x108)) + x104), uint64(fiatScalarUint1(x112)))
	var x115 uint64
	var x116 uint64
	x115, x116 = bits.Add64(x97, x101, uint64(fiatScalarUint1(x114)))
	var x117 uint64
	var x118 uint64
	x118, x117 = bits.Mul64(x3, 0x399411b7c309a3d)
	var x119 uint64
	var x120 uint64
	x120, x119 = bits.Mul64(x3, 0xceec73d217f5be65)
	var x121 uint64
	var x122 uint64
	x122, x121 = bits.Mul64(x3, 0xd00e1ba768859347)
	var x123 uint64
	var x124 uint64
	x124, x123 = bits.Mul64(x3, 0xa40611e3449c0f01)
	var x125 uint64
	var x126 uint64
	x125, x126 = bits.Add64(x124, x121, uint64(0x0))
	var x127 uint64
	var x128 uint64
	x127, x128 = bits.Add64(x122, x119, uint64(fiatScalarUint1(x126)))
	var x129 uint64
	var x130 uint64
	x129, x130 = bits.Add64(x120, x117, uint64(fiatScalarUint1(x128)))
	var x131 uint64
	var x132 uint64
	x131, x132 = bits.Add64(x111, x123, uint64(0x0))
	var x133 uint64
	var x134 uint64
	x133, x134 = bits.Add64(x113, x125, uint64(fiatScalarUint1(x132)))
	var x135 uint64
	var x136 uint64
	x135, x136 = bits.Add64(x115, x127, uint64(fiatScalarUint1(x134)))
	var x137 uint64
	var x138 uint64
	x137, x138 = bits.Add64(((uint64(fiatScalarUint1(x116)) + (uint64(fiatScalarUint1(x98)) + (uint64(fiatScalarUint1(x90)) + x78))) + x102), x129, uint64(fiatScalarUint1(x136)))
	var x139 uint64
	_, x139 = bits.Mul64(x131, 0xd2b51da312547e1b)
	var x141 uint64
	var x142 uint64
	x142, x141 = bits.Mul64(x139, 0x1000000000000000)
	var x143 uint64
	var x144 uint64
	x144, x143 = bits.Mul64(x139, 0x14def9dea2f79cd6)
	var x145 uint64
	var x146 uint64
	x146, x145 = bits.Mul64(x139, 0x5812631a5cf5d3ed)
	var x147 uint64
	var x148 uint64
	x147, x148 = bits.Add64(x146, x143, uint64(0x0))
	var x150 uint64
	_, x150 = bits.Add64(x131, x145, uint64(0x0))
	var x151 uint64
	var x152 uint64
	x151, x152 = bits.Add64(x133, x147, uint64(fiatScalarUint1(x150)))
	var x153 uint64
	var x154 uint64
	x153, x154 = bits.Add64(x135, (uint64(fiatScalarUint1(x148)) + x144), uint64(fiatScalarUint1(x152)))
	var x155 uint64
	var x156 uint64
	x155, x156 = bits.Add64(x137, x141, uint64(fiatScalarUint1(x154)))
	x157 := ((uint64(fiatScalarUint1(x156)) + (uint64(fiatScalarUint1(x138)) + (uint64(fiatScalarUint1(x130)) + x118))) + x142)
	var x158 uint64
	var x159 uint64
	x158, x159 = bits.Sub64(x151, 0x5812631a5cf5d3ed, uint64(0x0))
	var x160 uint64
	var x161 uint64
	x160, x161 = bits.Sub64(x153, 0x14def9dea2f79cd6, uint64(fiatScalarUint1(x159)))
	var x162 uint64
	var x163 uint64
	x162, x163 = bits.Sub64(x155, uint64(0x0), uint64(fiatScalarUint1(x161)))
	var x164 uint64
	var x165 uint64
	x164, x165 = bits.Sub64(x157, 0x1000000000000000, uint64(fiatScalarUint1(x163)))
	var x167 uint64
	_, x167 = bits.Sub64(uint64(0x0), uint64(0x0), uint64(fiatScalarUint1(x165)))
	var x168 uint64
	fiatScalarCmovznzU64(&x168, fiatScalarUint1(x167), x158, x151)
	var x169 uint64
	fiatScalarCmovznzU64(&x169, fiatScalarUint1(x167), x160, x153)
	var x170 uint64
	fiatScalarCmovznzU64(&x170, fiatScalarUint1(x167), x162, x155)
	var x171 uint64
	fiatScalarCmovznzU64(&x171, fiatScalarUint1(x167), x164, x157)
	out1[0] = x168
	out1[1] = x169
	out1[2] = x170
	out1[3] = x171
}

// fiatScalarToBytes serializes a field element NOT in the Montgomery domain to bytes in little-endian order.
//
// Preconditions:
//   0 ≤ eval arg1 < m
// Postconditions:
//   out1 = map (λ x, ⌊((eval arg1 mod m) mod 2^(8 * (x + 1))) / 2^(8 * x)⌋) [0..31]
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0x1fffffffffffffff]]
// Output Bounds:
//   out1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x1f]]
func fiatScalarToBytes(out1 *[32]uint8, arg1 *[4]uint64) {
	x1 := arg1[3]
	x2 := arg1[2]
	x3 := arg1[1]
	x4 := arg1[0]
	x5 := (uint8(x4) & 0xff)
	x6 := (x4 >> 8)
	x7 := (uint8(x6) & 0xff)
	x8 := (x6 >> 8)
	x9 := (uint8(x8) & 0xff)
	x10 := (x8 >> 8)
	x11 := (uint8(x10) & 0xff)
	x12 := (x10 >> 8)
	x13 := (uint8(x12) & 0xff)
	x14 := (x12 >> 8)
	x15 := (uint8(x14) & 0xff)
	x16 := (x14 >> 8)
	x17 := (uint8(x16) & 0xff)
	x18 := uint8((x16 >> 8))
	x19 := (uint8(x3) & 0xff)
	x20 := (x3 >> 8)
	x21 := (uint8(x20) & 0xff)
	x22 := (x20 >> 8)
	x23 := (uint8(x22) & 0xff)
	x24 := (x22 >> 8)
	x25 := (uint8(x24) & 0xff)
	x26 := (x24 >> 8)
	x27 := (uint8(x26) & 0xff)
	x28 := (x26 >> 8)
	x29 := (uint8(x28) & 0xff)
	x30 := (x28 >> 8)
	x31 := (uint8(x30) & 0xff)
	x32 := uint8((x30 >> 8))
	x33 := (uint8(x2) & 0xff)
	x34 := (x2 >> 8)
	x35 := (uint8(x34) & 0xff)
	x36 := (x34 >> 8)
	x37 := (uint8(x36) & 0xff)
	x38 := (x36 >> 8)
	x39 := (uint8(x38) & 0xff)
	x40 := (x38 >> 8)
	x41 := (uint8(x40) & 0xff)
	x42 := (x40 >> 8)
	x43 := (uint8(x42) & 0xff)
	x44 := (x42 >> 8)
	x45 := (uint8(x44) & 0xff)
	x46 := uint8((x44 >> 8))
	x47 := (uint8(x1) & 0xff)
	x48 := (x1 >> 8)
	x49 := (uint8(x48) & 0xff)
	x50 := (x48 >> 8)
	x51 := (uint8(x50) & 0xff)
	x52 := (x50 >> 8)
	x53 := (uint8(x52) & 0xff)
	x54 := (x52 >> 8)
	x55 := (uint8(x54) & 0xff)
	x56 := (x54 >> 8)
	x57 := (uint8(x56) & 0xff)
	x58 := (x56 >> 8)
	x59 := (uint8(x58) & 0xff)
	x60 := uint8((x58 >> 8))
	out1[0] = x5
	out1[1] = x7
	out1[2] = x9
	out1[3] = x11
	out1[4] = x13
	out1[5] = x15
	out1[6] = x17
	out1[7] = x18
	out1[8] = x19
	out1[9] = x21
	out1[10] = x23
	out1[11] = x25
	out1[12] = x27
	out1[13] = x29
	out1[14] = x31
	out1[15] = x32
	out1[16] = x33
	out1[17] = x35
	out1[18] = x37
	out1[19] = x39
	out1[20] = x41
	out1[21] = x43
	out1[22] = x45
	out1[23] = x46
	out1[24] = x47
	out1[25] = x49
	out1[26] = x51
	out1[27] = x53
	out1[28] = x55
	out1[29] = x57
	out1[30] = x59
	out1[31] = x60
}

// fiatScalarFromBytes deserializes a field element NOT in the Montgomery domain from bytes in little-endian order.
//
// Preconditions:
//   0 ≤ bytes_eval arg1 < m
// Postconditions:
//   eval out1 mod m = bytes_eval arg1 mod m
//   0 ≤ eval out1 < m
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x1f]]
// Output Bounds:
//   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0x1fffffffffffffff]]
func fiatScalarFromBytes(out1 *[4]uint64, arg1 *[32]uint8) {
	x1 := (uint64(arg1[31]) << 56)
	x2 := (uint64(arg1[30]) << 48)
	x3 := (uint64(arg1[29]) << 40)
	x4 := (uint64(arg1[28]) << 32)
	x5 := (uint64(arg1[27]) << 24)
	x6 := (uint64(arg1[26]) << 16)
	x7 := (uint64(arg1[25]) << 8)
	x8 := arg1[24]
	x9 := (uint64(arg1[23]) << 56)
	x10 := (uint64(arg1[22]) << 48)
	x11 := (uint64(arg1[21]) << 40)
	x12 := (uint64(arg1[20]) << 32)
	x13 := (uint64(arg1[19]) << 24)
	x14 := (uint64(arg1[18]) << 16)
	x15 := (uint64(arg1[17]) << 8)
	x16 := arg1[16]
	x17 := (uint64(arg1[15]) << 56)
	x18 := (uint64(arg1[14]) << 48)
	x19 := (uint64(arg1[13]) << 40)
	x20 := (uint64(arg1[12]) << 32)
	x21 := (uint64(arg1[11]) << 24)
	x22 := (uint64(arg1[10]) << 16)
	x23 := (uint64(arg1[9]) << 8)
	x24 := arg1[8]
	x25 := (uint64(arg1[7]) << 56)
	x26 := (uint64(arg1[6]) << 48)
	x27 := (uint64(arg1[5]) << 40)
	x28 := (uint64(arg1[4]) << 32)
	x29 := (uint64(arg1[3]) << 24)
	x30 := (uint64(arg1[2]) << 16)
	x31 := (uint64(arg1[1]) << 8)
	x32 := arg1[0]
	x33 := (x31 + uint64(x32))
	x34 := (x30 + x33)
	x35 := (x29 + x34)
	x36 := (x28 + x35)
	x37 := (x27 + x36)
	x38 := (x26 + x37)
	x39 := (x25 + x38)
	x40 := (x23 + uint64(x24))
	x41 := (x22 + x40)
	x42 := (x21 + x41)
	x43 := (x20 + x42)
	x44 := (x19 + x43)
	x45 := (x18 + x44)
	x46 := (x17 + x45)
	x47 := (x15 + uint64(x16))
	x48 := (x14 + x47)
	x49 := (x13 + x48)
	x50 := (x12 + x49)
	x51 := (x11 + x50)
	x52 := (x10 + x51)
	x53 := (x9 + x52)
	x54 := (x7 + uint64(x8))
	x55 := (x6 + x54)
	x56 := (x5 + x55)
	x57 := (x4 + x56)
	x58 := (x3 + x57)
	x59 := (x2 + x58)
	x60 := (x1 + x59)
	out1[0] = x39
	out1[1] = x46
	out1[2] = x53
	out1[3] = x60
}
