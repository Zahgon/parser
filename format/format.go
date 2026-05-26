// Copyright (c) 2014 The sortutil Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSES/STRUTIL-LICENSE file.

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

package format

import (
	"io"
)

const (
	st0 = iota
	stBOL
	stPERC
	stBOLPERC
)

// Formatter is an io.Writer extended formatter by a fmt.Printf like function Format.
type Formatter interface {
	io.Writer
	Format(format string, args ...interface{}) (n int, errno error)
}

type indentFormatter struct {
	io.Writer
	indent      []byte
	indentLevel int
	state       int
}

var replace = map[rune]string{
	'\000': "\\0",
	'\'':   "''",
	'\n':   "\\n",
	'\r':   "\\r",
}

// IndentFormatter returns a new Formatter which interprets %i and %u in the
// Format() formats string as indent and unindent commands. The commands can
// nest. The Formatter writes to io.Writer 'w' and inserts one 'indent'
// string per current indent level value.
// Behaviour of commands reaching negative indent levels is undefined.
//
//	IndentFormatter(os.Stdout, "\t").Format("abc%d%%e%i\nx\ny\n%uz\n", 3)
//
// output:
//
//	abc3%e
//	    x
//	    y
//	z
//
// The Go quoted string literal form of the above is:
//
//	"abc%%e\n\tx\n\tx\nz\n"
//
// The commands can be scattered between separate invocations of Format(),
// i.e. the formatter keeps track of the indent level and knows if it is
// positioned on start of a line and should emit indentation(s).
// The same output as above can be produced by e.g.:
//
//	f := IndentFormatter(os.Stdout, " ")
//	f.Format("abc%d%%e%i\nx\n", 3)
//	f.Format("y\n%uz\n")
func IndentFormatter(w io.Writer, indent string) Formatter {
	_ = "STUB: not implemented"
	return *new(Formatter)
}

func (f *indentFormatter) format(flat bool, format string, args ...interface{}) (n int, errno error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Format implements Format interface.
func (f *indentFormatter) Format(format string, args ...interface{}) (n int, errno error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type flatFormatter indentFormatter

// FlatFormatter returns a newly created Formatter with the same functionality as the one returned
// by IndentFormatter except it allows a newline in the 'format' string argument of Format
// to pass through if the indent level is current zero.
//
// If the indent level is non-zero then such new lines are changed to a space character.
// There is no indent string, the %i and %u format verbs are used solely to determine the indent level.
//
// The FlatFormatter is intended for flattening of normally nested structure textual representation to
// a one top level structure per line form.
//
//	FlatFormatter(os.Stdout, " ").Format("abc%d%%e%i\nx\ny\n%uz\n", 3)
//
// output in the form of a Go quoted string literal:
//
//	"abc3%%e x y z\n"
func FlatFormatter(w io.Writer) Formatter { _ = "STUB: not implemented"; return *new(Formatter) }

// Format implements Format interface.
func (f *flatFormatter) Format(format string, args ...interface{}) (n int, errno error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// OutputFormat output escape character with backslash.
func OutputFormat(s string) string { _ = "STUB: not implemented"; return "" }

// RestoreFlag mark the Restore format
type RestoreFlags uint64

// Mutually exclusive group of `RestoreFlags`:
// [RestoreStringSingleQuotes, RestoreStringDoubleQuotes]
// [RestoreKeyWordUppercase, RestoreKeyWordLowercase]
// [RestoreNameUppercase, RestoreNameLowercase]
// [RestoreNameDoubleQuotes, RestoreNameBackQuotes]
// The flag with the left position in each group has a higher priority.
const (
	RestoreStringSingleQuotes RestoreFlags = 1 << iota
	RestoreStringDoubleQuotes
	RestoreStringEscapeBackslash

	RestoreKeyWordUppercase
	RestoreKeyWordLowercase

	RestoreNameUppercase
	RestoreNameLowercase
	RestoreNameDoubleQuotes
	RestoreNameBackQuotes

	RestoreSpacesAroundBinaryOperation
	RestoreBracketAroundBinaryOperation

	RestoreStringWithoutCharset
	RestoreStringWithoutDefaultCharset

	RestoreTiDBSpecialComment
)

const (
	DefaultRestoreFlags = RestoreStringSingleQuotes | RestoreKeyWordUppercase | RestoreNameBackQuotes
)

func (rf RestoreFlags) has(flag RestoreFlags) bool { _ = "STUB: not implemented"; return false }

// HasStringSingleQuotesFlag returns a boolean indicating when `rf` has `RestoreStringSingleQuotes` flag.
func (rf RestoreFlags) HasStringSingleQuotesFlag() bool { _ = "STUB: not implemented"; return false }

// HasStringDoubleQuotesFlag returns a boolean indicating whether `rf` has `RestoreStringDoubleQuotes` flag.
func (rf RestoreFlags) HasStringDoubleQuotesFlag() bool { _ = "STUB: not implemented"; return false }

// HasStringEscapeBackslashFlag returns a boolean indicating whether `rf` has `RestoreStringEscapeBackslash` flag.
func (rf RestoreFlags) HasStringEscapeBackslashFlag() bool { _ = "STUB: not implemented"; return false }

// HasKeyWordUppercaseFlag returns a boolean indicating whether `rf` has `RestoreKeyWordUppercase` flag.
func (rf RestoreFlags) HasKeyWordUppercaseFlag() bool { _ = "STUB: not implemented"; return false }

// HasKeyWordLowercaseFlag returns a boolean indicating whether `rf` has `RestoreKeyWordLowercase` flag.
func (rf RestoreFlags) HasKeyWordLowercaseFlag() bool { _ = "STUB: not implemented"; return false }

// HasNameUppercaseFlag returns a boolean indicating whether `rf` has `RestoreNameUppercase` flag.
func (rf RestoreFlags) HasNameUppercaseFlag() bool { _ = "STUB: not implemented"; return false }

// HasNameLowercaseFlag returns a boolean indicating whether `rf` has `RestoreNameLowercase` flag.
func (rf RestoreFlags) HasNameLowercaseFlag() bool { _ = "STUB: not implemented"; return false }

// HasNameDoubleQuotesFlag returns a boolean indicating whether `rf` has `RestoreNameDoubleQuotes` flag.
func (rf RestoreFlags) HasNameDoubleQuotesFlag() bool { _ = "STUB: not implemented"; return false }

// HasNameBackQuotesFlag returns a boolean indicating whether `rf` has `RestoreNameBackQuotes` flag.
func (rf RestoreFlags) HasNameBackQuotesFlag() bool { _ = "STUB: not implemented"; return false }

// HasSpacesAroundBinaryOperationFlag returns a boolean indicating whether `rf` has `RestoreSpacesAroundBinaryOperation` flag.
func (rf RestoreFlags) HasSpacesAroundBinaryOperationFlag() bool {
	_ = "STUB: not implemented"
	return false
}

func (rf RestoreFlags) HasRestoreBracketAroundBinaryOperation() bool {
	_ = "STUB: not implemented"
	return false
}

func (rf RestoreFlags) HasStringWithoutDefaultCharset() bool {
	_ = "STUB: not implemented"
	return false
}

func (rf RestoreFlags) HasStringWithoutCharset() bool { _ = "STUB: not implemented"; return false }

func (rf RestoreFlags) HasTiDBSpecialCommentFlag() bool { _ = "STUB: not implemented"; return false }

// RestoreCtx is `Restore` context to hold flags and writer.
type RestoreCtx struct {
	Flags     RestoreFlags
	In        io.Writer
	DefaultDB string
	CTENames  []string
}

// NewRestoreCtx returns a new `RestoreCtx`.
func NewRestoreCtx(flags RestoreFlags, in io.Writer) *RestoreCtx {
	_ = "STUB: not implemented"
	return nil
}

// WriteKeyWord writes the `keyWord` into writer.
// `keyWord` will be converted format(uppercase and lowercase for now) according to `RestoreFlags`.
func (ctx *RestoreCtx) WriteKeyWord(keyWord string) { _ = "STUB: not implemented"; return }

func (ctx *RestoreCtx) WriteWithSpecialComments(featureID string, fn func()) {
	_ = "STUB: not implemented"
	return
}

// WriteString writes the string into writer
// `str` may be wrapped in quotes and escaped according to RestoreFlags.
func (ctx *RestoreCtx) WriteString(str string) { _ = "STUB: not implemented"; return }

// WriteName writes the name into writer
// `name` maybe wrapped in quotes and escaped according to RestoreFlags.
func (ctx *RestoreCtx) WriteName(name string) { _ = "STUB: not implemented"; return }

// WritePlain writes the plain text into writer without any handling.
func (ctx *RestoreCtx) WritePlain(plainText string) { _ = "STUB: not implemented"; return }

// WritePlainf write the plain text into writer without any handling.
func (ctx *RestoreCtx) WritePlainf(format string, a ...interface{}) {
	_ = "STUB: not implemented"
	return
}
