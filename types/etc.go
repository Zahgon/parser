// Copyright 2014 The ql Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSES/QL-LICENSE file.

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
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/parser/terror"
)

// IsTypeBlob returns a boolean indicating whether the tp is a blob type.
func IsTypeBlob(tp byte) bool { _ = "STUB: not implemented"; return false }

// IsTypeChar returns a boolean indicating
// whether the tp is the char type like a string type or a varchar type.
func IsTypeChar(tp byte) bool { _ = "STUB: not implemented"; return false }

var type2Str = map[byte]string{
	mysql.TypeBit:         "bit",
	mysql.TypeBlob:        "text",
	mysql.TypeDate:        "date",
	mysql.TypeDatetime:    "datetime",
	mysql.TypeUnspecified: "unspecified",
	mysql.TypeNewDecimal:  "decimal",
	mysql.TypeDouble:      "double",
	mysql.TypeEnum:        "enum",
	mysql.TypeFloat:       "float",
	mysql.TypeGeometry:    "geometry",
	mysql.TypeInt24:       "mediumint",
	mysql.TypeJSON:        "json",
	mysql.TypeLong:        "int",
	mysql.TypeLonglong:    "bigint",
	mysql.TypeLongBlob:    "longtext",
	mysql.TypeMediumBlob:  "mediumtext",
	mysql.TypeNull:        "null",
	mysql.TypeSet:         "set",
	mysql.TypeShort:       "smallint",
	mysql.TypeString:      "char",
	mysql.TypeDuration:    "time",
	mysql.TypeTimestamp:   "timestamp",
	mysql.TypeTiny:        "tinyint",
	mysql.TypeTinyBlob:    "tinytext",
	mysql.TypeVarchar:     "varchar",
	mysql.TypeVarString:   "var_string",
	mysql.TypeYear:        "year",
}

// TypeStr converts tp to a string.
func TypeStr(tp byte) (r string) { _ = "STUB: not implemented"; return "" }

// TypeToStr converts a field to a string.
// It is used for converting Text to Blob,
// or converting Char to Binary.
// Args:
//
//	tp: type enum
//	cs: charset
func TypeToStr(tp byte, cs string) (r string) { _ = "STUB: not implemented"; return "" }

var (
	dig2bytes = [10]int{0, 1, 1, 2, 2, 3, 3, 4, 4, 4}
)

// constant values.
const (
	digitsPerWord = 9 // A word holds 9 digits.
	wordSize      = 4 // A word is 4 bytes int32.
)

// ErrInvalidDefault is returned when meet a invalid default value.
var ErrInvalidDefault = terror.ClassTypes.NewStd(mysql.ErrInvalidDefault)
