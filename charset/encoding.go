// Copyright 2021 PingCAP, Inc.
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

package charset

import (
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/parser/terror"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

const encodingLegacy = "utf-8" // utf-8 encoding is compatible with old default behavior.

var errInvalidCharacterString = terror.ClassParser.NewStd(mysql.ErrInvalidCharacterString)

type EncodingLabel string

// Format trim and change the label to lowercase.
func Format(label string) EncodingLabel { _ = "STUB: not implemented"; return *new(EncodingLabel) }

// Formatted is used when the label is already trimmed and it is lowercase.
func Formatted(label string) EncodingLabel { _ = "STUB: not implemented"; return *new(EncodingLabel) }

// Encoding provide a interface to encode/decode a string with specific encoding.
type Encoding struct {
	enc        encoding.Encoding
	name       string
	charLength func([]byte) int
}

// Enabled indicates whether the non-utf8 encoding is used.
func (e *Encoding) Enabled() bool { _ = "STUB: not implemented"; return false }

// Name returns the name of the current encoding.
func (e *Encoding) Name() string {
	_ = "STUB: not implemented"

	// NewEncoding creates a new Encoding.
	return ""
}

func NewEncoding(label string) *Encoding { _ = "STUB: not implemented"; return nil }

// UpdateEncoding updates to a new Encoding.
func (e *Encoding) UpdateEncoding(label EncodingLabel) { _ = "STUB: not implemented"; return }

// Encode convert bytes from utf-8 charset to a specific charset.
func (e *Encoding) Encode(dest, src []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decode convert bytes from a specific charset to utf-8 charset.
func (e *Encoding) Decode(dest, src []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Encoding) transform(transformer transform.Transformer, dest, src []byte, isDecoding bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip the source bytes that cannot be decoded normally.

// The source bytes are exhausted.

func (e *Encoding) nextCharLenInSrc(srcRest []byte, isDecoding bool) int {
	_ = "STUB: not implemented"
	return 0
}

func enlargeCapacity(dest []byte) []byte { _ = "STUB: not implemented"; return nil }

func (e *Encoding) generateErr(srcRest []byte, srcNextLen int) error {
	_ = "STUB: not implemented"
	return nil
}

// replacementBytes are bytes for the replacement rune 0xfffd.
var replacementBytes = []byte{0xEF, 0xBF, 0xBD}

// beginWithReplacementChar check if dst has the prefix '0xEFBFBD'.
func beginWithReplacementChar(dst []byte) bool { _ = "STUB: not implemented"; return false }
