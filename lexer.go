// Copyright 2016 PingCAP, Inc.
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

package parser

import (
	"bytes"

	"github.com/pingcap/parser/charset"
	"github.com/pingcap/parser/mysql"
)

var _ = yyLexer(&Scanner{})

// Pos represents the position of a token.
type Pos struct {
	Line   int
	Col    int
	Offset int
}

// Scanner implements the yyLexer interface.
type Scanner struct {
	r   reader
	buf bytes.Buffer

	encoding charset.Encoding

	errs         []error
	warns        []error
	stmtStartPos int

	// inBangComment is true if we are inside a `/*! ... */` block.
	// It is used to ignore a stray `*/` when scanning.
	inBangComment bool

	sqlMode mysql.SQLMode

	// If the lexer should recognize keywords for window function.
	// It may break the compatibility when support those keywords,
	// because some application may already use them as identifiers.
	supportWindowFunc bool

	// Whether record the original text keyword position to the AST node.
	skipPositionRecording bool

	// lastScanOffset indicates last offset returned by scan().
	// It's used to substring sql in syntax error message.
	lastScanOffset int

	// lastKeyword records the previous keyword returned by scan().
	// determine whether an optimizer hint should be parsed or ignored.
	lastKeyword int
	// lastKeyword2 records the keyword before lastKeyword, it is used
	// to disambiguate hint after for update, which should be ignored.
	lastKeyword2 int
	// lastKeyword3 records the keyword before lastKeyword2, it is used
	// to disambiguate hint after create binding for update, which should
	// be pertained.
	lastKeyword3 int

	// hintPos records the start position of the previous optimizer hint.
	lastHintPos Pos

	// true if a dot follows an identifier
	identifierDot bool
}

// Errors returns the errors and warns during a scan.
func (s *Scanner) Errors() (warns []error, errs []error) {
	_ = "STUB: not implemented"
	return nil,

		// reset resets the sql string to be scanned.
		nil
}

func (s *Scanner) reset(sql string) { _ = "STUB: not implemented"; return }

func (s *Scanner) stmtText() string { _ = "STUB: not implemented"; return "" }

// trim new line

// Errorf tells scanner something is wrong.
// Scanner satisfies yyLexer interface which need this function.
func (s *Scanner) Errorf(format string, a ...interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// AppendError sets error into scanner.
// Scanner satisfies yyLexer interface which need this function.
func (s *Scanner) AppendError(err error) { _ = "STUB: not implemented"; return }

func (s *Scanner) tryDecodeToUTF8String(sql string) string { _ = "STUB: not implemented"; return "" }

func (s *Scanner) getNextToken() int { _ = "STUB: not implemented"; return 0 }

// Lex returns a token and store the token value in v.
// Scanner satisfies yyLexer interface.
// 0 and invalid are special token id this function would return:
// return 0 tells parser that scanner meets EOF,
// return invalid tells parser that scanner meets illegal character.
func (s *Scanner) Lex(v *yySymType) int { _ = "STUB: not implemented"; return 0 }

// LexLiteral returns the value of the converted literal
func (s *Scanner) LexLiteral() interface{} { _ = "STUB: not implemented"; return nil }

// SetSQLMode sets the SQL mode for scanner.
func (s *Scanner) SetSQLMode(mode mysql.SQLMode) {
	_ = "STUB: not implemented"

	// GetSQLMode return the SQL mode of scanner.
	return
}

func (s *Scanner) GetSQLMode() mysql.SQLMode {
	_ = "STUB: not implemented"

	// EnableWindowFunc controls whether the scanner recognize the keywords of window function.
	return *new(mysql.SQLMode)
}

func (s *Scanner) EnableWindowFunc(val bool) { _ = "STUB: not implemented"; return }

// InheritScanner returns a new scanner object which inherits configurations from the parent scanner.
func (s *Scanner) InheritScanner(sql string) *Scanner { _ = "STUB: not implemented"; return nil }

// NewScanner returns a new scanner object.
func NewScanner(s string) *Scanner { _ = "STUB: not implemented"; return nil }

func (s *Scanner) handleIdent(lval *yySymType) int {
	_ = "STUB: not implemented"

	// A character string literal may have an optional character set introducer and COLLATE clause:
	// [_charset_name]'string' [COLLATE collation_name]
	// See https://dev.mysql.com/doc/refman/5.7/en/charset-literal.html
	return 0
}

func (s *Scanner) skipWhitespace() rune { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) scan() (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// when scanner meets EOF, the returned token should be 0,
// because 0 is a special token id to remind the parser that stream is end.

// search a trie to get a token.

func startWithXx(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func startWithNn(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// The National Character Set, N'some text' or n'some test'.
// See https://dev.mysql.com/doc/refman/5.7/en/string-literals.html
// and https://dev.mysql.com/doc/refman/5.7/en/charset-national.html

func startWithBb(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func startWithSharp(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func startWithDash(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func startWithSlash(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// we see '/*' so far.

// '/*!' MySQL-specific comments
// See http://dev.mysql.com/doc/refman/5.7/en/comments.html
// in '/*!', which we always recognize regardless of version.

// '/*T' maybe TiDB-specific comments

// '/*TX' is just normal comment.

// in '/*T!', try to match the pattern '/*T![feature1,feature2,...]'.

// '/*M' maybe MariaDB-specific comments
// no special treatment for now.

// '/*+' optimizer hints
// See https://dev.mysql.com/doc/refman/5.7/en/optimizer-hints.html

// only recognize optimizers hints directly followed by certain
// keywords like SELECT, INSERT, etc., only a special case "FOR UPDATE" needs to be handled
// we will report a warning in order to match MySQL's behavior, but the hint content will be ignored

// special case of `create binding for update`

// '/**' if the next char is '/' it would close the comment.

// standard C-like comment. read until we see '*/' then drop it.

// Meets */, means comment end.

// unclosed comment or other errors.

func startWithStar(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0,

		// skip and exit '/*!' if we see '*/'
		*new(Pos), ""
}

// otherwise it is just a normal star.

func startWithAt(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func scanIdentifier(s *Scanner) (int, Pos, string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func scanIdentifierOrString(s *Scanner) (tok int, lit string) {
	_ = "STUB: not implemented"
	return 0, ""
}

var (
	quotedIdentifier = -identifier
)

func scanQuotedIdent(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// don't return identifier in case that it's interpreted as keyword token later.

func startString(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0,

		// lazyBuf is used to avoid allocation if possible.
		// it has a useBuf field indicates whether bytes.Buffer is necessary. if
		// useBuf is false, we can avoid calling bytes.Buffer.String(), which
		// make a copy of data and cause allocation.
		*new(Pos), ""
}

type lazyBuf struct {
	useBuf bool
	r      *reader
	b      *bytes.Buffer
	p      *Pos
}

func (mb *lazyBuf) setUseBuf(str string) { _ = "STUB: not implemented"; return }

func (mb *lazyBuf) writeRune(r rune, w int) { _ = "STUB: not implemented"; return }

func (mb *lazyBuf) data() string { _ = "STUB: not implemented"; return "" }

func (s *Scanner) scanString() (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// handleEscape handles the case in scanString when previous char is '\'.
func handleEscape(s *Scanner) rune { _ = "STUB: not implemented"; return 0 }

/*
	\" \' \\ \n \0 \b \Z \r \t ==> escape to one char
	\% \_ ==> preserve both char
	other ==> remove \
*/

func startWithNumber(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

// 0x, 0x7fz3 are identifier

// 0b, 0b123, 0b1ab are identifier

// Identifiers may begin with a digit but unless quoted may not consist solely of digits.

func startWithDot(s *Scanner) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return 0, *new(Pos), ""
}

func (s *Scanner) scanOct() { _ = "STUB: not implemented"; return }

func (s *Scanner) scanHex() { _ = "STUB: not implemented"; return }

func (s *Scanner) scanBit() { _ = "STUB: not implemented"; return }

func (s *Scanner) scanFloat(beg *Pos) (tok int, pos Pos, lit string) {
	_ = "STUB: not implemented"

	// float = D1 . D2 e D3
	return 0, *new(Pos), ""
}

// D1 . D2 e XX when XX is not D3, parse the result to an identifier.
// 9e9e = 9e9(float) + e(identifier)
// 9est = 9est(identifier)

func (s *Scanner) scanDigits() string { _ = "STUB: not implemented"; return "" }

// scanVersionDigits scans for `min` to `max` digits (range inclusive) used in
// `/*!12345 ... */` comments.
func (s *Scanner) scanVersionDigits(min, max int) { _ = "STUB: not implemented"; return }

func (s *Scanner) scanFeatureIDs() (featureIDs []string) { _ = "STUB: not implemented"; return nil }

func (s *Scanner) lastErrorAsWarn() { _ = "STUB: not implemented"; return }

type reader struct {
	s string
	p Pos
	w int

	peekRune        rune
	peekRuneUpdated bool
}

var eof = Pos{-1, -1, -1}

func (r *reader) eof() bool { _ = "STUB: not implemented"; return false }

// peek() peeks a rune from underlying reader.
// if reader meets EOF, it will return unicode.ReplacementChar. to distinguish from
// the real unicode.ReplacementChar, the caller should call r.eof() again to check.
func (r *reader) peek() rune { _ = "STUB: not implemented"; return 0 }

// illegal encoding

// inc increase the position offset of the reader.
// peek must be called before calling inc!
func (r *reader) inc() { _ = "STUB: not implemented"; return }

func (r *reader) incN(n int) { _ = "STUB: not implemented"; return }

func (r *reader) readByte() (ch rune) { _ = "STUB: not implemented"; return 0 }

func (r *reader) pos() Pos { _ = "STUB: not implemented"; return *new(Pos) }

func (r *reader) updatePos(pos Pos) { _ = "STUB: not implemented"; return }

func (r *reader) data(from *Pos) string { _ = "STUB: not implemented"; return "" }

func (r *reader) incAsLongAs(fn func(rune) bool) rune { _ = "STUB: not implemented"; return 0 }
