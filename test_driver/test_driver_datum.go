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
	"github.com/pingcap/parser/types"
)

// Kind constants.
const (
	KindNull          byte = 0
	KindInt64         byte = 1
	KindUint64        byte = 2
	KindFloat32       byte = 3
	KindFloat64       byte = 4
	KindString        byte = 5
	KindBytes         byte = 6
	KindBinaryLiteral byte = 7 // Used for BIT / HEX literals.
	KindMysqlDecimal  byte = 8
	KindMysqlDuration byte = 9
	KindMysqlEnum     byte = 10
	KindMysqlBit      byte = 11 // Used for BIT table column values.
	KindMysqlSet      byte = 12
	KindMysqlTime     byte = 13
	KindInterface     byte = 14
	KindMinNotNull    byte = 15
	KindMaxValue      byte = 16
	KindRaw           byte = 17
	KindMysqlJSON     byte = 18
)

// Datum is a data box holds different kind of data.
// It has better performance and is easier to use than `interface{}`.
type Datum struct {
	k byte        // datum kind.
	i int64       // i can hold int64 uint64 float64 values.
	b []byte      // b can hold string or []byte values.
	x interface{} // x hold all other types.
}

// Kind gets the kind of the datum.
func (d *Datum) Kind() byte {
	_ = "STUB: not implemented"

	// GetInt64 gets int64 value.
	return 0
}

func (d *Datum) GetInt64() int64 {
	_ = "STUB: not implemented"

	// SetInt64 sets int64 value.
	return 0
}

func (d *Datum) SetInt64(i int64) { _ = "STUB: not implemented"; return }

// GetUint64 gets uint64 value.
func (d *Datum) GetUint64() uint64 {
	_ = "STUB: not implemented"

	// SetUint64 sets uint64 value.
	return 0
}

func (d *Datum) SetUint64(i uint64) { _ = "STUB: not implemented"; return }

// GetFloat64 gets float64 value.
func (d *Datum) GetFloat64() float64 { _ = "STUB: not implemented"; return 0 }

// SetFloat64 sets float64 value.
func (d *Datum) SetFloat64(f float64) { _ = "STUB: not implemented"; return }

// GetFloat32 gets float32 value.
func (d *Datum) GetFloat32() float32 { _ = "STUB: not implemented"; return 0 }

// SetFloat32 sets float32 value.
func (d *Datum) SetFloat32(f float32) { _ = "STUB: not implemented"; return }

// GetString gets string value.
func (d *Datum) GetString() string {
	_ = "STUB: not implemented"

	// SetString sets string value.
	return ""
}

func (d *Datum) SetString(s string) { _ = "STUB: not implemented"; return }

// GetBytes gets bytes value.
func (d *Datum) GetBytes() []byte {
	_ = "STUB: not implemented"

	// SetBytes sets bytes value to datum.
	return nil
}

func (d *Datum) SetBytes(b []byte) { _ = "STUB: not implemented"; return }

// SetBytesAsString sets bytes value to datum as string type.
func (d *Datum) SetBytesAsString(b []byte) { _ = "STUB: not implemented"; return }

// GetInterface gets interface value.
func (d *Datum) GetInterface() interface{} {
	_ = "STUB: not implemented"

	// SetInterface sets interface to datum.
	return nil
}

func (d *Datum) SetInterface(x interface{}) { _ = "STUB: not implemented"; return }

// SetNull sets datum to nil.
func (d *Datum) SetNull() { _ = "STUB: not implemented"; return }

// GetBinaryLiteral gets Bit value
func (d *Datum) GetBinaryLiteral() BinaryLiteral {
	_ = "STUB: not implemented"

	// SetBinaryLiteral sets Bit value
	return *new(BinaryLiteral)
}

func (d *Datum) SetBinaryLiteral(b BinaryLiteral) { _ = "STUB: not implemented"; return }

// GetMysqlDecimal gets Decimal value
func (d *Datum) GetMysqlDecimal() *MyDecimal { _ = "STUB: not implemented"; return nil }

// SetMysqlDecimal sets Decimal value
func (d *Datum) SetMysqlDecimal(b *MyDecimal) { _ = "STUB: not implemented"; return }

// GetValue gets the value of the datum of any kind.
func (d *Datum) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

// SetValue sets any kind of value.
func (d *Datum) SetValue(val interface{}) { _ = "STUB: not implemented"; return }

// Store as BinaryLiteral for Bit and Hex literals

// NewDatum creates a new Datum from an interface{}.
func NewDatum(in interface{}) (d Datum) { _ = "STUB: not implemented"; return *new(Datum) }

// NewBytesDatum creates a new Datum from a byte slice.
func NewBytesDatum(b []byte) (d Datum) {
	_ = "STUB: not implemented"
	return *

	// NewStringDatum creates a new Datum from a string.
	new(Datum)
}

func NewStringDatum(s string) (d Datum) {
	_ = "STUB: not implemented"
	return *

	// MakeDatums creates datum slice from interfaces.
	new(Datum)
}

func MakeDatums(args ...interface{}) []Datum { _ = "STUB: not implemented"; return nil }

// BinaryLiteral is the internal type for storing bit / hex literal type.
type BinaryLiteral []byte

// BitLiteral is the bit literal type.
type BitLiteral BinaryLiteral

// HexLiteral is the hex literal type.
type HexLiteral BinaryLiteral

// ZeroBinaryLiteral is a BinaryLiteral literal with zero value.
var ZeroBinaryLiteral = BinaryLiteral{}

// String implements fmt.Stringer interface.
func (b BinaryLiteral) String() string { _ = "STUB: not implemented"; return "" }

// ToString returns the string representation for the literal.
func (b BinaryLiteral) ToString() string {
	_ = "STUB: not implemented"

	// ToBitLiteralString returns the bit literal representation for the literal.
	return ""
}

func (b BinaryLiteral) ToBitLiteralString(trimLeadingZero bool) string {
	_ = "STUB: not implemented"
	return ""
}

// ParseBitStr parses bit string.
// The string format can be b'val', B'val' or 0bval, val must be 0 or 1.
// See https://dev.mysql.com/doc/refman/5.7/en/bit-value-literals.html
func ParseBitStr(s string) (BinaryLiteral, error) {
	_ = "STUB: not implemented"
	return *new(BinaryLiteral), nil
}

// format is b'val' or B'val'

// here means format is not b'val', B'val' or 0bval.

// Pad with zero (slice from `-alignedLength`)

// NewBitLiteral parses bit string as BitLiteral type.
func NewBitLiteral(s string) (BitLiteral, error) {
	_ = "STUB: not implemented"
	return *new(BitLiteral), nil
}

// ToString implement ast.BinaryLiteral interface
func (b BitLiteral) ToString() string { _ = "STUB: not implemented"; return "" }

// ParseHexStr parses hexadecimal string literal.
// See https://dev.mysql.com/doc/refman/5.7/en/hexadecimal-literals.html
func ParseHexStr(s string) (BinaryLiteral, error) {
	_ = "STUB: not implemented"
	return *new(BinaryLiteral), nil
}

// format is x'val' or X'val'

// here means format is not x'val', X'val' or 0xval.

// NewHexLiteral parses hexadecimal string as HexLiteral type.
func NewHexLiteral(s string) (HexLiteral, error) {
	_ = "STUB: not implemented"
	return *new(HexLiteral), nil
}

// ToString implement ast.BinaryLiteral interface
func (b HexLiteral) ToString() string { _ = "STUB: not implemented"; return "" }

// SetBinChsClnFlag sets charset, collation as 'binary' and adds binaryFlag to FieldType.
func SetBinChsClnFlag(ft *types.FieldType) { _ = "STUB: not implemented"; return }

// DefaultFsp is the default digit of fractional seconds part.
// MySQL use 0 as the default Fsp.
const DefaultFsp = int8(0)

// DefaultTypeForValue returns the default FieldType for the value.
func DefaultTypeForValue(value interface{}, tp *types.FieldType, charset string, collate string) {
	_ = "STUB: not implemented"
	return
}

// TODO: tp.Flen should be len(x) * 3 (max bytes length of CharsetUTF8)
