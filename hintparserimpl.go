// Copyright 2020 PingCAP, Inc.
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
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/parser/terror"
)

var (
	ErrWarnOptimizerHintUnsupportedHint = terror.ClassParser.NewStd(mysql.ErrWarnOptimizerHintUnsupportedHint)
	ErrWarnOptimizerHintInvalidToken    = terror.ClassParser.NewStd(mysql.ErrWarnOptimizerHintInvalidToken)
	ErrWarnMemoryQuotaOverflow          = terror.ClassParser.NewStd(mysql.ErrWarnMemoryQuotaOverflow)
	ErrWarnOptimizerHintParseError      = terror.ClassParser.NewStd(mysql.ErrWarnOptimizerHintParseError)
	ErrWarnOptimizerHintInvalidInteger  = terror.ClassParser.NewStd(mysql.ErrWarnOptimizerHintInvalidInteger)
)

// hintScanner implements the yyhintLexer interface
type hintScanner struct {
	Scanner
}

func (hs *hintScanner) Errorf(format string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (hs *hintScanner) Lex(lval *yyhintSymType) int { _ = "STUB: not implemented"; return 0 }

type hintParser struct {
	lexer  hintScanner
	result []*ast.TableOptimizerHint

	// the following fields are used by yyParse to reduce allocation.
	cache  []yyhintSymType
	yylval yyhintSymType
	yyVAL  *yyhintSymType
}

func newHintParser() *hintParser { _ = "STUB: not implemented"; return nil }

func (hp *hintParser) parse(input string, sqlMode mysql.SQLMode, initPos Pos) ([]*ast.TableOptimizerHint, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skipped the initial '/*+'

// skip the final '*/' (we need the '*/' for reporting warnings)

// ParseHint parses an optimizer hint (the interior of `/*+ ... */`).
func ParseHint(input string, sqlMode mysql.SQLMode, initPos Pos) ([]*ast.TableOptimizerHint, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hp *hintParser) warnUnsupportedHint(name string) { _ = "STUB: not implemented"; return }

func (hp *hintParser) lastErrorAsWarn() { _ = "STUB: not implemented"; return }
