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

import (
	"math"
)

func isSpace(c byte) bool { _ = "STUB: not implemented"; return false }

func isDigit(c byte) bool { _ = "STUB: not implemented"; return false }

func myMin(a, b int) int { _ = "STUB: not implemented"; return 0 }

func pow10(x int) int32 { _ = "STUB: not implemented"; return 0 }

func Abs(n int64) int64 { _ = "STUB: not implemented"; return 0 }

// uintSizeTable is used as a table to do comparison to get uint length is faster than doing loop on division with 10
var uintSizeTable = [21]uint64{
	0, // redundant 0 here, so to make function StrLenOfUint64Fast to count from 1 and return i directly
	9, 99, 999, 9999, 99999,
	999999, 9999999, 99999999, 999999999, 9999999999,
	99999999999, 999999999999, 9999999999999, 99999999999999, 999999999999999,
	9999999999999999, 99999999999999999, 999999999999999999, 9999999999999999999,
	math.MaxUint64,
} // math.MaxUint64 is 18446744073709551615 and it has 20 digits

// StrLenOfUint64Fast efficiently calculate the string character lengths of an uint64 as input
func StrLenOfUint64Fast(x uint64) int { _ = "STUB: not implemented"; return 0 }

// StrLenOfInt64Fast efficiently calculate the string character lengths of an int64 as input
func StrLenOfInt64Fast(x int64) int { _ = "STUB: not implemented"; return 0 }

// add "-" sign on the length count
