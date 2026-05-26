// Copyright 2015 PingCAP, Inc.
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

package types

import (
	"io"

	"github.com/pingcap/parser/format"
)

// UnspecifiedLength is unspecified length.
const (
	UnspecifiedLength = -1
)

// TiDBStrictIntegerDisplayWidth represent whether return warnings when integerType with (length) was parsed.
// The default is `false`, it will be parsed as warning, and the result in show-create-table will ignore the
// display length when it set to `true`. This is for compatibility with MySQL 8.0 in which integer max display
// length is deprecated, referring this issue #6688 for more details.
var (
	TiDBStrictIntegerDisplayWidth bool
)

// FieldType records field type information.
type FieldType struct {
	Tp      byte
	Flag    uint
	Flen    int
	Decimal int
	Charset string
	Collate string
	// Elems is the element list for enum and set type.
	Elems []string
}

// NewFieldType returns a FieldType,
// with a type and other information about field type.
func NewFieldType(tp byte) *FieldType { _ = "STUB: not implemented"; return nil }

// Clone returns a copy of itself.
func (ft *FieldType) Clone() *FieldType { _ = "STUB: not implemented"; return nil }

// Equal checks whether two FieldType objects are equal.
func (ft *FieldType) Equal(other *FieldType) bool {
	_ = "STUB: not implemented"
	// We do not need to compare whole `ft.Flag == other.Flag` when wrapping cast upon an Expression.
	// but need compare unsigned_flag of ft.Flag.
	// When Tp is float or double with Decimal unspecified, do not check whether Flen is equal,
	// because Flen for them is useless.
	// The Decimal field can be ignored if the type is int or string.
	return false
}

// EvalType gets the type in evaluation.
func (ft *FieldType) EvalType() EvalType { _ = "STUB: not implemented"; return *new(EvalType) }

// Hybrid checks whether a type is a hybrid type, which can represent different types of value in specific context.
func (ft *FieldType) Hybrid() bool { _ = "STUB: not implemented"; return false }

// Init initializes the FieldType data.
func (ft *FieldType) Init(tp byte) { _ = "STUB: not implemented"; return }

// CompactStr only considers Tp/CharsetBin/Flen/Deimal.
// This is used for showing column type in infoschema.
func (ft *FieldType) CompactStr() string { _ = "STUB: not implemented"; return "" }

// displayFlen and displayDecimal are flen and decimal values with `-1` substituted with default value.

// Format is ENUM ('e1', 'e2') or SET ('e1', 'e2')

// 1. Flen Not Default, Decimal Not Default -> Valid
// 2. Flen Not Default, Decimal Default (-1) -> Invalid
// 3. Flen Default, Decimal Not Default -> Valid
// 4. Flen Default, Decimal Default -> Valid (hide)

// Referring this issue #6688, the integer max display length is deprecated in MySQL 8.0.
// Since the length doesn't take any effect in TiDB storage or showing result, we remove it here.

// InfoSchemaStr joins the CompactStr with unsigned flag and
// returns a string.
func (ft *FieldType) InfoSchemaStr() string { _ = "STUB: not implemented"; return "" }

// String joins the information of FieldType and returns a string.
// Note: when flen or decimal is unspecified, this function will use the default value instead of -1.
func (ft *FieldType) String() string { _ = "STUB: not implemented"; return "" }

// Restore implements Node interface.
func (ft *FieldType) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// RestoreAsCastType is used for write AST back to string.
func (ft *FieldType) RestoreAsCastType(ctx *format.RestoreCtx, explicitCharset bool) {
	_ = "STUB: not implemented"
	return
}

// FormatAsCastType is used for write AST back to string.
func (ft *FieldType) FormatAsCastType(w io.Writer, explicitCharset bool) {
	_ = "STUB: not implemented"
	return
}

// VarStorageLen indicates this column is a variable length column.
const VarStorageLen = -1

// StorageLength is the length of stored value for the type.
func (ft *FieldType) StorageLength() int { _ = "STUB: not implemented"; return 0 }

// This may not be the accurate length, because we may encode them as varint.

// HasCharset indicates if a COLUMN has an associated charset. Returning false here prevents some information
// statements(like `SHOW CREATE TABLE`) from attaching a CHARACTER SET clause to the column.
func HasCharset(ft *FieldType) bool { _ = "STUB: not implemented"; return false }
