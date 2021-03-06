// Package bits provides various constructs that work with bits
package bits // import "github.com/spaskalev/bits"

import (
	"strconv"
)

var hamming = [256]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
}

// Returns the number of raised bits in the given byte
func Hamming(b byte) int {
	return hamming[b]
}

var reversed = [256]byte{
	0, 128, 64, 192, 32, 160, 96, 224, 16, 144, 80, 208, 48, 176, 112, 240,
	8, 136, 72, 200, 40, 168, 104, 232, 24, 152, 88, 216, 56, 184, 120, 248,
	4, 132, 68, 196, 36, 164, 100, 228, 20, 148, 84, 212, 52, 180, 116, 244,
	12, 140, 76, 204, 44, 172, 108, 236, 28, 156, 92, 220, 60, 188, 124, 252,
	2, 130, 66, 194, 34, 162, 98, 226, 18, 146, 82, 210, 50, 178, 114, 242,
	10, 138, 74, 202, 42, 170, 106, 234, 26, 154, 90, 218, 58, 186, 122, 250,
	6, 134, 70, 198, 38, 166, 102, 230, 22, 150, 86, 214, 54, 182, 118, 246,
	14, 142, 78, 206, 46, 174, 110, 238, 30, 158, 94, 222, 62, 190, 126, 254,
	1, 129, 65, 193, 33, 161, 97, 225, 17, 145, 81, 209, 49, 177, 113, 241,
	9, 137, 73, 201, 41, 169, 105, 233, 25, 153, 89, 217, 57, 185, 121, 249,
	5, 133, 69, 197, 37, 165, 101, 229, 21, 149, 85, 213, 53, 181, 117, 245,
	13, 141, 77, 205, 45, 173, 109, 237, 29, 157, 93, 221, 61, 189, 125, 253,
	3, 131, 67, 195, 35, 163, 99, 227, 19, 147, 83, 211, 51, 179, 115, 243,
	11, 139, 75, 203, 43, 171, 107, 235, 27, 155, 91, 219, 59, 187, 123, 251,
	7, 135, 71, 199, 39, 167, 103, 231, 23, 151, 87, 215, 55, 183, 119, 247,
	15, 143, 79, 207, 47, 175, 111, 239, 31, 159, 95, 223, 63, 191, 127, 255,
}

// Returns a bits-reversed byte based on the input
func Reverse(b byte) byte {
	return reversed[b]
}

// The Vector interface defines methods on the bit vector
type Vector interface {
	// Retrieves the bit at the designated position
	Peek(uint) bool
	// Sets the bit at the designated position
	Poke(uint, bool)
	// Flips the bit at the designated position
	Flip(uint)
	// Returns the total number of elements supported by the vector
	Len() uint
}

// A Vector type that simply stores booleans in a slice
type boolvector []bool

// Retrieves the bit at the designated position
func (v boolvector) Peek(pos uint) bool {
	return v[pos]
}

// Sets the bit at the designated position
func (v boolvector) Poke(pos uint, val bool) {
	v[pos] = val
}

// Flips the bit at the designated position
func (v boolvector) Flip(pos uint) {
	v[pos] = !v[pos]
}

// Returns the total number of elements supported by the vector
func (v boolvector) Len() uint {
	return uint(len(v))
}

func NewBool(size uint) Vector {
	var slice []bool = make([]bool, size)
	return boolvector(slice)
}

type vector []uint

// Retrieves the bit at the designated position
func (v vector) Peek(pos uint) bool {
	var slot, offset uint = at(pos)
	return (v[slot] & (1 << offset)) > 0
}

// Sets the bit at the designated position
func (v vector) Poke(pos uint, val bool) {
	var slot, offset uint = at(pos)
	if val {
		v[slot] |= (1 << offset)
	} else {
		v[slot] &^= (1 << offset)
	}
}

// Flips the bit at the designated position
func (v vector) Flip(pos uint) {
	var slot, offset uint = at(pos)
	v[slot] ^= (1 << offset)
}

// Returns the total number of elements supported by the vector
func (v vector) Len() uint {
	return uint(len(v) * strconv.IntSize)
}

func at(pos uint) (uint, uint) {
	return pos / strconv.IntSize, pos % strconv.IntSize
}

// Create a new bit vector sized up to the desired number of elements
func NewBit(size uint) Vector {
	var length uint = size / strconv.IntSize

	if size%strconv.IntSize > 0 {
		// Allocate one additional slot if the desired
		// size is does not divide by 32/64
		length++
	}

	var slice []uint = make([]uint, length)
	return vector(slice)
}
