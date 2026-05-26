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

package parser

import (
	"regexp"

	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/auth"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/parser/terror"
)

var (
	// ErrSyntax returns for sql syntax error.
	ErrSyntax = terror.ClassParser.NewStd(mysql.ErrSyntax)
	// ErrParse returns for sql parse error.
	ErrParse = terror.ClassParser.NewStd(mysql.ErrParse)
	// ErrUnknownCharacterSet returns for no character set found error.
	ErrUnknownCharacterSet = terror.ClassParser.NewStd(mysql.ErrUnknownCharacterSet)
	// ErrInvalidYearColumnLength returns for illegal column length for year type.
	ErrInvalidYearColumnLength = terror.ClassParser.NewStd(mysql.ErrInvalidYearColumnLength)
	// ErrWrongArguments returns for illegal argument.
	ErrWrongArguments = terror.ClassParser.NewStd(mysql.ErrWrongArguments)
	// ErrWrongFieldTerminators returns for illegal field terminators.
	ErrWrongFieldTerminators = terror.ClassParser.NewStd(mysql.ErrWrongFieldTerminators)
	// ErrTooBigDisplayWidth returns for data display width exceed limit .
	ErrTooBigDisplayWidth = terror.ClassParser.NewStd(mysql.ErrTooBigDisplaywidth)
	// ErrTooBigPrecision returns for data precision exceed limit.
	ErrTooBigPrecision = terror.ClassParser.NewStd(mysql.ErrTooBigPrecision)
	// ErrUnknownAlterLock returns for no alter lock type found error.
	ErrUnknownAlterLock = terror.ClassParser.NewStd(mysql.ErrUnknownAlterLock)
	// ErrUnknownAlterAlgorithm returns for no alter algorithm found error.
	ErrUnknownAlterAlgorithm = terror.ClassParser.NewStd(mysql.ErrUnknownAlterAlgorithm)
	// ErrWrongValue returns for wrong value
	ErrWrongValue = terror.ClassParser.NewStd(mysql.ErrWrongValue)
	// ErrWarnDeprecatedSyntaxNoReplacement return when the syntax was deprecated and there is no replacement.
	ErrWarnDeprecatedSyntaxNoReplacement = terror.ClassParser.NewStd(mysql.ErrWarnDeprecatedSyntaxNoReplacement)
	// ErrWarnDeprecatedIntegerDisplayWidth share the same code 1681, and it will be returned when length is specified in integer.
	ErrWarnDeprecatedIntegerDisplayWidth = terror.ClassParser.NewStdErr(mysql.ErrWarnDeprecatedSyntaxNoReplacement, mysql.Message("Integer display width is deprecated and will be removed in a future release.", nil))
	// ErrWrongUsage returns for incorrect usages.
	ErrWrongUsage = terror.ClassParser.NewStd(mysql.ErrWrongUsage)
	// SpecFieldPattern special result field pattern
	SpecFieldPattern = regexp.MustCompile(`(\/\*!(M?[0-9]{5,6})?|\*\/)`)
	specCodeStart    = regexp.MustCompile(`^\/\*!(M?[0-9]{5,6})?[ \t]*`)
	specCodeEnd      = regexp.MustCompile(`[ \t]*\*\/$`)
)

// TrimComment trim comment for special comment code of MySQL.
func TrimComment(txt string) string { _ = "STUB: not implemented"; return "" }

type ParserConfig struct {
	EnableWindowFunction        bool
	EnableStrictDoubleTypeCheck bool
	SkipPositionRecording       bool
	CharsetClient               string // CharsetClient indicates how to decode the original SQL.
}

// Parser represents a parser instance. Some temporary objects are stored in it to reduce object allocation during Parse function.
type Parser struct {
	charset    string
	collation  string
	result     []ast.StmtNode
	src        string
	lexer      Scanner
	hintParser *hintParser

	explicitCharset       bool
	strictDoubleFieldType bool

	// the following fields are used by yyParse to reduce allocation.
	cache  []yySymType
	yylval yySymType
	yyVAL  *yySymType
}

func yySetOffset(yyVAL *yySymType, offset int) { _ = "STUB: not implemented"; return }

func yyhintSetOffset(_ *yyhintSymType, _ int) { _ = "STUB: not implemented"; return }

type stmtTexter interface {
	stmtText() string
}

// New returns a Parser object with default SQL mode.
func New() *Parser { _ = "STUB: not implemented"; return nil }

func (parser *Parser) SetStrictDoubleTypeCheck(val bool) { _ = "STUB: not implemented"; return }

func (parser *Parser) SetParserConfig(config ParserConfig) { _ = "STUB: not implemented"; return }

// Parse parses a query string to raw ast.StmtNode.
// If charset or collation is "", default charset and collation will be used.
func (parser *Parser) Parse(sql, charset, collation string) (stmt []ast.StmtNode, warns []error, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (parser *Parser) lastErrorAsWarn() { _ = "STUB: not implemented"; return }

// ParseOneStmt parses a query and returns an ast.StmtNode.
// The query must have one statement, otherwise ErrSyntax is returned.
func (parser *Parser) ParseOneStmt(sql, charset, collation string) (ast.StmtNode, error) {
	_ = "STUB: not implemented"
	return *new(ast.StmtNode), nil
}

// SetSQLMode sets the SQL mode for parser.
func (parser *Parser) SetSQLMode(mode mysql.SQLMode) { _ = "STUB: not implemented"; return }

// EnableWindowFunc controls whether the parser to parse syntax related with window function.
func (parser *Parser) EnableWindowFunc(val bool) { _ = "STUB: not implemented"; return }

// ParseErrorWith returns "You have a syntax error near..." error message compatible with mysql.
func ParseErrorWith(errstr string, lineno int) error { _ = "STUB: not implemented"; return nil }

// The select statement is not at the end of the whole statement, if the last
// field text was set from its offset to the end of the src string, update
// the last field text.
func (parser *Parser) setLastSelectFieldText(st *ast.SelectStmt, lastEnd int) {
	_ = "STUB: not implemented"
	return
}

func (parser *Parser) startOffset(v *yySymType) int { _ = "STUB: not implemented"; return 0 }

func (parser *Parser) endOffset(v *yySymType) int { _ = "STUB: not implemented"; return 0 }

func (parser *Parser) parseHint(input string) ([]*ast.TableOptimizerHint, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toInt(l yyLexer, lval *yySymType, str string) int { _ = "STUB: not implemented"; return 0 }

// TODO: toDecimal maybe out of range still.
// This kind of error should be throw to higher level, because truncated data maybe legal.
// For example, this SQL returns error:
// create table test (id decimal(30, 0));
// insert into test values(123456789012345678901234567890123094839045793405723406801943850);
// While this SQL:
// select 1234567890123456789012345678901230948390457934057234068019438509023041874359081325875128590860234789847359871045943057;
// get value 99999999999999999999999999999999999999999999999999999999999999999

func toDecimal(l yyLexer, lval *yySymType, str string) int { _ = "STUB: not implemented"; return 0 }

func toFloat(l yyLexer, lval *yySymType, str string) int { _ = "STUB: not implemented"; return 0 }

// See https://dev.mysql.com/doc/refman/5.7/en/hexadecimal-literals.html
func toHex(l yyLexer, lval *yySymType, str string) int { _ = "STUB: not implemented"; return 0 }

// See https://dev.mysql.com/doc/refman/5.7/en/bit-type.html
func toBit(l yyLexer, lval *yySymType, str string) int { _ = "STUB: not implemented"; return 0 }

func getUint64FromNUM(num interface{}) uint64 { _ = "STUB: not implemented"; return 0 }

func getInt64FromNUM(num interface{}) (val int64, errMsg string) {
	_ = "STUB: not implemented"
	return 0, ""
}

// convertToRole tries to convert elements of roleOrPrivList to RoleIdentity
func convertToRole(roleOrPrivList []*ast.RoleOrPriv) ([]*auth.RoleIdentity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertToPriv tries to convert elements of roleOrPrivList to PrivElem
func convertToPriv(roleOrPrivList []*ast.RoleOrPriv) ([]*ast.PrivElem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
