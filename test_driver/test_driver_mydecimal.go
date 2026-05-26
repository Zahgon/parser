// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !codes
// +build !codes

package test_driver

const panicInfo = "This branch is not implemented. " +
	"This is because you are trying to test something specific to TiDB's MyDecimal implementation. " +
	"It is recommended to do this in TiDB repository."

// constant values.
const (
	maxWordBufLen = 9 // A MyDecimal holds 9 words.
	digitsPerWord = 9 // A word holds 9 digits.
	digMask       = 100000000
)

var (
	wordBufLen = 9
)

// fixWordCntError limits word count in wordBufLen, and returns overflow or truncate error.
func fixWordCntError(wordsInt, wordsFrac int) (newWordsInt int, newWordsFrac int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

/*
countLeadingZeroes returns the number of leading zeroes that can be removed from fraction.

@param   i    start index
@param   word value to compare against list of powers of 10
*/
func countLeadingZeroes(i int, word int32) int { _ = "STUB: not implemented"; return 0 }

func digitsToWords(digits int) int { _ = "STUB: not implemented"; return 0 }

// MyDecimal represents a decimal value.
type MyDecimal struct {
	digitsInt int8 // the number of *decimal* digits before the point.

	digitsFrac int8 // the number of decimal digits after the point.

	resultFrac int8 // result fraction digits.

	negative bool

	// wordBuf is an array of int32 words.
	// A word is an int32 value can hold 9 digits.(0 <= word < wordBase)
	wordBuf [maxWordBufLen]int32
}

// String returns the decimal string representation rounded to resultFrac.
func (d *MyDecimal) String() string { _ = "STUB: not implemented"; return "" }

func (d *MyDecimal) stringSize() int {
	_ = "STUB: not implemented"
	// sign, zero integer and dot.
	return 0
}

func (d *MyDecimal) removeLeadingZeros() (wordIdx int, digitsInt int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// ToString converts decimal to its printable string representation without rounding.
//
//	RETURN VALUE
//
//	    str       - result string
//	    errCode   - eDecOK/eDecTruncate/eDecOverflow
func (d *MyDecimal) ToString() (str []byte) { _ = "STUB: not implemented"; return nil }

/* symbol 0 before digital point */

// FromString parses decimal from string.
func (d *MyDecimal) FromString(str []byte) error { _ = "STUB: not implemented"; return nil }
